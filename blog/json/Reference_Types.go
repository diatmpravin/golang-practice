package main

import (
	"encoding/json"
	"fmt"
)

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

// type Bar struct {
// 	Name string
// }

// type Foo struct {
//     Bar *Bar
// }

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var m FamilyMember
	fmt.Printf("%+v\n",m)

	// var f Foo
	// fmt.Printf("%+v\n",f)

	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n",f)
}
