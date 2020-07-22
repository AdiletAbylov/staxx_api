package staxxapi

import (

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSnapshotsFail(t *testing.T) {
	Init("http://localhost", "2000")
	_, err := ListSnapshots("geth")

	assert.NotNil(t, err, "Error should be error")
}

func TestListSnapshots(t *testing.T) {
	Init("http://localhost", "4000")
	resp, _ := ListSnapshots("ganache")

	assert.NotNil(t, resp, "Should not be nil")
	assert.Equal(t, true, resp.IsResponseStatusOK(), "response status should be OK")
}

func TestRemoveSnapshot(t *testing.T) {
	Init("http://localhost", "4000")
	resp, err := RemoveSnapshot("someid")
	
	assert.Equal(t, true, resp.IsResponseStatusOK(), "response status should be OK")
	assert.Nil(t, err, "Error should be nil")
}
