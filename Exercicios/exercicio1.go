package Exercicios

import "fmt"

func Swap(a, b *int) {
	fmt.Println("================Exercicio 1=================")
	aux := *b
	*b = *a
	*a = aux
	fmt.Println(*a, *b)
	fmt.Printf("============================================\n\n")
}
