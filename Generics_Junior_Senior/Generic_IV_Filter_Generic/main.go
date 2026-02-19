package main

import "fmt"

func Filter[T any](slice []T, condicao func(T) bool)[]T{
	//crio um slice auxiliar
	auxiliar:= make([]T, 0)

	//varro o slice
	for i:=0; i<len(slice); i++{
		if condicao(slice[i]) {
			auxiliar=append(auxiliar, slice[i])
		}
	}
	return auxiliar
}




func main() {

	vetor := []int{1, 2, 3, 24, 854, 321}

	react:=Filter(vetor, func (n int) bool {
		return n%2==0
	})

	fmt.Println(":", react)
	fmt.Println("======================")
	
}
