package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 2
const delay = 5

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
	fmt.Println("")
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
	fmt.Print("Monitorando..\n")
	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.caelum.com.br", "https://www.google.com/"}

	for i := 0; i < monitoramentos; i++ {
		for j, site := range sites {
			fmt.Println("Testando site:", j, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	res, _ := http.Get(site)
	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas", res.StatusCode)
	}
}

//	SLICE
// func exibeNomes() {
// 	nomes := []string{"Douglas", "Valchan", "Bernardo", "Maria"}
// 	fmt.Println(nomes)
// 	fmt.Println(reflect.TypeOf(nomes))
// 	fmt.Println("O meu slice tem", len(nomes), "items")
// 	fmt.Println("O meu slice tem a capacidade para", cap(nomes), "items")
// 	nomes = append(nomes, "Aparecida")
// 	fmt.Println(nomes)
// 	fmt.Println("O meu slice tem", len(nomes), "items")
// 	fmt.Println("O meu slice tem a capacidade para", cap(nomes), "items")
// }
