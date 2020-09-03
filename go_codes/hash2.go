package main

import "fmt"

func main() {
	m := make(map[int]bool)

	arr := []int{1, 2, 3, 4, 5, 6, 2, 4, 3}
	var ans int
	ans = 0

	for _, v := range arr {
		if _, ok := m[v]; ok {
			ans++
		} else {
			m[v] = true
		}
	}
	fmt.Println(ans)
}
