package manifest

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// GetLocalRepo - read the git repo url from the local repository
func (manifest *Manifest) GetLocalRepo() error {
	// Execute the "git remote -v" command
	cmd := exec.Command("git", "remote", "-v")
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(errSetGitRepoFailed, err)
	}

	// Split the output by lines and extract the first line
	lines := strings.Split(out.String(), LineEnding)
	if len(lines) == 0 || len(lines[0]) == 0 {
		return fmt.Errorf(errMissingGitRepoUrl) // No output
	}

	// Split the first line by whitespace and return the second column
	fields := strings.Fields(lines[0])
	if len(fields) < 2 {
		return nil // Unexpected format
	}
	return manifest.Global.GitRepoUrl.Set(fields[1])
}
