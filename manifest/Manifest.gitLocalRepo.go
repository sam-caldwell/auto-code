package manifest

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// gitLocalRepo - execute 'git remote -v' to fetch the local repo url
func (manifest *Manifest) gitLocalRepo() (localRepo GitRepoUrlString, err error) {

	// Execute the "git remote -v" command
	cmd := exec.Command("git", "remote", "-v")
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		return localRepo, fmt.Errorf(errGitRemoteUrlFailed, err)
	}

	// Split the output by lines and extract the first line
	lines := strings.Split(out.String(), LineEnding)
	if len(lines) == 0 || len(lines[0]) == 0 {
		return localRepo, fmt.Errorf(errMissingLocalGitRepoUrl) // No output
	}

	// Split the first line by whitespace and return the second column
	if fields := strings.Fields(lines[0]); len(fields) < 2 {
		return localRepo, nil // Unexpected format
	} else {
		err = localRepo.Set(fields[1])
	}

	return localRepo, nil
}
