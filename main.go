package main

import "fmt"

func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro encontrado:", r)
		}
	}()
	return a / b
}

func main() {

	fmt.Println(divide(4, 0))
	fmt.Println(divide(4, 2))
	fmt.Println("CONTINUEI EXECUTANDO")
}
