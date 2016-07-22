package basic

import (
	"bytes"
	"strings"
)

func modifyString(str string) (r string) {
	if len(str) == 0 {
		return str
	}

	var b bytes.Buffer

	for _, v := range str {
		if v >= 97 && v <= 122 {
			b.WriteString(strings.ToUpper(string(v)))
		} else {
			b.WriteString(strings.ToLower(string(v)))
		}
	}

	return b.String()
}
