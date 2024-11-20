package license

func (l *Name) String() string {
	return [...]string{"Proprietary", "MIT", "Apache2", "BSD"}[*l]
}
