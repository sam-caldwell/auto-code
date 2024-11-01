package manifest

const (
	// globalCopyrightPattern - copyright string pattern
	//
	// All copyrights must follow the US pattern:
	//   (c) YEAR Author. All Rights Reserved.
	// or
	//   ©YEAR Author. All Rights Reserved.
	globalCopyrightPattern = `^©\s?(\d{4}(-\d{4})?)(,\s?\d{4}(-\d{4})?)*\s[A-Za-z .,'-]+$`

	// NameIdentifierPattern - a valid name identifier (must be alphanumeric plus _ or .)
	NameIdentifierPattern = `[a-zA-Z][a-zA-Z0-9_\\.]{1,256}`

	// globalEmailPattern - a valid email address for the program project / author.
	globalEmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$\n`

	// globalLanguagePattern - a list of supported languages
	globalLanguagePattern = `^(go|c++|node)$`

	// globalLicensePattern - a list of supported licenses
	globalLicensePattern = `^(MIT|BSD|Apache2|Proprietary)$`

	// globalVersionPattern - a valid semantic string
	globalVersionPattern = `
    	^([0-9]+)\.                 # Matches MAJOR
    	([0-9]+)\.                  # Matches MINOR
    	([0-9]+)                    # Matches PATCH
    	(?:-([0-9A-Za-z-]+          # Optional pre-release label
		(?:\.[0-9A-Za-z-]+)*)?)     # Additional segments
    	(?:\+([0-9A-Za-z-]+         # Optional build metadata
        (?:\.[0-9A-Za-z-]+)*)?)?$   # Additional segments
	` /*end of globalVersionPattern*/
)

const (
	// configSourcesPattern - a supported config.sources value
	configSourcesPattern = `(file|environment|commandline)`

	// configPropertyTypes - data type for the config properties
	configPropertyTypes = `
		bool,
		int,	int8,	int16,	int32,	int64,
		uint,	uint8,	uint16,	uint32,	uint64,
		float32,float64,
		string
	`
	environmentVariableNamePattern = `^[a-zA-Z][a-zA-Z0-9_]{0-64}$`

	fileNameStringPattern = `^(/[^/ ]+)+/[^/ ]+\.[a-zA-Z0-9]+$`

	fileFormatPattern = `^(yaml|json)$`

	objectPropertyStringPattern = `^(\.?([a-zA-Z_]\w*|\["[^"]+"\]))+(\[\d*\])*(\.\*|(\|.+))*$`

	shortArgumentPattern = `^-[a-zA-Z0-9]$`

	longArgumentPattern = `^--([a-zA-Z0-9][a-zA-Z0-9_\-\\.]{0,254)[a-zA-Z0-9]{0,1}$`
)
