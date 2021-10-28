package infection

import (
	"io/ioutil"
	"log"
)

func GetSelf() []byte {
	var worm []byte
	worm, err := ioutil.ReadFile("/proc/self/exe")
	if err != nil {
		log.Fatal("Unable to read worm binary: %v", err)
	}
	log.Printf("Length: %d",len(worm))
	return worm
}