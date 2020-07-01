package main

import (
	"fmt"
	"os"

	"github.com/gabrielmq/site-monitor/src/logs"
	"github.com/gabrielmq/site-monitor/src/monitor"
)

func main() {
	exibirIntroducao()

	for {
		exibirMenu()
		comando := leComando()

		switch comando {
		case 1:
			monitor.IniciarMonitoramento()
		case 2:
			logs.ImprimeLogs()
		case 0:
			fmt.Println("Finalizando programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}

func exibirIntroducao() {
	nome := "Gabriel"
	versao := 1.1

	fmt.Println()
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println()
}

func exibirMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Finalizar programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println()
	return comando
}
