package main

import (
	"guilherme.nono/estudo/Exercicios"
)

func main() {
	//Exercicios.Swap(new(2), new(5))
	//Exercicios.Inc(new(5))
	//Exercicios.Reset(&Exercicios.Counter{
	//	Value: 5,
	//})
	//Exercicios.NewCounterA()
	//Exercicios.NewCounterB()
	//Exercicios.SafeInc(new(11))
	//Exercicios.UpdateMax(new(11), 15)
	//Exercicios.SliceAndPointers()
	//fmt.Println("================Exercicio 8=================")
	//
	//point := Exercicios.Point{X: 0, Y: 0}
	//
	//point.MovePtr(9, 5)
	//point.MoveVal(2, 9)
	//
	//fmt.Println(point)
	//fmt.Printf("============================================\n\n")
	//var p *int
	//Exercicios.AllocIfNil(&p)
	//Examples.GoRoutine()
	//Examples.AnnonimousFuncWGoRoutine()
	//Examples.RunningInTheSameTime()
	//Examples.WaitGroup()
	//Exercicios.GoRoutineWaitGroup()
	//Exercicios.GoRoutineChannelSoma()
	//Examples.MultiProdutorSoma()
	//Exercicios.SayConcurrentWaitGroup()
	//Exercicios.SyncWithWG()
	//Exercicios.CalculatePowerSquare()
	list := []int{4, 12, 5, 29}
	Exercicios.FanOutFanIn(list)
}
