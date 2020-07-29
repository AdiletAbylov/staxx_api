package staxxapi

import (

	"github.com/adiletabylov/staxxapi/client"
	"github.com/adiletabylov/staxxapi/helpers"
	"github.com/adiletabylov/staxxapi/model"
)

// CreateEnv creates and starts new environment with given testchain config
func CreateEnv(testchain *model.Testchain) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "environments", "start")
	json, err := testchain.ToReader()
	if err != nil {
		return nil, err
	}
	return client.Post(url, json)
}

// StopEnv stops running environment with given config
func StopEnv(envID string) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "environments", envID, "stop")
	return client.Get(url)
}

// EnvDetails return details of environment with the given id
func EnvDetails(envID string) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "environments", envID)
	return client.Get(url)
}

// ListEnv returns list of all environments
func ListEnv() (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "environments")
	return client.Get(url)
}

// RemoveEnv removes environment with the given id
func RemoveEnv(envID string) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "environments", envID)
	return client.Delete(url)
}

// TakeSnapshot creates snapshot of environment with given id
func TakeSnapshot(envID string, description string) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "environments", envID, "take_snapshot")
	data, err := model.DataForTakingSnapshotRequest(envID, description)
	if err != nil {
		return nil, err
	}
	return client.Post(url, data)
}

// RevertSnapshot reverts snapshot with the given id and environment ID
func RevertSnapshot(envID string, snapshotID string) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "environments", envID, "revert_snapshot")
	data, err := model.DataForTakingSnapshotRequest(envID, snapshotID)
	if err != nil {
		return nil, err
	}
	return client.Post(url, data)
}
