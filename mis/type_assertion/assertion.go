package main

import (
	"fmt"
)

func main() {
	var anything interface{} = "string"
	var moreAnything interface{} = 123

	fmt.Println(anything, moreAnything)

	// You can’t use a ‘type conversion’ here since they’re not similar types. You must use a type assertion
	//aString := anything.(string)

	//We convert the ‘anything’ interface{} type using a dot and the required type in parentheses.
	//If it fails, golang will panic and fail (unless you ‘recover’ from the panic). You can get around this by passing back two parameters:
	//aString, found := anything.(string)
	// Now if the assertion fails, it won’t panic and crash, but will set found to false.

	// If you’re not sure about what the interface{} type could be, you can use a type switch:

	switch v := anything.(type) {
	case string:
		fmt.Println(v)
	case int32, int64:
		fmt.Println(v)
	//case SomeCustomType:
	//	fmt.Println(v)
	default:
		fmt.Println("unknown")
	}

}
