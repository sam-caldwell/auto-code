package patterns

const (

	// ObjectPropertyStringPattern - a dot-delimited object string pattern for looking up a property in YAML/JSON
	ObjectPropertyStringPattern = `^(\.?([a-zA-Z_]\w*|\["[^"]+"\]))+(\[\d*\])*(\.\*|(\|.+))*$`
)
