package main

import (
	"os/exec"
	"strconv"
	"sync"
)

func main() {
	NumEl := 8   // Number of times the external program is called
	NumCore := 4 // Number of available cores
	c := make(chan bool, NumCore-1)
	wg := new(sync.WaitGroup)
	wg.Add(NumEl) // Set the number of goroutines to (0 + NumEl)
	for i := 0; i < NumEl; i++ {
		go callProg(i, c, wg)
		c <- true // At the NumCoreth iteration, c is blocking
	}
	wg.Wait() // Wait for all the children to die
	close(c)
}

func callProg(i int, c chan bool, wg *sync.WaitGroup) {
	defer func() {
		<-c
		wg.Done() // Decrease the number of alive goroutines
	}()
	cmd := exec.Command("zenity", "--info", "--text='Hello from iteration n."+strconv.Itoa(i)+"'")
	cmd.Run()
}
