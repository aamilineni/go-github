package restclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// Get makes the request and gets the http response
func Get(url string, headers http.Header, httpResponse interface{}) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	request.Header = headers
	response, err := Client.Do(request)

	// handle the error
	if err != nil {
		return err
	}

	// handle the status code
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid request")
	}

	defer response.Body.Close()
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBytes, httpResponse)
	if err != nil {
		return err
	}

	return nil
}
