package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	version   = "1.0"
	apiKey    = "ea71af1b0d732835af33e7bb1a7b7984ee705169c7c395b79ad4d6a06cd8f246"
	secretKey = "57aeabbba4864e3c5b39f2b2f39c27d1dcfa4a0fef9e8c17bcc67bd0ee815d29"
	host      = "https://paashq.shephertz.com/paas/1.0/app"
)

type Param struct {
	ApiKey    string `json:"apiKey"`
	Version   string `json:"version"`
	TimeStamp string `json:"timeStamp"`
	Host      string `json:"host"`
}

func main() {
	t := time.Now()
	time := t.Format(time.RFC3339)

	p := &Param{
		ApiKey:    apiKey,
		Version:   version,
		Host:      host,
		TimeStamp: time,
	}

	fmt.Println("Params get valued-------------------------->", p)
	fmt.Println(" ")

	params, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Json Encoding Error:", err)
	}

	fmt.Println("Params after Marshal----------------------->", params)
	fmt.Println(" ")

	var pro Param

	_ = json.Unmarshal([]byte(params), &pro)

	fmt.Printf("%#v Params after Unmarshal-------------------->\n", pro)
	fmt.Println(" ")

	for v, i := range pro {
		fmt.Println("Value==>", v)
		fmt.Println("Index==>", i)
	}
}
