package mapview

import "fmt"

type PathError struct {
	Path 	string			//object path of property in map where error occurred
	Err 	error			//actual error message
}

type PathErrors []PathError

func (errs *PathErrors) GetMapErrors() MapErrors {
	mapErrors = make(MapErrors)
	for _, e := range errs {
		append(mapErrors[e.Path()], e)
	}
	
	return mapErrors
}

// map of string property paths to arrays of errors for that path
type MapErrors map[string][]error

// Define some specific error types: TODO: error messages

// path is required, but not present
type MissingPathError PathError {}
func (e MissingPathError) Error() string {
	
}

// path is of a type that is not allowed
type UnsupportedTypeError PathError {}
func (e UnsupportedTypeError) Error() string {
	
}

// path is present, but should not be
type ExtantPathError PathError {}
func (e UnsupportedTypeError) Error() string {
	
}
