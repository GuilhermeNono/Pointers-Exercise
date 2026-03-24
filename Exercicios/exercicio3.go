package Exercicios

import "fmt"

type Counter struct{ Value int }

func Reset(c *Counter) {
	fmt.Println("================Exercicio 3=================")

	c.Value = 0

	fmt.Println(*c)

	fmt.Printf("============================================\n\n")
}
