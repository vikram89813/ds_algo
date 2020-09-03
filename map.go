package main

import (
	"fmt"
)

type test struct {
	a, b string
}

func main() {
	m := make(map[string]int)
	n := make(map[test]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	n[test{"hello", "world"}] = 1

	v1, _ := n[test{"hello", "world"}]
	fmt.Println(v1)

}
