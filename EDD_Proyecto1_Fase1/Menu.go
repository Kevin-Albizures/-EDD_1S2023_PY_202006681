package main

import (
	"fmt"
	"github.com/inancgumus/screen"
)

func main() {

	var Op int

	// Variables de Verificación
	var usuario string
	var contraseña string

	// Menú de inicio
	fmt.Print("----------------------------------------------\n")
	fmt.Print("-                                           --\n")
	fmt.Print("-             INICIO DE SESION              --\n")
	fmt.Print("-                                           --\n")
	fmt.Print("-  1. Inicio de sesión como administrador   --\n")
	fmt.Print("-  2. Inicio de sesión como estudiante      --\n")
	fmt.Print("-                                           --\n")
	fmt.Print("----------------------------------------------\n")
	fmt.Scanln(Op)

	switch Op {

	case 1:
		intentos:=0
		Op1:=0
		for intentos==3 {
			screen.Clear()
			fmt.Print("Ingrese usuario: ")
			fmt.Scanln(usuario)
			fmt.Print("Ingrese contraseña: ")
			fmt.Scanln(contraseña)
			
			if usuario=="Admin" && contraseña=="Admin" {
				intentos=3
				// Menú de admin
				fmt.Print("----------------------------------------------\n")
				fmt.Print("-                                           --\n")
				fmt.Print("-             MENU ADMINISTRADOR            --\n")
				fmt.Print("-                                           --\n")
				fmt.Print("-  1. Ver Estudiantes Pendientes            --\n")
				fmt.Print("-  2. Ver Estudiantes del Sistema           --\n")
				fmt.Print("-  3. Registrar Nuevo Estudiante            --\n")
				fmt.Print("-  4. Carga Masiva de Estudiantes           --\n")
				fmt.Print("-  5. Generar Bítacora                      --\n")
				fmt.Print("-  6. Cerrar Sesión                         --\n")
				fmt.Print("-                                           --\n")
				fmt.Print("----------------------------------------------\n")
			} else {
				fmt.Print("Usuario o contraseña incorrectos. \n")
				intentos++
				fmt.Print("Regresar ingrese: 1\n")
				fmt.Print("Intentar de nuevo: ENTER \n")
				fmt.Scanln(Op1)

				if Op1==1 {
					intentos=3
				}
			}
		}
		
	case 2:
		intentos:=0
		Op1:=0
		for intentos==3 {
			screen.Clear()
			fmt.Print("Ingrese usuario: ")
			fmt.Scanln(usuario)
			fmt.Print("Ingrese contraseña: ")
			fmt.Scanln(contraseña)

			if usuario=="Admin" && contraseña=="Admin" {
				intentos=3
				// Menú de admin
				
			} else {
				fmt.Print("Usuario o contraseña incorrectos. \n")
				intentos++
				fmt.Print("Regresar ingrese: 1 \n")
				fmt.Print("Intentar de nuevo: ENTER \n")
				fmt.Scanln(Op1)

				if Op1==1 {
					intentos=3
				}

			}

		}


	}
	


}