package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// 1 2 3 4 5 8 13

	//posicao := uint(10)
	//fmt.Println(fibonacci(posicao))

	posicao := uint(5)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
