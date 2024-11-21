package enum

// String - return a string representation of the dataContract contract storage type
func (d *DataDriver) String() string {
	return [...]string{"file", "postgres", "aws-s3"}[*d]
}
