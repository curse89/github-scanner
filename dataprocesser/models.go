package dataprocesser

type DataResult struct {
	account string
	repos   []RepoResult
	err     error
}

type RepoResult struct {
	name        string
	versions    []string
	versionsErr error
}
