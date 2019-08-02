package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	// if comando == 1 {
	// 	fmt.Println("Monitorando..")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs..")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa..")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	switch lerComando() {
	case 1:
		fmt.Println("Monitorando..")
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
