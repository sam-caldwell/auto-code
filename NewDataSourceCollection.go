package arguments

import (
	"errors"
)

// NewDataSourceCollection - Initialize and return a data source collection with the sources input.
//
// Because golang will almost certainly have to make a copy of an array to a larger contiguous
// memory block, NewDataSourceCollection() allows the caller to pass an expected size so that
// the array is defined with the proper size to start with.
//
// We create a list of pointers to reduce memory and processing overhead when the list is passed
// to functions.  Likewise, we pass the sources list as a pointer to make the copied information
// in the function call faster (64-bits versus n-bytes).
func NewDataSourceCollection(sources *[]*DataSource) (p DataSourceCollection, err error) {

	// nil pointer check
	if sources == nil {
		panic(nilDataSourceListPointer)
	}

	// Make sure we allocate a DataSourceCollection to start out with.
	p = make(DataSourceCollection, 0, len(*sources))
	if len(*sources) == 0 {
		return p, errors.New(emptyDataSourceList)
	}

	// We allocate a list of 0 elements with a capacity of len(sources) to reduce the risk of
	// memory reallocations.  This results in a DataCollection list which is the size of the
	// non-nil elements.
	//
	// Note: calling copy() would be faster, but we need to verify there are no nil pointers
	//       in the data.  If we encounter a nil, we need to return an error.
	for i, ds := range *sources {
		if ds == nil {
			// Note: We do not stop on a nil-element, we just skip that element.
			//       If we stop early, we may cause other errors which could
			//       make it harder to debug issues.
			err = errors.New(nilSourceEncountered)
		} else {
			p[i] = ds
		}
	}

	return p, err
}
