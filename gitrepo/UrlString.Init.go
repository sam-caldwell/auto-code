package gitrepo

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/words"
	"os/exec"
)

// Init - initialize local git repository and add a remote origin.
func (repo *UrlString) Init() error {

	if s := string(*repo); s == words.EmptyString {
		return fmt.Errorf(messages.ErrEmptyGitRepoUrl)
	}

	if err := exec.Command("git", "init").Run(); err != nil {
		return fmt.Errorf(messages.ErrGitInitFailed, err)
	}

	if err := exec.Command("git", "remote", "add", "origin", string(*repo)).Run(); err != nil {
		return fmt.Errorf(messages.ErrGitRemoteAddFailed, err)
	}

	return nil

}
