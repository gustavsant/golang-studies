package main

import "fmt"

type humano interface {
	cumprimento()
}

type pessoa struct {
	nome  string
	idade int
}

type pintor struct {
	pessoa
	tintaQueUsa string
}

type programador struct {
	pessoa
	linguaQuePrograma string
}

func (pi pintor) cumprimento() {
	fmt.Println("Olá! Me chamo " + pi.nome + " e uso " + pi.tintaQueUsa)
}

func (pr programador) cumprimento() {
	fmt.Println("Olá! Me chamo " + pr.nome + " e uso " + pr.linguaQuePrograma)
}

func cumprimentar(h humano) {
	h.cumprimento()
}

func main() {
	lindalva := pintor{
		pessoa: pessoa{
			"Lindalva",
			21,
		},
		tintaQueUsa: "Tinta acrílico",
	}

	gustavo := programador{
		pessoa: pessoa{
			"Gustavo",
			22,
		},
		linguaQuePrograma: "Go",
	}
	cumprimentar(lindalva)
	cumprimentar(gustavo)

}
