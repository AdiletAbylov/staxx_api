package helpers

import (
	"net/url"
	"path"
)

// BuildURL returns full url from given service url and array of elements
func BuildURL(connectionString string, elem ...string) string {
	u, _ := url.Parse(connectionString)
	u.Path = path.Join(elem...)
	return u.String()
}
