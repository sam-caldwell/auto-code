package enum

// Language - represents the language auto-code will use to generate an artifact
//
// Supported languages:
//   - golang
//   - typescript-react
//   - sql
//
// We may add more languages over time
type Language int

const (
	Golang Language = iota
	TypeScriptReact
	Sql
)
