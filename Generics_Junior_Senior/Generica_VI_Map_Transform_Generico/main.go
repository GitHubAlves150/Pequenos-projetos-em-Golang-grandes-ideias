//Autor: Lucas Lorenço Alves
//Programa: Map/Transform Generico
//Objetivo:Transformar slice de um tipo para outro

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//Template de função
func transformar[T, U any]( slice []T, transformadora func(T) U) []U{
	//crio um slice auxiliar'
	auxiliar:=make([]U, len(slice))

	
	//crio um laço para percorrer os indice, para poder aplciar a operação em cada indice

	for indice, valor:=range slice{
		auxiliar[indice]= transformadora(valor)
	}
	return auxiliar
}


func main(){

	
vetorint:=[]int{22, 43, 65, 77, 6545}
 vetorstring:=transformar(vetorint,  func(n int) string{
	s := strconv.Itoa(n)	
	return s

 })

 fmt.Println(":", vetorstring)
fmt.Println("Tipo convertido", reflect.TypeOf(vetorstring))


fmt.Println("=====================================================")

vetorString2:=transformar(vetorstring, func(n string) int{
	i, err:= strconv.Atoi(n)
	if err !=nil{
	 	fmt.Println("Erro na conversão: ", err)
		return 0
	}

	return i
} )

fmt.Println("Convertido: ", vetorString2)
fmt.Println(reflect.TypeOf(vetorString2))

	





}


