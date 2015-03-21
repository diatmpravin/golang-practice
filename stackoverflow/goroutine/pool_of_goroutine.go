package main

import (
	"os/exec"
	"strconv"
)

func main() {
	NumEI := 8
	for i := 0; i < NumEI; i++ {
		cmd := exec.Command("zenity", "--info", "--text='Hello from iteration n." + strconv.Itoa(i) + "'")
		cmd.Run()
	}
}
