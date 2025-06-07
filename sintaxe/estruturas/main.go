package main

import "fmt"

type Produto struct {
	codigo     int
	descricao  string
	preco      float64
	quantidade int
}

func main() {
	var produto Produto = Produto{
		codigo:     10,
		descricao:  "teclado",
		preco:      275.5,
		quantidade: 50,
	}

	fmt.Printf("produto: %s\n", produto.descricao)
	fmt.Printf("produto: %d\n", produto.codigo)
	fmt.Printf("produto: %2f\n", produto.preco)
	fmt.Printf("produto: %.2f\n", float64(produto.quantidade)*produto.preco)
}
