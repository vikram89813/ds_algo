package main

import (
	"fmt"
	"sync"
)

func main() {
	gap := func(gap int) int {
		if gap <= 1 {
			return 0
		}

		return (gap / 2) + (gap % 2)
	}

	swap := func(a int, b int) (int, int) {
		return b, a
	}

	arr1 := []int{10, 27, 38, 43, 82}
	arr2 := []int{3, 9}

	merge := func(arr1 []int, arr2 []int, wg *sync.WaitGroup) {
		defer wg.Done()
		n := len(arr1)
		m := len(arr2)

		for g := gap(n + m); g > 0; g = gap(g) {
			i := 0
			for ; i+g < n; i++ {
				if arr1[i] > arr1[i+g] {
					arr1[i], arr1[i+g] = swap(arr1[i], arr1[i+g])
				}
			}

			j := 0
			if g > n {
				j = g - n
			}

			for ; j < m && i < n; j, i = j+1, i+1 {
				if arr1[i] > arr2[j] {
					arr1[i], arr2[j] = swap(arr1[i], arr2[j])
				}
			}

			if j < m {
				for ; j+g < m; j++ {
					if arr2[j] > arr2[j+g] {
						arr2[j], arr2[i+g] = swap(arr2[j], arr2[i+g])
					}
				}
			}
		}
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go merge(arr1, arr2, &wg)

	wg.Wait()
	fmt.Println(arr1)
	fmt.Println(arr2)
}
