package arguments

// DataSourceCollection - A collection of data source objects (tracked by reference).
//
// This is a list of DataSource pointers, which must preserve the order in which they were registered
// with the arguments.Parser object via Parameter().  This must also ensure that no DataSource pointer
// is nil.  Nil pointers in the collection would be illogical and dangerous.
type DataSourceCollection []*DataSource
