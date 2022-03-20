package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aamilineni/go-github/api/model"
	"github.com/aamilineni/go-github/restclient"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var GetDoFunc func(req *http.Request) (*http.Response, error)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func TestGetWithStatus200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	GetDoFunc = func(req *http.Request) (*http.Response, error) {

		var data interface{}

		switch req.URL.String() {
		case "https://mock/branchesURL":
			data = getGithubRepoModel()
		default:
			data = getGithubModel()
		}

		mockResponse := httptest.NewRecorder()
		bytes, _ := json.Marshal(data)
		mockResponse.Write(bytes)
		return mockResponse.Result(), nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "name", Value: "Anil"}}

	restclient.Client = &MockClient{}
	NewGithubHandler(restclient.Client).Get(c)

	assert.Equal(t, c.Writer.Status(), http.StatusOK)
}

func TestGetWithStatus404(t *testing.T) {
	gin.SetMode(gin.TestMode)

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		mockResponse := httptest.NewRecorder()
		mockResponse.WriteHeader(http.StatusNotFound)
		return mockResponse.Result(), nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "name", Value: "Anil"}}

	restclient.Client = &MockClient{}
	NewGithubHandler(restclient.Client).Get(c)

	assert.Equal(t, c.Writer.Status(), http.StatusNotFound)
}

func TestGetWithStatus500(t *testing.T) {
	gin.SetMode(gin.TestMode)

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		mockResponse := httptest.NewRecorder()
		var data interface{}

		switch req.URL.String() {
		case "https://mock/branchesURL":
			mockResponse.WriteHeader(http.StatusInternalServerError)
			return mockResponse.Result(), nil
		default:
			data = getGithubModel()
			bytes, _ := json.Marshal(data)
			mockResponse.Write(bytes)
		}

		return mockResponse.Result(), nil

	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "name", Value: "Anil"}}

	restclient.Client = &MockClient{}
	NewGithubHandler(restclient.Client).Get(c)

	assert.Equal(t, c.Writer.Status(), http.StatusInternalServerError)
}

func TestGetForRepoModelWithStatus500(t *testing.T) {
	gin.SetMode(gin.TestMode)

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		mockResponse := httptest.NewRecorder()
		mockResponse.WriteHeader(http.StatusInternalServerError)
		return mockResponse.Result(), nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "name", Value: "Anil"}}

	restclient.Client = &MockClient{}
	NewGithubHandler(restclient.Client).Get(c)

	assert.Equal(t, c.Writer.Status(), http.StatusInternalServerError)
}

func getGithubModel() *[]model.GithubModel {
	githubModel := []model.GithubModel{}

	for i := 0; i < 2; i++ {
		githubModel = append(githubModel, model.GithubModel{
			ID:       i,
			Name:     fmt.Sprintf("%d", i),
			FullName: "Full name",
			Owner: struct {
				Login    string "json:\"login\""
				ReposURL string "json:\"repos_url\""
			}{
				Login:    "Owner Login",
				ReposURL: "Repos URL",
			},
			BranchesURL: "https://mock/branchesURL{branch}",
		})
	}

	return &githubModel
}

func getGithubRepoModel() *[]model.GithubRepoModel {
	githubRepoModelArr := []model.GithubRepoModel{}

	for i := 0; i < 2; i++ {
		githubRepoModelArr = append(githubRepoModelArr, model.GithubRepoModel{
			Name: fmt.Sprintf("%d", i),
			Commit: struct {
				SHA string "json:\"sha\""
			}{
				SHA: fmt.Sprintf("SHA%d", i),
			},
		})
	}

	return &githubRepoModelArr
}
