package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"strings"
)

// Set - Set the local git repo url using the given string
func (repo *GitRepoUrlString) Set(value string) error {

	if strings.TrimSpace(value) == EmptyString {
		return fmt.Fatalf(messages.ErrMissingLocalGitRepoUrl)
	}

	*repo = GitRepoUrlString(value)

	return nil

}
