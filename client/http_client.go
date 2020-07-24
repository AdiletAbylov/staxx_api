package client

import (
	"bytes"

	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"

	"github.com/adiletabylov/staxxapi/helpers"
	"github.com/adiletabylov/staxxapi/model"
)

const contentTypeJSON string = "application/json"

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFile(filepath string, url string, printer helpers.ProgressPrinter) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		os.Remove(filepath + ".tmp")
		return err
	}
	defer resp.Body.Close()

	printer.SetTotalLengthFromString(resp.Header.Values("Content-Length")[0])

	if _, err = io.Copy(out, io.TeeReader(resp.Body, &printer)); err != nil {
		out.Close()
		os.Remove(filepath + ".tmp")
		return err
	}

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}

	return nil
}

// UploadFile uploads file from given filepath to the given upload path
// with given params map and using POST request.
func UploadFile(filepath string, params map[string]string, uploadpath string, printer helpers.ProgressPrinter) (*model.Response, error) {
	req, err := newfileUploadRequest(uploadpath, params, "snapshot[file]", filepath, printer)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return model.NewResponseFromBody(resp.Body)
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string, printer helpers.ProgressPrinter) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileInfo, _ := file.Stat()

	printer.Total = uint64(fileInfo.Size())
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, paramName, filepath.Base(path)))
	h.Set("Content-Type", "application/gzip")
	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, err
	}

	// _, err = io.Copy(part, file)
	_, err = io.Copy(part, io.TeeReader(file, &printer))
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

// Get makes GET request by given url and parses and returns Response.
func Get(url string) (*model.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return model.NewResponseFromBody(resp.Body)
}

// Post makes POST request by given url and parses and returns Response.
func Post(url string, data io.Reader) (*model.Response, error) {
	resp, err := http.Post(url, contentTypeJSON, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return model.NewResponseFromBody(resp.Body)
}

// Delete makes DELETE request by given url and parses and returns Response.
func Delete(url string) (*model.Response, error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return model.NewResponseFromBody(resp.Body)
}
