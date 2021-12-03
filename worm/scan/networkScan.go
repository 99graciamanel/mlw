package scan

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"strings"
	"net"
	"bufio"
)

func ScanIp(ip string) [3]int {
	client := http.Client{
		Timeout: 1 * time.Second,
	}	
	ret := [3]int{0,0,0}
	//Apache Scan
	resp, err := client.Get(fmt.Sprintf("http://%s/",ip))
	if err == nil {
		header := resp.Header.Get("Server")
		if strings.Contains(header,"Apache/2.4.50") {
			ret[0] = 80
		}
	}
	//SSH Scan
	timeout,_ := time.ParseDuration("1s")
	conn, err := net.DialTimeout("tcp", ip+":22",timeout)
	if err == nil {
		defer conn.Close()
		message,err := bufio.NewReader(conn).ReadString('\n')
		if err == nil {
			if strings.Contains(message,"SSH") {
				ret[1] = 22
			}
		}
	}
	//Confluence Scan
	resp, err = client.Get(fmt.Sprintf("http://%s:8090/",ip))
	if err == nil {
		header := resp.Header.Get("X-Confluence-Request-Time")
		if header != "" {
			ret[2] = 8090
		}
	}
	return ret
}

func GetRandomIp(baseIp [4]int) string {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		if baseIp[i] == -1 {
			//Obtenir un número entre 0 i 255 si no és l'últim octet 
			baseIp[i] = rand.Intn(255)
			if i == 3 && baseIp[i] == 0 {
				baseIp[i] = 1
			}
		}
	}
	return fmt.Sprintf("%d.%d.%d.%d",baseIp[0],baseIp[1],baseIp[2],baseIp[3])

}

