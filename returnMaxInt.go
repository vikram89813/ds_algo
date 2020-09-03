package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var random = func(endpoint string, output chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(int64(time.Now().Unix()))
	randomValue := rand.Intn(1000)
	//fmt.Println(randomValue)
	output <- randomValue
}

func max(resources []string) (int, error) {
	maxOutput := 0
	var wg sync.WaitGroup
	output := make(chan int, 10)

	wg.Add(len(resources))

	for _, endpoint := range resources {
		go random(endpoint, output, &wg)
	}

	wg.Wait()

	close(output)

	fmt.Println("output")
	for element := range output {
		fmt.Printf("%d %d\n", maxOutput, element)
		if element > maxOutput {
			maxOutput = element
		}
	}

	return maxOutput, nil
}

// func TestMax(t *testing.T) {
// 	// randomMock := func(endPoint string) int {
// 	// 	return 5
// 	// }

// 	//random := randomMock

// 	randomArgs := []string{"abc.htm", "b.htm", "c.htm", "d.html", "e.htm"}

// 	value, _ := max(randomArgs)

// 	assert.Equal(t, value, 5)
// }

func main() {
	result, _ := max([]string{"abc.htm", "b.htm", "c.htm", "d.html", "e.htm"})
	fmt.Println(result)
}
