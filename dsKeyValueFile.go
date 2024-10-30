package arguments

type dsKeyValueFile struct{}

func (ds dsKeyValueFile) Get() (any, error) {
	return nil, nil
}

func KeyValueFile(fileName, recordName string) DataSource {
	return dsKeyValueFile{}
}
