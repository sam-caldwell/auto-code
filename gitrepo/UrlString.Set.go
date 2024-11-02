package gitrepo

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/words"
	"strings"
)

// Set - Set the local git repo url using the given string
func (repo *UrlString) Set(value string) error {

	if strings.TrimSpace(value) == words.EmptyString {
		return fmt.Errorf(messages.ErrMissingLocalGitRepoUrl)
	}

	*repo = UrlString(value)

	return nil

}
