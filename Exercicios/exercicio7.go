package Exercicios

import "fmt"

func SliceAndPointers() {
	fmt.Println("================Exercicio 7=================")
	slice := []int{10, 20, 30}
	fmt.Println(slice)
	slicePtrs := []*int{&slice[0], &slice[1], &slice[2]}
	fmt.Println(*slicePtrs[0], *slicePtrs[1], *slicePtrs[2])
	*slicePtrs[1] = 5

	fmt.Println(*slicePtrs[0], *slicePtrs[1], *slicePtrs[2])
	fmt.Printf("============================================\n\n")
}
