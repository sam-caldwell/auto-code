package arguments

type dsJsonFile struct{}

func (ds dsJsonFile) Get() (any, error) {
	return nil, nil
}
func JsonFile(fileName, recordName string) DataSource {
	return dsJsonFile{}
}
