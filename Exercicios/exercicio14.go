package Exercicios

import "fmt"

func CalculatePowerSquare() {
	ch := make(chan int)

	go powerSquare(ch)

	for item := range ch {
		fmt.Println(item)
	}
}

func powerSquare(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i * i
	}
	close(ch)
}
