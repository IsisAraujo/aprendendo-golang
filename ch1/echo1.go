package main // sempre precisará de um main

import ( //importo pacotes
	"fmt" // formata strings
	"os"  // manipula arquivos
)

func main() {
	var s, sep string                   //tipo de dado
	for i := 1; i < len(os.Args); i++ { //comeca do 1
		s += sep + os.Args[i] // salva de acordo com o i
		sep = " "             //recebe o valor
	}
	fmt.Println(s) // print na proxima linha
}
