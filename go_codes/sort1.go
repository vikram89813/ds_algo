package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "sdgesbesbsrtab"
	str1 := strings.Split(str, "")
	sort.Strings(str1)
	strings.Join(str1, "")
	fmt.Println(str1)
}
