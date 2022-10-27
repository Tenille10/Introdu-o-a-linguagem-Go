// servidor client - servidor relacionado ao que o usuário precisa
//nesse código vamos emitir uma solicitação de um cliente a um servidor http (aula 2)
//vamos codar em português , depois lemos em "golanguês"

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	//Emita uma solicitação HTTP GET para um servidor. http.Get é um atalho conveniente
	//para criar um client server objeto e chamar seu (Get) método;
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//ele usa o objeto que tem configurações padrão úteis, como predefinição do client server.
	//Imprima o status da resposta HTTP.
	fmt.Println("Qual é o status:", resp.Status) //O HTTP (Hypertext Transfer Protocol,
	//RFC 2616)
	//é o protocolo responsável por fazer a comunicação entre o cliente e o servidor.
	//Dessa forma, a cada “solicitação” feita, o HTTP responde se você obteve sucesso,
	// se não, se há algum erro na página, etc. Estas mensagens de erro são os
	//“status HTTP”. Por exemplo, os erros de HTTP mais comuns são: erro http 404,
	//erro http 500 e erro http 403.

	//Imprima as primeiras 5 linhas do corpo da resposta.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

//Classe de status 2XX – Sucesso
//Essa classe indica que a solicitação foi recebida, entendida e que será processada com êxito pelo servidor. Assim, são as respostas http para sucesso. Ou seja, tudo correu bem na transação.
//Por exemplo, os códigos mais comuns dessa classe são:

//200 OK;
//201 Criado;
//202 Aceito;
//203 não-autorizado;
//204 Nenhum conteúdo;
//205 Reset;
//206 Conteúdo parcial;
//207-Status Multi.
