package gitrepo

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"strings"
)

// Set - Set the local git repo url using the given string
func (repo *UrlString) Set(value string) error {

	if strings.TrimSpace(value) == manifest.EmptyString {
		return fmt.Errorf(messages.ErrMissingLocalGitRepoUrl)
	}

	*repo = UrlString(value)

	return nil

}
