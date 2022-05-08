package main

import (
	"fmt"
)

func MudeMe(p *Pessoa) {
	(*p).sobrenome = "Oliveira"
}

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	pessoa1 := Pessoa{
		nome:      "Marcelo",
		sobrenome: "Azevedo",
		idade:     32,
	}

	MudeMe(&pessoa1)

	fmt.Println(pessoa1)
}
