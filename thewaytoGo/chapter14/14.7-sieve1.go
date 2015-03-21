// 14.7-sieve1.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package main
package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel ch.
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send i to channel ch.
	}
}

// Copy the values from channel in to channel out,
// removing those divisible by prime.
func filter(in, out chan int, prime int) {

	for {
		i := <-in
		if i%prime != 0 {

			// Receive value of new variable i from in.
			out <- i // Send i to channel out.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func main() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)
	for {
		// Start generate() as a goroutine.
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1

	}
}
