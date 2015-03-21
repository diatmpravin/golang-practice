package main

import (
	"os/exec"
	"strconv"
)

func main() {
	NumEl := 8   // Number of times the external program is called
	NumCore := 4 // Number of available cores
	c := make(chan bool, NumCore-1)
	for i := 0; i < NumEl; i++ {
		go callProg(i, c)
		c <- true // At the NumCoreth iteration, c is blocking
	}
}

func callProg(i int, c chan bool) {
	defer func() { <-c }()
	cmd := exec.Command("zenity", "--info", "--text='Hello from iteration n."+strconv.Itoa(i)+"'")
	cmd.Run()
}
