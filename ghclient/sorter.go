package ghclient

func (reposList ReposList) Len() int {
	return len(reposList.Repos)
}

func (reposList ReposList) Less(i, j int) bool {
	return reposList.Repos[i].Name < reposList.Repos[j].Name
}

func (reposList ReposList) Swap(i, j int) {
	reposList.Repos[i], reposList.Repos[j] = reposList.Repos[j], reposList.Repos[i]
}
