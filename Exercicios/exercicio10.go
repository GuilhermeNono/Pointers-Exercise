package Exercicios

import (
	"fmt"
	"sync"
	"time"
)

func GoRoutineWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(3)

	go api(&wg)
	go banco(&wg)
	go cache(&wg)

	wg.Wait()

	fmt.Println("Sistema pronto")

}

func api(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Iniciando Api")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Api Finalizada")
}

func banco(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Iniciando Banco")
	time.Sleep(600 * time.Millisecond)
	fmt.Println("Banco Finalizado")
}

func cache(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Iniciando Cache")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Cache Finalizado")
}
