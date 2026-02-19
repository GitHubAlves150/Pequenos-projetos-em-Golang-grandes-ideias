//Autor Lucas Alves
//Exercicio 2: Minimos e Maximos
//Encontrar menor e maior valor em um slice

package main

import (
	"fmt"
	"golang.org/x/exp/constraints" //è uma biblioteca experimental do golang que define um conjunto de restiçoes(constraints) uteis para usar com Generics
)

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}

}
func Max[T constraints.Ordered](a, b T) T { //T é o retorno do tipo genérico
	if a > b {
		return a
	}
	return b
}

/*
"T pode ser qualquer tipo que suporte os operadores <, <=, >, >="

Tipos que implementam ordered
- Todos os inteiros (int, int8, int16, int32, int64, uint, etc)
- Todos os floats (float32, float64)
- Strings (string)

func Min[T constraints.Ordered](a, b T) T {
         └─────┬──────┘  └──┬──┘  └┬┘
               │            │      └── 3. Retorna T
               │            └── 2. Parâmetros a,b do tipo T
               └── 1. T deve ser Ordered (comparável)
}

Tipos que não implementam
- bool     // não faz sentido 10 < true
- structs  // como comparar {nome} < {idade}?
- slices   // não é possível comparar slices diretamente
- maps
*/

func main() {
	fmt.Println("Min:",  Min(10, 2))
	fmt.Println("Max:", Max(22, 33))

}
