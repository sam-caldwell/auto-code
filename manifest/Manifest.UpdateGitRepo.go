package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/gitrepo"
	"github.com/sam-caldwell/auto-code/manifest/messages"
)

// UpdateGitRepo - merge the manifest.Global.git_repo and local git repo and ensure agreement
//
// We are going to merge global.git_repo from the manifest YAML file with the local git repo
// and verify the two agree. In an optimal state
func (manifest *Manifest) UpdateGitRepo() (err error) {

	// get the local git repo url
	var localRepo gitrepo.UrlString
	if localRepo, err = manifest.gitLocalRepo(); err != nil {
		// if the manifest does not have a git_repo url, we cannot go forward.
		if manifest.Global.GitRepoUrl == EmptyString {
			// the local git repo could not be loaded or was not defined.
			return fmt.Errorf(messages.ErrMissingGitRepoUrl, err)
		}

		// if the manifest has a git_repo url, we can set our local repo to that value.
		if err = manifest.Global.GitRepoUrl.Init(); err != nil {
			// on error, bail.  We cannot fix the local repo, so there is an unrecoverable problem.
			return err
		}

		// now we can load the local repo again to verify our errors are resolved.
		if localRepo, err = manifest.gitLocalRepo(); err != nil {
			return err // We have proven we have an unrecoverable error
		}
	} else {
		// localRepo was loaded without error.
		if manifest.Global.GitRepoUrl == EmptyString {
			// the local git repo could not be loaded or was not defined.
			return fmt.Errorf(messages.ErrMissingGitRepoUrl, err)
		} else {
			// we had no manifest.Global.git_repo in the manifest, but we had a localRepo
			// so we will set our manifest version to the localRepo value.
			manifest.Global.GitRepoUrl = localRepo
		}
	}

	// Check final agreement to ensure local and manifest git repo are the same.
	if localRepo != manifest.Global.GitRepoUrl {
		return fmt.Errorf(messages.ErrGitRepoUrlMismatch)
	}

	return err

}
