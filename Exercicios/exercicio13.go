package Exercicios

import (
	"fmt"
	"sync"
	"time"
)

func SyncWithWG() {
	var wg sync.WaitGroup

	wg.Add(3)

	go task("task 1", &wg)
	go task("task 2", &wg)
	go task("task 3", &wg)

	wg.Wait()
}

func task(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
	time.Sleep(3 * time.Second)
}
