package main

import (
	"os/exec"
	"strconv"
	"sync"
	"fmt"
)

func main() {
	tasks := make(chan *exec.Cmd, 64)
	fmt.Printf("%+v", tasks)

	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for cmd := range tasks {
				cmd.Run()
			}
			wg.Done()
		}()
	}

	// generate some tasks
	for i := 0; i < 10; i++ {
		tasks <- exec.Command("zenity", "--info", "--text='Hello from iteration n."+strconv.Itoa(i)+"'")
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}
