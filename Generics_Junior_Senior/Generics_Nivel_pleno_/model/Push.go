package model


//Método push
func (i *Stack[T]) Push(valor T){
	i.elemento= append(i.elemento, valor)
	
}