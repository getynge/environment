// Package filter provides facilities for filtering and modifying environment variables
package filter

type Filter interface {
    // Filter filters a given key value pair.
    // The return value of the function is what the provided key value pair will be replaced with upon the function's
    // return. If the function returns a non-nil error, the other return values are discarded.
    Filter(keyIn, valueIn string) (keyOut, valueOut string, err error)
}