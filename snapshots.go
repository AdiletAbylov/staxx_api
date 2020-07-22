package staxxapi

import (
	"github.com/adiletabylov/staxxapi/client"
	"github.com/adiletabylov/staxxapi/helpers"
	"github.com/adiletabylov/staxxapi/model"
)

// DownloadSnapshot downloads snapshot and saves it to the given filepath
func DownloadSnapshot(snapshotID string, filepath string, printFunc func(bytesWrited uint64, bytesTotal uint64)) error {
	url := helpers.BuildURL(connectionString(), "snapshots", snapshotID, "download")
	printer := helpers.NewProgressPrinter(printFunc)
	return client.DownloadFile(filepath, url, printer)

}

// UploadSnapshot uploads snapshot file
func UploadSnapshot(filepath string, description string, chainType string, printFunc func(bytesWrited uint64, bytesTotal uint64)) error {
	params := map[string]string{
		"snapshot[description]": description,
		"snapshot[type]":        chainType,
	}
	printer := helpers.NewProgressPrinter(printFunc)
	url := helpers.BuildURL(connectionString(), "snapshots")
	return client.UploadFile(filepath, params, url, printer)
}

// RemoveSnapshot makes DELETE request to remove snapshot with given id
func RemoveSnapshot(snapshotID string) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "snapshots", snapshotID)
	resp, err := client.Delete(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ListSnapshots makes GET request and returns list of snapshots of given evm type
func ListSnapshots(evmType string) (*model.Response, error) {
	url := helpers.BuildURL(connectionString(), "snapshots", evmType)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}
