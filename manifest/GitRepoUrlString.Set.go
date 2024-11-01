package manifest

import (
	"fmt"
	"strings"
)

// Set - Set the local git repo url using the given string
func (repo *GitRepoUrlString) Set(value string) error {

	if strings.TrimSpace(value) == EmptyString {
		return fmt.Errorf(errMissingLocalGitRepoUrl)
	}

	*repo = GitRepoUrlString(value)

	return nil

}
