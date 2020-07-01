package logs

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// RegistraLog ...
func RegistraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	defer arquivo.Close()

	arquivo.WriteString(fmt.Sprintf("%s - %s - online: %t \n", time.Now().Format("02/01/2006 15:04:05"), site, status))
}

// ImprimeLogs ...
func ImprimeLogs() {
	fmt.Println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(string(arquivo))
}
