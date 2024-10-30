package arguments

// dsYamlFile - A DataSource
type dsYamlFile struct{}

func (ds *dsYamlFile) Get() (any, error) {
	return nil, nil
}
func YamlFile(name, help string, validator *Validator) DataSource {
	return dsCliOption{}
}
