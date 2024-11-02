package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"os/exec"
)

// Init - initialize local git repository and add a remote origin.
func (repo *GitRepoUrlString) Init() error {

	if s := string(*repo); s == EmptyString {
		return fmt.Fatalf(messages.ErrEmptyGitRepoUrl)
	}

	if err := exec.Command("git", "init").Run(); err != nil {
		return fmt.Fatalf(messages.ErrGitInitFailed, err)
	}

	if err := exec.Command("git", "remote", "add", "origin", string(*repo)).Run(); err != nil {
		return fmt.Fatalf(messages.ErrGitRemoteAddFailed, err)
	}

	return nil

}
