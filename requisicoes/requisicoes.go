package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	for {
		exibeMenu()
		switch lerComando() {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs..")
		case 0:
			fmt.Println("Saindo do programa..")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) //Houve algum problema
		}
	}

}

func exibeIntroducao() {
	nome := "Valchan"
	var versao float64 = 1.1
	fmt.Println("Hello,", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando..")
	site := "https://random-status-code.herokuapp.com/"
	res, _ := http.Get(site)
	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas", res.StatusCode)
	}
}
