package Examples

import (
	"fmt"
	"sync"
	"time"
)

func GoRoutine() {
	go ola()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Fim do main")
}

func AnnonimousFuncWGoRoutine() {
	go func() {
		fmt.Println("Executando uma funcao anonima na go routine.")
	}()

	time.Sleep(100 * time.Millisecond)
}

func RunningInTheSameTime() {
	go goImprimir("goroutine A")
	go goImprimir("goroutine B")
	time.Sleep(1 * time.Second)
	fmt.Println("Fim do main")
}

func goImprimir(texto string) {
	for i := 0; i < 3; i++ {
		fmt.Println(texto, i)
		time.Sleep(200 * time.Millisecond)
	}
}

func ola() {
	fmt.Println("Ola da Goroutine.")
}

func WaitGroup() {
	var wg sync.WaitGroup

	wg.Add(2)

	go tarefa("Tarefa 1", &wg)
	go tarefa("Tarefa 2", &wg)

	wg.Wait()
	fmt.Println("Todas as tarefas terminaram")
}

func tarefa(nome string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Iniciou a Tarefa da Goroutine.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Finalizou a Tarefa da Goroutine.")
}

func MultiProdutorSoma() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(2)

	go produtor("P1", 1, 5, ch, &wg)
	go produtor("P2", 6, 10, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	var soma int

	for item := range ch {
		fmt.Println("Consumidor recebeu:", item)
		soma += item
	}

	fmt.Println("Soma final:", soma)
}

func produtor(nome string, inicio, fim int, channel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := inicio; i <= fim; i++ {
		fmt.Printf("%s enviou %d\n", nome, i)
		channel <- i
	}
}

func WorkerPoolBasico() {
	const totalJobs = 10
	const totalWorkers = 3

	jobs := make(chan int, totalJobs)
	results := make(chan string, totalJobs)

	var wg sync.WaitGroup

	//1) Subindo os Workers
	for w := 1; w <= totalWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	//2) Envia Jobs
	for j := 1; j <= totalJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//3) Espera os outros terminarem e fecha o result
	go func() {
		wg.Wait()
		close(results)
	}()

	//4)Consome Resultados
	for r := range results {
		fmt.Println(r)
	}

	fmt.Println("Pool Finalizada")
}

func worker(id int, jobs chan int, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		//Simulação de um processamento
		time.Sleep(400 * time.Millisecond)
		results <- fmt.Sprintf("worker %d processou job %d", id, job)
	}

}
