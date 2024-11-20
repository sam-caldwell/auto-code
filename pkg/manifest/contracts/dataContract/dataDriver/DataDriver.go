package dataDriver

// DataDriver - an enumerated type describing the database/storage driver used in the contract.
type DataDriver int

const (
	File DataDriver = iota
	Postgres
	AwsS3
)
