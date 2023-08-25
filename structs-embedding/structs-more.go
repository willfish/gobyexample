package main

import "fmt"

type Position struct {
	x float64
	y float64
}

type SpecialPosition struct {
	Position
}

func (p *Position) Move(dx, dy float64) {
	p.x += dx
	p.y += dy
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

func (p *SpecialPosition) Move(dx, dy float64) {
	p.x += dx * 2
	p.y += dy * 2
}

type Player struct {
	*Position
}

type Enemy struct {
	*SpecialPosition
}

func NewPlayer(x, y float64) *Player {
	return &Player{&Position{x, y}}
}

func NewEnemy(x, y float64) *Enemy {
	return &Enemy{&SpecialPosition{Position{x, y}}}
}

func main() {
	player := NewPlayer(0, 0)
	player.Move(1, 1)
	player.Teleport(10, 10)

	enemy := NewEnemy(0, 0)
	enemy.Move(1.56121, 1)

	fmt.Println(player.x, player.y)

	fmt.Println(enemy.x, enemy.y)
}
