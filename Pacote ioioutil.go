/* transformando o exemplo da aula em comentário:
package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Teste go!")
	err := ioutil.WriteFile("teste", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
Motivo:
De acordo com a própria documentação exemplo acima (da aula) foi descontinuado este é o link do exemplo acima na documentação: https://pkg.go.dev/io/ioutil#example-WriteFile
Contudo ainda acordo com a própria documentação a função agora simplesmente se chama os.WtriteFile e pode ser vista em: https://pkg.go.dev/os#WriteFile
E está aqui abaixo (com comentários):
*/

package main

import (
	"log" // pacote de log
	"os"  // pacote de sistema operacional (entrada e saída de dados do sistema operacional)
)

func main() {
	err := os.WriteFile("testdata/hello", []byte("Hello, Gophers!"), 0666)
	// escreve no arquivo "testdata/hello" a string "Hello, Gophers!"
	//e permite que o arquivo seja editado por qualquer usuário
	if err != nil {
		// se houver erro, imprime o erro e encerra o programa
		// se não houver erro, imprime "Arquivo salvo com sucesso!"
		log.Fatal(err) // imprime o erro e encerra o programa
	}
}
