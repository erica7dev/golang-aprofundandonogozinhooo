package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	idade int
}

type humanos interface {
	falar()
}

func dizerAlgo(h humanos) {
	h.falar()
}

func (p *pessoa) falar(){
	fmt.Println(p.nome, "E aí!")
}

func main() {
	p1 := pessoa{"João",1000}

	//p1.falar()

	(&p1).falar()
	dizerAlgo(&p1)
}
