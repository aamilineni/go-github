package handlers

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/aamilineni/go-github/api/model"
	"github.com/aamilineni/go-github/constants"
	"github.com/aamilineni/go-github/restclient"
	"github.com/gin-gonic/gin"
)

type githubHandler struct {
	client restclient.HTTPClient
}

func NewGithubHandler(client restclient.HTTPClient) *githubHandler {
	return &githubHandler{
		client: client,
	}
}

func (me *githubHandler) Get(ctx *gin.Context) {

	// get the repo name from query param
	name := ctx.Param(constants.QUERY_PARAM_NAME)

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Basic %s", os.Getenv(constants.GITHUB_AUTHTOKEN)))

	// Get the REPO'S information for the given user from github API
	reposURL := fmt.Sprintf(constants.GET_REPOS_URL, name)
	githubModel := &[]model.GithubModel{}
	err := restclient.Get(reposURL, headers, githubModel)
	if err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))

		return
	}

	var wg sync.WaitGroup

	githubRepoResponses := []*model.GithubRepoResponse{}

	for _, repo := range *githubModel {
		wg.Add(1)
		go func(repoModel model.GithubModel) {
			defer wg.Done()
			githubRepoModelArr := &[]model.GithubRepoModel{}
			err := restclient.Get(repoModel.GetBranchesURL(), headers, githubRepoModelArr)
			if err != nil {
				ctx.Errors = append(ctx.Errors, ctx.Error(err))

				return
			}

			githubRepoResponse := &model.GithubRepoResponse{
				OwnerLogin: repoModel.Owner.Login,
				RepoName:   repoModel.Name,
				Branches:   []model.GithubRepoBranchResponse{},
			}

			for _, githubRepoModel := range *githubRepoModelArr {
				githubRepoResponse.Branches = append(githubRepoResponse.Branches, model.GithubRepoBranchResponse{
					LastCommitSHA: githubRepoModel.Commit.SHA,
					BranchName:    githubRepoModel.Name,
				})
			}
			githubRepoResponses = append(githubRepoResponses, githubRepoResponse)

		}(repo)

	}
	wg.Wait()

	ctx.JSON(http.StatusOK, &githubRepoResponses)
}
