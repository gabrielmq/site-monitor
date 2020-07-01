package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// LeSitesDoArquivo ...
func LeSitesDoArquivo() []string {
	arquivo, err := os.Open("sites.txt")

	var sites []string
	if err != nil {
		fmt.Println("Erro:", err)
		return sites
	}

	defer arquivo.Close()

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		sites = append(sites, strings.TrimSpace(linha))

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Erro ao ler linha: %s", err)
		}
	}

	return sites
}
