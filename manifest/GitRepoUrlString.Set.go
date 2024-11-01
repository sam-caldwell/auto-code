package manifest

// Set - Set the local git repo url using the given string
func (repo *GitRepoUrlString) Set(value string) error {

	*repo = GitRepoUrlString(value)

	return repo.Verify()

}
