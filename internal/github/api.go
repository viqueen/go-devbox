package github

type Api interface {
	SearchRepositories(query string) ([]Repository, error)
}

type SearchRepositoriesResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}

type Repository struct {
	Id       int    `json:"id"`
	HtmlUrl  string `json:"html_url"`
	Fork     bool   `json:"fork"`
	Language string `json:"language"`
	Archived bool   `json:"archived"`
	Owner    struct {
		Login   string `json:"login"`
		Id      int    `json:"id"`
		HtmlUrl string `json:"html_url"`
	} `json:"owner"`
}
