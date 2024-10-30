package arguments

type dsEnvironment struct{}

func (ds dsEnvironment) Get() (any, error) {
	return nil, nil
}
func Environment(name string, eName string, validator *Validator) DataSource {
	return dsEnvironment{}
}
