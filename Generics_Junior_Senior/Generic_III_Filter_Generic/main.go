//Autor Lucas Alves
//Exercicio 4: Filter Generic
//Filtra elementos baseados em função

package main

import (
	"fmt"
)

// Filter que retorna novo slice com elementos que satisfazem a condição
func Filter[T any](slice []T, condicao func(T) bool) []T {
	resultado := make([]T, 0)
	for _, valor := range slice {

		if condicao(valor) {
			resultado = append(resultado, valor)
		}

	}
	return resultado

}

func main() {

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Filtra pares
	pares := Filter(number, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println("Pares", pares)

	//Filtrar maior que 5
	Maior5 := Filter(number, func(n int) bool {
		if n > 5 {
			return true
		}
		return false
	})

	fmt.Println("Maior que 5: ", Maior5)
}
