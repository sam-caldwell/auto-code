package gitrepo

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"regexp"
)

// Verify - Validate the git repo url
func (repo *UrlString) Verify() error {

	if pattern := regexp.MustCompile(pattern.GitRepoUrlPattern); pattern.MatchString(string(*repo)) {

		return fmt.Errorf(messages.ErrInvalidGitRepoUrl, *repo)

	}

	return nil

}
