package manifest

func (l *LicenseName) String() string {
	return [...]string{"Proprietary", "MIT", "Apache2", "BSD"}[*l]
}
