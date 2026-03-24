package Exercicios

import "fmt"

func NewCounterA() *Counter { // usando new
	fmt.Println("================Exercicio 4=================")
	c := new(Counter{
		Value: 0,
	})

	fmt.Println(*c)

	return c
}
func NewCounterB() *Counter { // usando &Counter{...}
	c := &Counter{Value: 0}

	fmt.Println(*c)
	fmt.Printf("============================================\n\n")
	return c
}
