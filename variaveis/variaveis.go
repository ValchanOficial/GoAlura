package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Valchan"
	var idade = 25
	var versao float64 = 1.1
	fmt.Println("Hello,", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão,", versao)
	fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
}
