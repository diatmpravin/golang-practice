package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type test_struct struct {
	Test string
}

// func test(rw http.ResponseWriter, req *http.Request) {
// 	req.ParseForm()
// 	log.Println(req.Form)
// 	//LOG: map[{"test": "that"}:[]]
// 	var t test_struct
// 	for key, _ := range req.Form {
// 		log.Println(key)
// 		//LOG: {"test": "that"}
// 		err := json.Unmarshal([]byte(key), &t)
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 	}
// 	log.Println(t.Test)
// 	//LOG: that
// }

func test(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(t.Test)
}

// func test(rw http.ResponseWriter, req *http.Request) {
// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		panic()
// 	}
// 	log.Println(string(body))
// 	var t test_struct
// 	err = json.Unmarshal(body, &t)
// 	if err != nil {
// 		panic()
// 	}
// 	log.Println(t.Test)
// }

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
