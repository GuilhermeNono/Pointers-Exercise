package Exercicios

import "fmt"

func UpdateMax(max *int, v int) bool {
	fmt.Println("================Exercicio 6=================")

	if v > *max {
		*max = v
		fmt.Println(*max)
		fmt.Printf("============================================\n\n")
		return true
	}
	fmt.Printf("============================================\n\n")
	return false
}
