package main

import "fmt"

type Person map[string]string
type Job map[string]string

func keys(m map[string]string) (keys []string) {
	for key, _ := range m {
		keys = append(keys, key)
	}

	return
}

func name(p Person) string {
	return p["first_name"] + " " + p["last_name"]
}

func main() {
	var person = Person{"first_name": "Rob", "last_name": "Pike"}
	var job = Job{"title": "Commander", "project": "Golang"}

	fmt.Printf("%v\n", name(person))
	//fmt.Printf("%v", name(job)) //this gives a compile error

	fmt.Printf("%v\n", keys(person))
	fmt.Printf("%v\n", keys(job))
}
