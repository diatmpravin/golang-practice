package main

import (
	"encoding/ascii85"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	var str string
	str = "9jqo^BlbD-BleB1DJ+*+F(f,q/0JhKF<GL>Cj@.4Gp$d7F!,L7@<6@)/0JDEF<G%<+EV:2F!,\n"

	decoderstream := ascii85.NewDecoder(strings.NewReader(str))

	dstreambuf, err := ioutil.ReadAll(decoderstream)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(dstreambuf))

}
