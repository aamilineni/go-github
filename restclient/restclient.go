package restclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aamilineni/go-github/api/model"
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
		if response.StatusCode == http.StatusNotFound {
			return model.ErrorModel{Status: http.StatusNotFound, Message: "No data found"}
		}
		return model.ErrorModel{
			Status:  response.StatusCode,
			Message: "error while fetching data from github",
		}
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
