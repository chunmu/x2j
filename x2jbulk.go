package x2j

import (
	"os"
)

func ReadXmlMsgsFromFile(fname string) error {
	println("get in")
	_, statErr := os.Stat(fname)
	if statErr != nil {
		return statErr
	}
	xmlFile, openErr := os.Open(fname)
	if openErr != nil {
		return openErr
	}
	defer xmlFile.Close()
	return nil
}
