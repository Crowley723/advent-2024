package utils

import (
	"os"
)

// LoadFile takes a filename and returns the two columns of integers as arrays.
func LoadFile[T any](filename string, parse func(*os.File) (T, error)) (T, error) {
	file, err := os.Open(filename)
	if err != nil {
		var zero T
		return zero, err
	}
	defer file.Close()
	return parse(file)

}
