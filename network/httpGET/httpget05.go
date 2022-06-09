/* RZFeeser | Alta3 Research 
   HTTP GET and parameters */

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
)

func main() {

    name := "John Doe"
    occupation := "gardener"

    params := "name=" + url.QueryEscape(name) + "&" +
        "occupation=" + url.QueryEscape(occupation)
    path := fmt.Sprintf("https://httpbin.org/get?%s", params)

    resp, err := http.Get(path)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}
/*
{
  "args": {
    "name": "John Doe",
    "occupation": "gardener"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-62a243a4-59949a4831a705b46133572e"
  },
  "origin": "71.251.147.237",
  "url": "https://httpbin.org/get?name=John+Doe&occupation=gardener"
}
*/
