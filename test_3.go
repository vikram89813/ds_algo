package main

import (
	"fmt"
)

func main() {
	var str string

	str = "abcdef"

	bstr := []byte(str)

	bstr[1] = 'g'

	fmt.Println(string(bstr))
}
