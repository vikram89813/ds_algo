package main

import (
    "fmt"
    "sort"
)

func main() {
    sl := []string{"mumbai", "london", "tokyo", "seattle"}
    sort.Sort(sort.StringSlice(sl))
    fmt.Println(sl)

    intSlice := []int{3,5,6,4,2,293,-34}
    sort.Sort(sort.IntSlice(intSlice))
    fmt.Println(intSlice)
}