package model


import "fmt"


//Método isEmpty
func (empty *Stack[T]) isEmpty() bool{
	return len(empty.elemento)==0
}


//Método peek
func (P *Stack[T]) Peek()(T, error){
	if P.isEmpty(){
		var zero T
		return zero, fmt.Errorf("Pilha vazia")
		fmt.Println("Pilha vazia")
	}
	return P.elemento[len(P.elemento)-1], nil
}

