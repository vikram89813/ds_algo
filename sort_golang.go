package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	sl := []string{"mumbai", "london", "tokyo", "seattle"}
	sort.Sort(sort.StringSlice(sl))
	fmt.Println(sl)

	intSlice := []int{3, 5, 6, 4, 2, 293, -34}
	sort.Sort(sort.IntSlice(intSlice))
	fmt.Println(intSlice)

	strSlice := []string{"apple", "durian", "kiwi", "banana"}
	sort.Sort(sort.Reverse(sort.StringSlice(strSlice)))
	fmt.Println(strSlice)

	str := "bcda"
	str2 := SortString(str)
	fmt.Println(str2)

}
