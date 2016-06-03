package reverse

import (
	"strings"
)

func firstReverse(str string) string {
	newStr := make([]string, len(str))
	arry := strings.Split(str, "")
	
	for i, c := range arry{
		newStr[len(arry)-(i+1)] = c
	}

	return strings.Join(newStr, "")
}

