package main

import (
	"fmt"
	"sort"
)

type test struct {
	x, y int
}

func main() {
	arr := make([]test, 0)
	arr = append(arr, test{3, 1})
	arr = append(arr, test{1, 4})
	arr = append(arr, test{0, 6})
	arr = append(arr, test{5, 9})

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].x < arr[j].x
	})

	fmt.Println(arr)
}
