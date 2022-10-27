package main

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.print(w, "Home Page")
}

func HandleRequest() {
    http.HandleFunc("/", Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    fmt.PrintIn("Iniciando o servidor Rest com Go")
    HandleResquest()
}