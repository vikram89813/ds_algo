package main

import "fmt"

//var dp [][]bool

func helper(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 || len(s2) == 0 || len(s3) == 0 {
		return true
	}

	if len(s3) == 0 {
		return false
	}

	//if dp[]

	return (s1[0] == s3[0] && helper(s1[1:], s2, s3[1:])) || (s2[0] == s3[0] && helper(s1, s2[1:], s3[1:]))
}

func makeDp(a int, b int) [][]bool {
	arr := make([][]bool, a)
	for i := 0; i < a; i++ {
		arr[i] = make([]bool, b)
	}
	return arr
}

func test(s1 string, s2 string, s3 string) {
	//dp = makeDp(len(s1), len(s2))
	fmt.Println(helper(s1, s2, s3))
}

func main() {
	test("XXY", "XXZ", "XXZXXXY")
	test("XY", "WZ", "WZXY")
	test("XY", "X", "XXY")
	test("YX", "X", "XXY")
	test("XXY", "XXZ", "XXXXZY")
}
