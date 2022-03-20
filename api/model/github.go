package model

import (
	"strings"
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

func (me GithubModel) GetBranchesURL() string {
	return me.BranchesURL[:strings.Index(me.BranchesURL, "{")]
}

type GithubRepoModel struct {
	Name   string `json:"name"`
	Commit struct {
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
