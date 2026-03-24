package Exercicios

import "fmt"

func AllocIfNil(n **int) {
	if *n == nil {
		*n = new(7)
	}
	fmt.Println("================Exercicio 9=================")
	fmt.Println("n  (tipo **int):", n)
	fmt.Println("*n (tipo  *int):", *n)
	fmt.Println("**n(tipo   int):", **n)
	fmt.Printf("============================================\n\n")
}
