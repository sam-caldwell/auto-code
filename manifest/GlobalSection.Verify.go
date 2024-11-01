package manifest

// Verify - Verify the global properties section of a manifest.yaml file
func (global GlobalSection) Verify() error {

	vTable := []func() error{

		global.Name.Verify,
		global.Description.Verify,
		global.Author.Verify,
		global.Email.Verify,
		global.Copyright.Verify,
		global.License.Verify,
		global.Language.Verify,
		global.Version.Verify,
		global.GitRepoUrl.Verify,

		//ToDo: If you add anything to the global section, add a function to this table to verify its content.

	}

	for _, f := range vTable {

		if err := f(); err != nil {
			return err
		}

	}

	return nil

}
