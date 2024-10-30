package arguments

// Count - Return the number of data sources
func (d *DataSourceCollection) Count() int {
	return len(*d)
}
