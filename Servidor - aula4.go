//vamos criar nosso primeiro site estático?
//passo a passo
//um exemplo simples
//mas real
//código fará: servir arquivos HTML e CSS de um local específico no disco.

//passo1: comece criando um diretório (pasta para armazenar o projeto) - servidor
//passo2: criar seu arquivo go - aula4.go
//passo3: criar um arquivo html - example.html (cdm:type nul > example.html)
//passo4: criar um arquivo css - example.css (cdm:type nul > example.css)
//abrir os arquivos e escrever os código (deixarei no ppt da aula)

// pacote principal
package main

//importar os pacotes que usaremos: net/http e log
import (
	"log"
	"net/http"
)

// teremos uma função principal que:
func main() {
	fs := http.FileServer(http.Dir("./servidor"))
	http.Handle("/", fs)
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//1. irá inicializar nossa função servidor arquivo e abrir nosso diretório SERVIDOR
//função para criar um manipulador que responde a todas as solicitações HTTP
// com o conteúdo de um determinado sistema de arquivos .
//** fs := http.FileServer(http.Dir("./servidor"))
//, usamos a http.Handle()função para registrar o servidor de arquivos como o
//manipulador de todas as solicitações e iniciar o servidor ouvindo na porta 3000.
//** http.Handle("/", fs)
// "/"corresponde a todos os caminhos de solicitação, em vez de apenas ao caminho vazio.
