package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "1,2,3,4,5,6"

	arr := strings.Split(str, ",")

	for _, v := range arr {
		fmt.Println(v)
	}
}
