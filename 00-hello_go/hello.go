// HELLO, WORD!
// package main // -> declara um pacote main -> um jeito de agrupar funções, é feito de todos os arquivos do mesmo diretório

// import "fmt" // -> pacote que contém funções para formatar texto, incluindo print e console. É uma das bibliotecas padrão

// func main() { // função que printa uma mensagem
// 	fmt.Println("Hello, Word!")
// }

// Quote module
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
