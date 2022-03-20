package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aamilineni/go-github/api/model"
	"github.com/aamilineni/go-github/constants"
	"github.com/aamilineni/go-github/restclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
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
	headers.Add("Authorization", fmt.Sprintf("token %s", os.Getenv(constants.GITHUB_AUTHTOKEN)))

	// Get the REPO'S information for the given user from github API
	reposURL := fmt.Sprintf(constants.GET_REPOS_URL, name)
	githubModel := &[]model.GithubModel{}
	err := restclient.Get(reposURL, headers, githubModel)
	if err != nil {
		handleError(err, ctx)

		return
	}

	var errGroup errgroup.Group

	githubRepoResponses := []*model.GithubRepoResponse{}

	for _, repo := range *githubModel {
		func(repoModel model.GithubModel) {
			errGroup.Go(func() error {
				githubRepoModelArr := &[]model.GithubRepoModel{}
				err := restclient.Get(repoModel.GetBranchesURL(), headers, githubRepoModelArr)
				if err != nil {
					return err
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
				return nil
			})

		}(repo)

	}

	if err := errGroup.Wait(); err != nil {
		handleError(err, ctx)

		return
	}

	ctx.JSON(http.StatusOK, &githubRepoResponses)
}

func handleError(err error, ctx *gin.Context) {

	switch v := err.(type) {
	case model.ErrorModel:
		ctx.JSON(v.Status, v)
	default:
		ctx.JSON(http.StatusInternalServerError, err)
	}

}
