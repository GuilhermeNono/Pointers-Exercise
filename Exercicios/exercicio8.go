package Exercicios

type Point struct{ X, Y int }

func (p Point) MoveVal(dx, dy int) {
	p.X = dx
	p.Y = dy
}
func (p *Point) MovePtr(dx, dy int) {
	p.X = dx
	p.Y = dy
}
