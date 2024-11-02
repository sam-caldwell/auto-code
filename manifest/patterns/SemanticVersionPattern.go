package patterns

const (

	// SemanticVersionPattern - a valid semantic string
	SemanticVersionPattern = `
    	^([0-9]+)\.                 # Matches MAJOR
    	([0-9]+)\.                  # Matches MINOR
    	([0-9]+)                    # Matches PATCH
    	(?:-([0-9A-Za-z-]+          # Optional pre-release label
		(?:\.[0-9A-Za-z-]+)*)?)     # Additional segments
    	(?:\+([0-9A-Za-z-]+         # Optional build metadata
        (?:\.[0-9A-Za-z-]+)*)?)?$   # Additional segments
	` /*end of SemanticVersionPattern*/
)
