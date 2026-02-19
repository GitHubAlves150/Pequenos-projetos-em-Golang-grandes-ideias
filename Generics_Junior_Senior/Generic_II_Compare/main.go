//Autor Lucas Alves
//Exercicio 3: Constraint Genérico
//Veririficar se elementos existem em slice

package main

import (
	"fmt"
)

// Constrain verifica se um elemento está presente no slice
func Contain[T comparable](slice []T, elemento T) bool {
	for _, valor := range slice {
		if valor == elemento {
			return true
		}

	}
	return false
}

func main() {

	numeros := []int{1, 32, 4, 5}
	fmt.Println(Contain(numeros, 4))
	fmt.Println(Contain(numeros, 100))

	fmt.Println("===========================")
	nomes := []string{"Maria", "Jordana", "Julius"}

	fmt.Println(Contain(nomes, "Roberto"))
	fmt.Println(Contain(nomes, "Jordana"))

	fmt.Println("===========================")
	PontosFlutuantes := []float64{12.3, 33.89, 123.09}

	fmt.Println(Contain(PontosFlutuantes, 12.44))
	fmt.Println(Contain(PontosFlutuantes, 33.89))


}
