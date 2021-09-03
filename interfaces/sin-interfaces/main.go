package main

import "fmt"

type perro struct {}

type pez struct{}

type pajaron struct{}

func (perro) caminar() string{
	return "Soy un perro y camino"
}

func (pez) nadar() string{
	return "Soy un pez y estoy nadando"
}

func (pajaron) volar() string{
	return "Soy un pájaro y estoy volando"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(p pez) {
	fmt.Println(p.nadar())
}

func moverPajaro(p pajaron) {
	fmt.Println(p.volar())
}

func main() {
	p := perro{}
	moverPerro(p)
	pe := pez{}
	moverPez(pe)
	pa := pajaron{}
	moverPajaro(pa)
}