package main

import (
	"fmt"
	"regexp"
)

func isMatch(s string, p string) bool {
	parse, _ := regexp.Compile(p)
	ans := parse.MatchString(s)
	return ans
}
func main() {
	s := "aa"
	p := "a"

	fmt.Println(isMatch(s, p))
}
