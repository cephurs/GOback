package helpers

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(sourcePath, destinationPath string) (int64, error){
	sourceFileStat, err := os.Stat(sourcePath)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", sourcePath)
	}

	source, err := os.Open(sourcePath)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(destinationPath)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	bytesWritten, err := io.Copy(destination, source)
	return bytesWritten, err
}
