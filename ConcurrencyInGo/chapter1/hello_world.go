package main

import (
	"fmt"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
}

func main() {
	hello := new(Job)
	world := new(Job)

	hello.i = 0
	hello.max = 3
	hello.text = "hello"

	world.i = 0
	world.max = 5000
	world.text = "world"

	go outputText(hello)
	outputText(world)
}
