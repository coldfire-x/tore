package libs

import (
	"io/ioutil"
	"net/http"
)

type httplib struct{}

// http request, GET method
func (h *httplib) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", nil
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	return string(content), nil
}
