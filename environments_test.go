package staxxapi

import (
	"fmt"
	"testing"
	"time"

	"github.com/adiletabylov/staxxapi/model"
	"github.com/stretchr/testify/assert"
)

func TestListEmv(t *testing.T) {
	Init("http://localhost", "4000")
	resp, err := ListEnv()
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, resp, "Response shouldn't be nil")
}

func TestCreateEnv(t *testing.T) {
	Init("http://localhost", "4000")
	config := model.TestchainConfig{
		Type:          "geth",
		BlockMineTime: 0,
		CleanOnStop:   true,
		Accounts:      2,
		DeployRef:     "refs/tags/staxx-testrunner",
		DeployStepID:  0,
	}
	testchain := model.Testchain{

		Config: config,
		Deps:   []string{},
	}
	resp, err := CreateEnv(&testchain)
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, resp, "Response shouldn't be nil")

	envID := resp.Data.(map[string]interface{})["id"].(string)
	// after 1 second
	time.Sleep(5 * time.Second)

	detailsResp, err := EnvDetails(envID)
	assert.Nil(t, err, "Error should be nil")

	respTitle := detailsResp.Data.(map[string]interface{})["testchain"].(map[string]interface{})["title"].(string)
	assert.Equal(t, envID, respTitle, "Testchain title from details shouls be the same with id of the created testchain")

	stopResp, err := StopEnv(envID)
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, stopResp, "Response shouldn't be nil")
	assert.True(t, stopResp.IsResponseStatusOK())

}

func TestRemoveEnv(t *testing.T) {
	Init("http://localhost", "4000")
	config := model.TestchainConfig{
		Type:          "geth",
		BlockMineTime: 0,
		CleanOnStop:   false,
		Accounts:      2,
		DeployRef:     "refs/tags/staxx-testrunner",
		DeployStepID:  0,
	}
	testchain := model.Testchain{

		Config: config,
		Deps:   []string{},
	}
	resp, err := CreateEnv(&testchain)
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, resp, "Response shouldn't be nil")

	envID := resp.Data.(map[string]interface{})["id"].(string)
	// after 1 second
	time.Sleep(5 * time.Second)

	stopResp, err := StopEnv(envID)
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, stopResp, "Response shouldn't be nil")
	assert.True(t, stopResp.IsResponseStatusOK())

	time.Sleep(5 * time.Second)

	removeResp, err := RemoveEnv(envID)
	assert.Nil(t, err, "Error should be nil")

	assert.NotNil(t, removeResp, "Response shouldn't be nil")
	assert.True(t, removeResp.IsResponseStatusOK())

	detailsResp, err := EnvDetails(envID)
	fmt.Printf("f: %+v", detailsResp)
	assert.Nil(t, err, "Error should be nil")
	assert.False(t, detailsResp.IsResponseStatusOK(), "Response status should be false")
	assert.Greater(t, len(detailsResp.Errors), 0, "There should be error for details request")
	assert.Equal(t, "Not Found", detailsResp.Errors[0].Detail)
}
