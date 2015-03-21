package main

import (
	"encoding/ascii85"
	"fmt"
)

func main() {

	// encoded message
	src := "9jqo^BlbD-BleB1DJ+*+F(f,q/0JhKF<GL>Cj@.4Gp$d7F!,L7@<6@)/0JDEF<G%<+EV:2F!,\n" +
		"O<DJ+*.@<*K0@<6L(Df-\\0Ec5e;DffZ(EZee.Bl.9pF\"AGXBPCsi+DGm>@3BB/F*&OCAfu2/AKY\n" +
		"i(DIb:@FD,*)+C]U=@3BN#EcYf8ATD3s@q?d$AftVqCh[NqF<G:8+EV:.+Cf>-FD5W8ARlolDIa\n" +
		"l(DId<j@<?3r@:F%a+D58'ATD4$Bl@l3De:,-DJs`8ARoFb/0JMK@qB4^F!,R<AKZ&-DfTqBG%G\n" +
		">uD.RTpAKYo'+CT/5+Cei#DII?(E,9)oF*2M7/c\n"

	newbuffer := make([]byte, len(src))

	_, _, err := ascii85.Decode(newbuffer, []byte(src), true)

	if err != nil {
		fmt.Println(err)
	}

	// print out the decoded message
	fmt.Println(string(newbuffer))

}
