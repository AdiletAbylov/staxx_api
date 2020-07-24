package staxxapi

import (
	"fmt"
	"os"

	"strings"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
)

func TestListSnapshotsFail(t *testing.T) {
	Init("http://localhost", "2000")
	_, err := ListSnapshots("geth")

	assert.NotNil(t, err, "Error should be error")
}

func TestListSnapshots(t *testing.T) {
	Init("http://localhost", "4000")
	resp, _ := ListSnapshots("geth")

	assert.NotNil(t, resp, "Should not be nil")
	assert.Equal(t, true, resp.IsResponseStatusOK(), "response status should be OK")
}

func TestRemoveSnapshot(t *testing.T) {
	Init("http://localhost", "4000")
	resp, err := RemoveSnapshot("someid")

	assert.Equal(t, true, resp.IsResponseStatusOK(), "response status should be OK")
	assert.Nil(t, err, "Error should be nil")
}

func TestUploadSnapshotAndDownload(t *testing.T) {
	Init("http://localhost", "4000")

	resp, err := UploadSnapshot("priv/test_snapshot.tgz", "test snapshot", "geth", printUploadFunction)

	assert.NotNil(t, resp, "Response shouldn't be nil")
	assert.Nil(t, err, "Error should be nil")
	// fmt.Printf("Resp: %+v", resp)
	// fmt.Printf("Err: %+v", err)
	snapshotID := resp.Data.(map[string]interface{})["id"].(string)

	filepathDownload := "priv/" + randomdata.Alphanumeric(5) + ".tgz"
	downerr := DownloadSnapshot(snapshotID, filepathDownload, printDownloadFunction)

	assert.Nil(t, downerr, "Error should be nil")
	_, err = os.Stat(filepathDownload)
	assert.Nil(t, err, "FileInfo should be without error")

	resp, err = RemoveSnapshot(snapshotID)
	assert.Nil(t, err, "Error should be nil")

	//Clean downloaded file
	os.Remove(filepathDownload)
}

func printUploadFunction(bytesWrited uint64, bytesTotal uint64) {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("Uploading %d bytes of %d", bytesWrited, bytesTotal)
}

func printDownloadFunction(bytesWrited uint64, bytesTotal uint64) {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("Downloading %d bytes of %d", bytesWrited, bytesTotal)
}
