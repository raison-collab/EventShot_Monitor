package main

import "fmt"

type EmptyDirError struct {
	Message string
}

func (e EmptyDirError) Error() string {
	return fmt.Sprintf("EmtyDirError: %s", e.Message)
}
