package main

import "fmt"

func main() {
	var i interface{}

	i = 10
	//str := "abn fsdf"

	fmt.Println(i.(int))
	//fmt.Println(str.(string))
}
