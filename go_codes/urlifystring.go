package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d", &T)
	parse, _ := regexp.Compile("[ ]")
	for T != 0 {
		str, _ := reader.ReadString('\n')
		str2 := parse.ReplaceAllString(str, "%20")
		fmt.Println(str2)
		T--
	}
}
