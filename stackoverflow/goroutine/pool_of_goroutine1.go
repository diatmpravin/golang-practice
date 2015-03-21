package main

import (
	"os/exec"
	"strconv"
)

func main() {
	NumEl := 8
	for i := 0; i < NumEl; i++ {
		go callProg(i) // <--- There!
	}
}

func callProg(i int) {
	cmd := exec.Command("zenity", "--info", "--text='Hello from iteration n."+strconv.Itoa(i)+"'")
	cmd.Run()
}
