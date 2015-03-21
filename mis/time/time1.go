package main

import "fmt"
import "time"

func main() {
	t := time.Now().UTC()
	fmt.Println(t)
	round := []time.Duration{

		time.Millisecond,
	}

	for _, d := range round {
                fmt.Println(d)
		fmt.Println(t.Round(d).Format("15:04:05.999999999"))
	}
}
