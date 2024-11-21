package archive

// HttpEndpointMethod - represents an HTTP method (e.g. GET, PUT, POST, DELETE, etc)
type HttpEndpointMethod int

const (
	GET HttpEndpointMethod = iota
	POST
	PUT
	DELETE
)
