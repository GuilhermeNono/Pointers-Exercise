package Exercicios

import (
	"fmt"
	"sync"
	"time"
)

func WorkerPool(){
	const totalJobs = 10
	const totalWorkers = 3
	const totalResults = 10

	jobs := make(chan int, totalJobs)
	results := make(chan string, totalResults)

	var wg sync.WaitGroup

	wg.Add(totalWorkers)

	for range totalWorkers{
		go worker(&wg, jobs, results)
	}

	for j := range totalJobs {
		jobs <- j
	}
	close(jobs)

	go func(){
		wg.Wait()
		close(results)
	}()

	for result := range results{
		fmt.Print(result)
	}

	fmt.Printf("☑️   Todos os itens foram processados")

}

func worker(wg *sync.WaitGroup, jobs <- chan int, results chan <- string){

	defer wg.Done()

	for item := range jobs{
		fmt.Printf("⏲️   Processando item %d...\n", item)
		time.Sleep(500 * time.Millisecond)
		results <- fmt.Sprintf("✅   O item %d foi processado com sucesso.\n", item)
	}
}