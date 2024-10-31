package arguments

// dsYamlFile - A DataSource
type dsYamlFile struct{}

func (ds *dsYamlFile) Get() (any, error) {
	return nil, nil
}

func YamlFile(fileName, dataObjectName string) DataSource {
	return dsCliOption{}
}
