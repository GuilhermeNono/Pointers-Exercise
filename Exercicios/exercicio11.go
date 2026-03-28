package Exercicios

import "fmt"

func GoRoutineChannelSoma() {
	ch := make(chan int)

	go sendNumbers(ch)

	var result int

	for item := range ch {
		result += item
	}

	fmt.Println("O resultado é", result)
}

// Apenas envia para o Consumer (ch chan <- int)
// Apenas recebe dos Producers (ch <- chan int)
func sendNumbers(channel chan<- int) {
	for i := 0; i <= 11; i++ {
		channel <- i
	}
	close(channel)
}
