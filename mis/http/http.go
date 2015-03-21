package main

import (
        "io/ioutil"
        "log"
        "net/http"
	"fmt"
)

func main() {
        client := &http.Client{}

        req, err := http.NewRequest("GET", "http://httpbin.org/ip", nil)
        if err != nil {
                log.Fatalln(err)
        }

        req.Header.Set("appName", "demo")


	fmt.Println(req)

        resp, err := client.Do(req)
        if err != nil {
                log.Fatalln(err)
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Fatalln(err)
        }

        log.Println(string(body))

}
