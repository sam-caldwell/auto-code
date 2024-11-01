package manifest

import (
	"fmt"
	"os/exec"
)

// Init - initialize local git repository and add a remote origin.
func (repo *GitRepoUrlString) Init() error {

	if s := string(*repo); s == EmptyString {
		return fmt.Errorf(errMissingGitRepoUrl)
	}

	if err := exec.Command("git", "init").Run(); err != nil {
		return err
	}

	if err := exec.Command("git", "remote", "add", "origin", string(*repo)).Run(); err != nil {
		return err
	}

	return nil

}
