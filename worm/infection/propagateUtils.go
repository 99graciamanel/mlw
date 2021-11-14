package infection

import (
	"io/ioutil"
	"log"
)

func GetFile(filename string) []byte {
	var file []byte
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Unable to read binary: %v", err)
	}
	return file
}