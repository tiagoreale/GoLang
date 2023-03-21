package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Entrando na função para verificar se o aluno está aprovado1")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado2")

	media := (n1 + n2) / 2

	if media >= 6 {
		fmt.Println("Média calculada. Resultado será retornado3")
		return true
	}

	fmt.Println("Média calculada. Resultado será retornado4")
	return false
}

func main() {
	defer funcao1()
	// DEFER = ADIAR
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
