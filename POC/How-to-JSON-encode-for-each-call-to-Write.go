package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type onceReader bool

func (r *onceReader) Read(p []byte) (int, error) {
	if *r {
		return 0, io.EOF
	}
	*r = true
	return copy(p, "hello"), nil
}

type jsonEncoder struct {
	io.Reader
}

func (e jsonEncoder) Read(p []byte) (int, error) {
	n, err := e.Reader.Read(p)
	if err != nil {
		return 0, err
	}

	b, err := json.Marshal(string(p[:n]))
	if err != nil {
		return 0, err
	}

	return copy(p, b), nil
}

func main() {
	var r onceReader
	if _, err := io.Copy(os.Stdout, jsonEncoder{&r}); err != nil {
		log.Print(err)
	}
}
