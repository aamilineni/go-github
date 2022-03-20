package restclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aamilineni/go-github/api/model"
	"github.com/stretchr/testify/assert"
)

type TestRestClientResponse struct {
	MetaData string `json:"metadata"`
	Name     string `json:"name"`
}

func TestGetWithStatus200(t *testing.T) {

	mockResp := &TestRestClientResponse{
		MetaData: "metadata",
		Name:     "anil",
	}

	// Create new test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := json.Marshal(mockResp)
		w.Write(bytes)
	}))
	data := &TestRestClientResponse{}
	err := Get(server.URL, nil, data)
	assert.NoError(t, err, "Error while making the Get Request")
	assert.Equal(t, data.MetaData, mockResp.MetaData)
	assert.Equal(t, data.Name, mockResp.Name)
}

func TestGetWithStatus404(t *testing.T) {
	// Create new test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	data := &TestRestClientResponse{}
	err := Get(server.URL, nil, data)
	if errModel, ok := err.(model.ErrorModel); ok {
		assert.Equal(t, errModel.Status, http.StatusNotFound)
		return
	}
	assert.Error(t, err)
}

func TestGetWithStatus500(t *testing.T) {
	// Create new test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	data := &TestRestClientResponse{}
	err := Get(server.URL, nil, data)
	if errModel, ok := err.(model.ErrorModel); ok {
		assert.Equal(t, errModel.Status, http.StatusInternalServerError)
		return
	}
	assert.Error(t, err)
}
