package Exercicios

import (
	"fmt"
	"sync"
	"time"
)

func SayConcurrent() {
	go say("Message say 1")
	go say("Message say 2")

	time.Sleep(5 * time.Second)
}

func SayConcurrentWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(2)

	go sayWWG("Message say 1", &wg)
	go sayWWG("Message say 2", &wg)

	wg.Wait()

	fmt.Println("Finalizou.")
}

func sayWWG(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}
}

func say(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}
}
