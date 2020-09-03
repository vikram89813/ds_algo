package main

import (
	"sync"
)

func main() {
	// ---------- Fan-out----------------------------
	// primeStream := primeFinder(done, randIntStream)
	// numFinders := runtime.NumCPU()
	// finders := make([]<-chan int, numFinders)
	// for i := 0; i < numFinders; i++ {
	// 	finders[i] = primeFinder(done, randIntStream)
	// }

	// ---------- Fan-in----------------------------
	fanIn := func(
		done <-chan interface{},
		channels ...<-chan interface{},
	) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})
		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}
		// Select from all the channels
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}
		// Wait for all the reads to complete
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}

	// test code
	// done := make(chan interface{})
	// defer close(done)
	// start := time.Now()
	// rand := func() interface{} { return rand.Intn(50000000) }
	// randIntStream := toInt(done, repeatFn(done, rand))
	// numFinders := runtime.NumCPU()
	// fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	// finders := make([]<-chan interface{}, numFinders)
	// fmt.Println("Primes:")
	// for i := 0; i < numFinders; i++ {
	// 	finders[i] = primeFinder(done, randIntStream)
	// }
	// for prime := range take(done, fanIn(done, finders...), 10) {
	// 	fmt.Printf("\t%d\n", prime)
	// }
	// fmt.Printf("Search took: %v", time.Since(start))
}
