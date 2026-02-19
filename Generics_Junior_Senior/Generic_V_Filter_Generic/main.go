package main

import "fmt"

func Filter1[T any](slice []T, condicao func(T) bool) []T {
	//slice auxiliar
	vetor := make([]T, 0)

	//varredura
	for _, elemento := range slice {
		if condicao(elemento) {
			vetor = append(vetor, elemento)
		}
	}
	return vetor

}

func Filter2[T any](slice []T, condicao func(T) bool) []T {
	//slice auxiliar
	vetor := make([]T, 0)

	//varredura do slice para verificação
	for i := 1; i < len(slice); i++ {
		if condicao(slice[i]) {
			vetor = append(vetor, slice[i])

		}
	}
	return vetor
}

func main() {

	vetor := []int{1, 2, 54, 66, 543, 21, 56}

	pares := Filter1(vetor, func(n int) bool { return n%2 == 0 })

	fmt.Println(":", pares)

	fmt.Println("==================com for normal==============================")
	Pares := Filter2(vetor, func(k int) bool { return k%2 == 0 })

	fmt.Println(":", Pares)

}
