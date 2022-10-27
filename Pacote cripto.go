// a cripto são os exemplos criptografados
// vamos usar uma muito comum a SHA-1
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()
	h.Write([]byte("código com pacote hash"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
