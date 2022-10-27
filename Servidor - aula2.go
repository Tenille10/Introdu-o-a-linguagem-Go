// servidor service - servidor de serviço, que irá prestar algo ao usuário
// vamos falar em português e depois entender o "golanguês"

package main

import (
	"fmt"
	"net/http"
)

// Um conceito fundamental em servidores net/http  são os funções(que estão guarda
//dos no nosso pacote http) .

func ola(w http.ResponseWriter, req *http.Request) {
	// As funções que servem como manipuladores recebem a http.ResponseWriter e a
	// http.Request como argumentos. O gravador de resposta é usado para preencher a
	// resposta HTTP. Aqui nossa resposta simples é apenas “olá\n”.
	fmt.Fprintf(w, "Ola\n")
}

func cabecalhos(w http.ResponseWriter, req *http.Request) {
	// Esse manipulador faz algo um pouco mais sofisticado lendo todos os cabeçalhos de
	// solicitação HTTP e ecoando-os no corpo da resposta.
	for nome, cabecalhos := range req.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v: %v\n", nome, c)
		}
	}
}

func main() {
	// Um manipulador ("função") é um objeto que implementa http.Handler .
	// Uma maneira comum de escrever um manipulador(handler) é usar
	// o http.HandlerFunc adaptado as funções com a assinatura própria.
	// Registramos nossos manipuladores nas rotas do servidor usando a
	// http.HandleFunc, rota da função que ele deve chamar "/ola" e "/cabecalho"
	// Ele configura o roteador padrão no pacote net/http e recebe uma função
	//como argumento.("".ola) e ("", cabecalhos)
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalho", cabecalhos)
	// Finalmente, chamamos o ListenAndServe com a porta ":8090"
	//e um manipulador nil para que seja usado o roteador padrão que acabamos de configurar.
	// Execute o servidor em segundo plano
	// acesse :http://localhost:8090/ola
	//http://localhost:8090/cabecalho
	http.ListenAndServe(":8090", nil)
}
