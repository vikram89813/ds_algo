package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var T int64
	fmt.Scanf("%d", &T)
	reader := bufio.NewReader(os.Stdin)

	for T != 0 {
		str, _ := reader.ReadString('\n')
		fmt.Println(strings.Title(str))
		T--
	}
}
