package monitor

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gabrielmq/site-monitor/src/logs"
	file "github.com/gabrielmq/site-monitor/src/utils"
)

const monitoramentos = 2
const delay = 2

// IniciarMonitoramento ...
func IniciarMonitoramento() {
	fmt.Println("Iniciando Monitoramento...")

	sites := file.LeSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			go testaSite(site)
		}

		fmt.Println()
		time.Sleep(delay * time.Second)
	}

	fmt.Println()
}

func testaSite(site string) {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "carregado com sucesso!")
		logs.RegistraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas; Status:", res.Status)
		logs.RegistraLog(site, false)
	}
}
