package x2j

import (
	"os"
)

func ReadXmlMsgsFromFile(fname string, r bool) error {
	_, statErr := os.Stat(fname)
	if statErr != nil {
		return statErr
	}
	o, oErr := os.Open(fname)
	if oErr != nil {
		return oErr
	}
	defer 
	return nil
}
