package manifest

// DocumentReference - a path/file or url reference to an external document.
//
// The DocumentReference represents the $ref used in a document which includes
// another document.  A $ref will cause the external source to be identified
// then loaded, parsed and fed to the containing document.
//
// This object will provide the load/parse functionality, but the wrapping object
// must provide any merge functionality to take the parsed document and include it
// into its content.
type DocumentReference string
