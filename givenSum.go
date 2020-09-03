package main

import "fmt"

func main() {
	arr := []int{1, 4, 45, 6, 10, 8}
	sum := 16

	fn := func(arr []int, sum int) (int, int) {
		m := make(map[int]int)
		a := -1
		b := -1

		for _, v := range arr {
			s := sum - v
			if _, ok := m[s]; ok {
				a = v
				b = s
				break
			} else {
				m[v] = v
			}
		}
		return a, b
	}

	a, b := fn(arr, sum)
	fmt.Println(a, b)
}
