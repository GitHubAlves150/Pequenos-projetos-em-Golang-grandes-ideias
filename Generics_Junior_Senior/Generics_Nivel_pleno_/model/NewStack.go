package model
//struct generica
type Stack[T any] struct{
	elemento []T//membro de dados do tipo generico T
}

//stack vazia-Inicializa uma stack vazia
//NewStack[T ay]() é um template que retorna *Stack[T]
func NewStack[T any]() *Stack[T]{
	return&Stack[T]{
		elemento : make([]T, 0),
	}
}
