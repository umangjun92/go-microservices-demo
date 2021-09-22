package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions "`
}

type RepoOwner struct {
	Url     string `json:"url"`
	Login   string `json:"login"`
	HtmlUrl string `json:"html_url"`
	Id      int64  `json:"id"`
}

type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPull bool `json:"pull"`
	HasPush bool `json:"push"`
}

// func CreateRepo(){
// 	request := map[string]interface{}{
// 		"name": "Hello World",
// 		"private": false,

// 	}
// }
