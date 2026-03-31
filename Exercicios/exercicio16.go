package Exercicios

import (
	"fmt"
	"sync"
)

func FanOutFanIn(n []int){

	const workers = 4

	result := 0

	numbersCh := make(chan int, 5)

	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(workers)

	for i := 0; i < workers; i++{
		go processor(&wg, &mu, numbersCh, &result)
	}

	for i := 0; i < len(n); i++{
		numbersCh <- n[i]
	}
	close(numbersCh)

	wg.Wait()

	fmt.Print(result)

}

func processor(wg *sync.WaitGroup, mu *sync.Mutex, ch chan int, result *int){
	defer wg.Done()

	for item := range ch{
		mu.Lock()
		*result += (item * 2)
		mu.Unlock() 
	}
}