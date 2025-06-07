package main

import "fmt"

// Definindo um tipo 'Animal'
type Animal struct {
	Nome string
}

// Definindo um método 'FazerBarulho' para 'Animal'
func (a Animal) FazerBarulho() {
	fmt.Println("Barulho genérico de animal.")
}

// Definindo um tipo 'Cachorro' que "herda" de 'Animal'
type Cachorro struct {
	Animal // Composição
}

// Sobrescrevendo o método 'FazerBarulho' para 'Cachorro'
func (c Cachorro) FazerBarulho() {
	fmt.Println("Au au!")
}

func main() {
	cachorro := Cachorro{Animal: Animal{Nome: "Rex"}}
	// Chamando o método 'FazerBarulho' do 'Cachorro'
	cachorro.FazerBarulho()
	// Chamando o método 'FazerBarulho' do 'Animal' (herança "light")
	cachorro.Animal.FazerBarulho()
}
