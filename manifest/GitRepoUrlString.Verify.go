package manifest

import (
	"fmt"
	"regexp"
)

// Verify - Validate the git repo url
func (repo *GitRepoUrlString) Verify() error {

	if pattern := regexp.MustCompile(gitRepoUrlPattern); pattern.MatchString(string(*repo)) {

		return fmt.Errorf(errInvalidGitRepoUrl, *repo)

	}

	return nil

}
