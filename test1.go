package main

import (
	"fmt"
	"runtime"
)

func main() {
	t := runtime.NumCPU()
	fmt.Println(t)
}
