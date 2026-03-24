package Exercicios

import "fmt"

func Inc(n *int) {
	fmt.Println("================Exercicio 2=================")
	*n += 1

	fmt.Println(*n)
	fmt.Printf("============================================\n\n")
}
