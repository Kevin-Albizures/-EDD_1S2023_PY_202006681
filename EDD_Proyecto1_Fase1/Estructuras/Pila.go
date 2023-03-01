package Estructura

import (
	"fmt"
)


type Pila struct {
    head *Bitacora
	tail *Bitacora
}

func (Pila *Pila) EnPila(texto string , fechaa string) {
    newNodoB1 := &Bitacora{texto:texto, fechaa: fechaa, back:nil, next:nil}
    if Pila.head == nil {
        Pila.head = newNodoB1
		Pila.tail = newNodoB1
    } else {

        Pila.head.next = newNodoB1
        newNodoB1.back =Pila.head
		Pila.head = newNodoB1
    }
	fmt.Print("Inserción a nodo hecha.")
}

// Método para imprimir la lista
func (Pila *Pila) Print() {
	temp := Pila.head
	for temp.back != nil {
		fmt.Printf("\n"+temp.texto+temp.fechaa)
		temp = temp.back
		fmt.Print("\n")
	}
	fmt.Printf("\n"+temp.texto+temp.fechaa)
	fmt.Print("\n")
}

func (Pila *Pila) DePila(carnet int) interface{} {
    if Pila.head == nil {
        return nil
    }
    value := Pila.head
    Pila.head = Pila.head.back
    if Pila.head == nil {
        Pila.head = nil
    }
	fmt.Print("Bitacora reducida")
    return value
}


