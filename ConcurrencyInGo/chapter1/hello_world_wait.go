package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job, g *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	g.Done()
}

func main() {
	goGroup := new(sync.WaitGroup)
	fmt.Println("Starting....")

	hello := new(Job)
	world := new(Job)

	hello.i = 0
	hello.max = 3
	hello.text = "hello"

	world.i = 0
	world.max = 5
	world.text = "world"

	go outputText(hello, goGroup)
	go outputText(world, goGroup)

	goGroup.Add(1)
	goGroup.Wait()
}
