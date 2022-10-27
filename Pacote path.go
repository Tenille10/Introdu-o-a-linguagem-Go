//implementa rotinas de utilitário para manipular caminhos de nome de arquivo

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
