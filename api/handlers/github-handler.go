package handlers

import (
	"fmt"
	"sync"

	"github.com/aamilineni/go-github/constants"
	"github.com/aamilineni/go-github/restclient"
	"github.com/gin-gonic/gin"
)

type GithubModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login    string `json:"login"`
		ReposURL string `json:"repos_url"`
	} `json:"owner"`
	BranchesURL string `json:"branches_url"`
}

type GithubRepoModel struct {
	Name  string `json:"name"`
	Owner struct {
		SHA string `json:"sha"`
	} `json:"commit"`
}

type GithubRepoBranchResponse struct {
	LastCommitSHA string `json:"lastCommitSHA"`
	BranchName    string `json:"branchName"`
}

type GithubRepoResponse struct {
	RepoName   string                     `json:"repoName"`
	OwnerLogin string                     `json:"ownerLogin"`
	Branches   []GithubRepoBranchResponse `json:"branches"`
}

type githubHandler struct {
	client restclient.HTTPClient
}

func NewGithubHandler(client restclient.HTTPClient) *githubHandler {
	return &githubHandler{
		client: client,
	}
}

func (me *githubHandler) Get(ctx *gin.Context) {
	name := ctx.Param(constants.QUERY_PARAM_NAME)

	reposURL := fmt.Sprintf(constants.GET_REPOS_URL, name)
	githubModel := []GithubModel{}
	err := restclient.Get(reposURL, nil, githubModel)
	if err != nil {
		ctx.Errors = append(ctx.Errors, ctx.Error(err))

		return
	}

	var wg sync.WaitGroup

	githubRepoResponses := []*GithubRepoResponse{}

	for _, repo := range githubModel {

		go func(repoModel GithubModel) {
			defer wg.Done()
			githubRepoModelArr := []GithubRepoModel{}
			err := restclient.Get(repoModel.BranchesURL, nil, githubRepoModelArr)
			if err != nil {
				ctx.Errors = append(ctx.Errors, ctx.Error(err))

				return
			}

			githubRepoResponse := &GithubRepoResponse{
				OwnerLogin: repoModel.Owner.Login,
				RepoName:   repoModel.Name,
				Branches:   []GithubRepoBranchResponse{},
			}

			for _, githubRepoModel := range githubRepoModelArr {
				githubRepoResponse.Branches = append(githubRepoResponse.Branches, GithubRepoBranchResponse{
					LastCommitSHA: githubRepoModel.Owner.SHA,
					BranchName:    githubRepoModel.Name,
				})
			}

		}(repo)

	}
	wg.Wait()

	ctx.BindJSON(githubRepoResponses)
}
