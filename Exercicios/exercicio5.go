package Exercicios

import "fmt"

func SafeInc(n *int) bool {
	fmt.Println("================Exercicio 5=================")
	if n == nil {
		fmt.Printf("============================================\n\n")
		return false
	}
	*n += 1
	fmt.Println(*n)
	fmt.Printf("============================================\n\n")
	return true
}
