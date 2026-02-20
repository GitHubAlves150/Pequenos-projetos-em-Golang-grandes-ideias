//Autor: Lucas Alves
//Nível Pleno
//Estrutura e constraints
//Implementar a estrutura de dados Lifo usando generic em Go, ou Template como preferir

package main
 import (
	"fmt"
	"GENERIC_NIVEL_PLENO/model"
 )



//Método pop


//Método size


 func main(){

     pilha :=model.NewStack[int]()
	pilha.Push(22)
	pilha.Push(32)
	pilha.Push(55)
	 
	_,err:= pilha.Peek()
	if err !=nil{
		fmt.Println("erro de pilah vazia")
	}
	fmt.Println(":", pilha)

	fmt.Println("-+===============Inserindo=====================")

	pilhaStr := model.NewStack[string]()
	pilhaStr.Push("Lucas")
	pilhaStr.Push("Jordan")

	_,erro:= pilhaStr.Peek()
	if erro!=nil{
		fmt.Println("erro de pilha vazia")
	}


	fmt.Println(":", pilhaStr)

	fmt.Println("-+==============Removendo======================")
	
	pilha.Pop()
	fmt.Println(":", pilha)

	
 	fmt.Println("...Fim....")

 }
