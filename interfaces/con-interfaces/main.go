package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct {}

type pez struct{}

type pajaron struct{}

func (perro) mover() string{
	return "Soy un perro y camino"
}

func (pez) mover() string{
	return "Soy un pez y estoy nadando"
}

func (pajaron) mover() string{
	return "Soy un pájaro y estoy volando"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
} 

func main() {
	p := perro{}
	moverAnimal(p)
	pe := pez{}
	moverAnimal(pe)
	pa := pajaron{}
	moverAnimal(pa)
}