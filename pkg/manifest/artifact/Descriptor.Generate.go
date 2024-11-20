package artifact

import "github.com/sam-caldwell/ansi"

// Generate - Generate the project artifact
func (a *Descriptor) Generate(target string, debug bool) error {
	if debug {
		ansi.Blue().Printf("Generate a %s: %s", a.Type.String(), a.Name.String())
	}
	switch a.Type {
	case Service:
	case External:
	case Binary:
	default:
		ansi.Red().Printf("unrecognized artifact type:'%s'" + a.Type.String()).LF().Fatal(1)
	}
	return nil
}
