package manifest

import (
	"fmt"
)

// UpdateGitRepo - merge the manifest.Global.git_repo and local git repo and ensure agreement
func (manifest *Manifest) UpdateGitRepo() (err error) {

	var localRepo GitRepoUrlString
	if localRepo, err = gitLocalRepo(); err != nil {

		if manifest.Global.GitRepoUrl == EmptyString {

			// We have no means of recovery.  There is no local repo and no repo in the manifest yaml file.
			return fmt.Errorf(errMissingGitRepoUrl)

		} else {

			// Local repo is not defined.  But the manifest yaml defines a git repo.
			// We can initialize the local git repo at this point.
			if err := manifest.Global.GitRepoUrl.Init(); err != nil {
				return err
			}

			// Verify the local repo initialization was correct.
			if localRepo, err = gitLocalRepo(); err != nil {
				return fmt.Errorf(errLocalRepoInitFailed, err)
			}

		}
		return err
	}

	if repo := string(manifest.Global.GitRepoUrl); repo == EmptyString {
		//repo is an empty string, we will need to set the value to the local repo.
	}

}
