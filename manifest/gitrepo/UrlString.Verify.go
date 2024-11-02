package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"regexp"
)

// Verify - Validate the git repo url
func (repo *GitRepoUrlString) Verify() error {

	if pattern := regexp.MustCompile(gitRepoUrlPattern); pattern.MatchString(string(*repo)) {

		return fmt.Fatalf(messages.ErrInvalidGitRepoUrl, *repo)

	}

	return nil

}
