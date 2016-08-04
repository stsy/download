package download

import (
	"io/ioutil"
	"net/http"
)

// String returns the sourcecode form URL.
// A successful call returns err == nil.
func String(url string) (html string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	html = string(bytes)
	return
}
