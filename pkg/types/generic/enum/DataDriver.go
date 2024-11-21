package enum

// DataDriver - an enumerated type describing the database/storage driver used in the contract.
type DataDriver uint8

const (
	File DataDriver = iota
	Postgres
	AwsS3
)
