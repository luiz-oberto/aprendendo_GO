package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Exported Name
	// nome de função que começa com letra maiúscula pode ser chamada por uma função que não está no memso pacote
	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// := -> esse operador é uma atalho para incializar e declarar uma variável em uma linha
	// só é possível utilizá-lo dentro de funções!!
	return message
}

// formas de declar variáveis:
// var mensagem_1 string = "isso é uma mensagem"
// var mensagem_2 = "isso é uma mensagem"
