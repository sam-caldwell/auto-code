package arguments

// DataSource - An interface to all the supported data sources
type DataSource interface {
	Get() (any, error)
}
