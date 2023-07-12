package main

import "fmt"
import "os"

func main(){
	fmt.Println("Bem-vindo ao programa X")
	fmt.Print("Qual é o seu nome? > ")
	var nome string
	fmt.Scan(&nome)
	println("Olá,",nome, ", vamos escolher algo:")
	fmt.Println("1 - Monitoramento")
	fmt.Println("2 - Logs")
	fmt.Println("0 - Sair")
	var escolha int
	fmt.Scan(&escolha)
	switch escolha{
		case 1:
			fmt.Println("Iniciando Monitoramento...")
		case 2:
			fmt.Println("Iniciando a verificação dos Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
	}

}