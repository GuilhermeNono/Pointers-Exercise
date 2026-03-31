package Exercicios

import (
	"fmt"
	"sync"
)

func FanOutFanIn(n []int) {

	const workers = 4

	resultCh := make(chan int)
	numbersCh := make(chan int)

	var wg sync.WaitGroup

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go processor(&wg, numbersCh, resultCh)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	go func() {
		for _, item := range n {
			numbersCh <- item
		}
		close(numbersCh)
	}()

	result := 0

	for i := range resultCh {
		result += i
	}

	fmt.Print(result)

}

func processor(wg *sync.WaitGroup, ch chan int, result chan int) {
	defer wg.Done()

	for item := range ch {
		result <- item * 2
	}
}
