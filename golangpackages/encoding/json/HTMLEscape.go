package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	dst := new(bytes.Buffer)

	src := []byte(`{
         "<script>Name":"Adam Ng",  // <----- look here
         "Age":36,
         "Job":"CEO"
         }`)

	json.HTMLEscape(dst, src)

	fmt.Println(dst)
}
