package httpContract

// HttpEndpointVersion - represents a httpContract endpoint version.  Used for versioning APIs.
//
// This should be a single unsigned integer value which will be returned by .String() as
// the version string 'v1', 'v2', etc.
type HttpEndpointVersion uint
