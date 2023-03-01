package Estructura

import (
	"fmt"
	"strconv"
)

// Declaración de la estructura
type List struct {
	head *Estudiante
	tail *Estudiante
}

// Se agrega el puntero hacia el struct para hacerlo parte de el

func (list *List) Insert(nombre string , Apellido string,carnet int, Contraseña string) {
	nnombre:=nombre+" "+Apellido
	newEstudiante := &Estudiante{nnombre: nnombre,carnet: carnet,Contraseña:Contraseña,next: nil,back: nil}
	
	if list.head == nil {
		list.head = newEstudiante
		list.tail = newEstudiante
	} else {
		
		temp := list.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newEstudiante
		newEstudiante.back=temp
		
	}
	fmt.Print("Estudiante "+nnombre+"ingresado...\n")
}



// Método para imprimir la lista
func (list *List) Print() {
	temp := list.head
	if temp!=nil{
		fmt.Print("\n-------------- LISTADO DE ALUMNOS -------------\n")
		for temp.next != nil {
			fmt.Printf("- Nombre: "+temp.nnombre+", Carnet: "+strconv.Itoa(temp.carnet))
			temp = temp.next
			fmt.Print("\n-----------------------------------------------\n")
		}
			fmt.Printf("- Nombre: "+temp.nnombre+", Carnet: "+strconv.Itoa(temp.carnet))
			fmt.Print("\n-----------------------------------------------\n")
		fmt.Print("\n")
	}else{
		fmt.Print(" *********************************************\n")
		fmt.Print("*          No hay alumnos asignados          *\n")	
		fmt.Print(" *********************************************\n")
	}
	
}

// Método para buscar
func (list *List) Find(carnet int) (result bool) {
	temp := list.head
	result = false
	if list.head!=nil{
		for temp != nil {
			if temp.carnet == carnet {
				result = true
			}
			temp = temp.next
		}
	}
	return result
}

//Comprobar inicio de sesión
func (list *List) Comprobar(carnet int, contraseña string) (result bool) {
	temp := list.head
	result = false
	if list.head!=nil{
		for temp != nil {
			if temp.carnet == carnet && temp.Contraseña==contraseña {
				result = true
			}
			temp = temp.next
		}
	}
	return result
}



func (list *List) Verifica() (result bool){
    return list.head!=nil
}