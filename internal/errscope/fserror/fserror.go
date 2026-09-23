package fserror

import "io/fs"

func Cause(err error) error {
	if pathErr, ok := err.(*fs.PathError); ok {
		return pathErr.Err
	}
	return err
}
