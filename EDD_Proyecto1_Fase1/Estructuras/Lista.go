package Lista

import (
	"fmt"
)

// Declaración de la estructura
type SimpleList struct {
	head *Estudiante
}

// Se agrega el puntero hacia el struct para hacerlo parte de el
func (list *SimpleList) Insert(nnombre string , Apellido string , Contraseña string) {
	//Declarar nuevo nodo
	newEstudiante := &Estudiante{nnombre: nnombre,
		 			Apellido: Apellido,
					Contraseña:Contraseña,
					next: nil}
	//Verificar si la lista está vacía
	if list.head == nil {
		list.head = newEstudiante
	} else {
		//Recorrer hasta encontrar el último nodo
		temp := list.head
		for temp.next != nil {
			temp = temp.next
		}
		//Agregar el nuevo nodo hasta el final
		temp.next = newEstudiante
	}
}

// Método para imprimir la lista
func (list *SimpleList) Print() {
	temp := list.head
	for temp.next != nil {
		fmt.Printf("  Nombre: "+temp.nnombre+
					"\n  Apellido: "+temp.Apellido+
					"\n  Contraseña"+temp.Contraseña)
		temp = temp.next
		fmt.Print("\n")
	}
	fmt.Printf(	"  Nombre: "+temp.nnombre+
				"\n  Apellido: "+temp.Apellido+
				"\n  Contraseña"+temp.Contraseña)
	fmt.Print("\n")
}

// Método para buscar un valor en la lista
func (list *SimpleList) Find(nnombre string) (result bool, index int) {
	temp := list.head
	index = 0
	result = false
	for temp != nil {
		if temp.nnombre == nnombre {
			result = true
		}
		index++
		temp = temp.next
	}

	if !result {
		index = -1
	}

	return result, index
}

// Método para eliminar valor de la lista
func (list *SimpleList) Delete(nnombre string) (result bool) {
	result = false
	// Si es el primero de la lista
	if list.head.nnombre == nnombre {
		result = true
		if list.head.next == nil {
			list.head = nil
		} else {
			list.head = list.head.next
		}
	} else {
		temp := list.head
		// Dos casos para eliminar [Se eliminará el 2]
		// 1 -> |2| -> 3 		----- Que el siguiente nodo no sea nulo
		// 1 -> |2| -> Null		----- Que el siguiente nodo sea nulo
		for temp != nil {
			//Encontrar el nodo anterior al que se elimina
			if temp.next != nil {
				if temp.next.nnombre == nnombre {
					break
				}
			}
			temp = temp.next
		}
		if temp == nil {
			fmt.Println("No se encontró el nombre D:")
			return false
		} else {
			//Se asigna el apuntador al siguiente del que se elimina
			// 1 -> |2| -> 3   		==>  1 -> 3
			// 1 -> |2| -> Null   	==>  1 -> Null (toma el valor del apuntando a nulo)
			temp.next = temp.next.next
			result = true
		}
	}
	return result
}