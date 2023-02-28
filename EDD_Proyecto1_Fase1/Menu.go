package main

import (
	"fmt"
	"os"
	//"os/exec"
	"bufio"
	"time"
)

func main() {

	var Op int

	// Variables de Verificación
	var usuario string
	var contraseña string

	scanner:=bufio.NewScanner(os.Stdin)

	// Menú de inicio
	
	fmt.Print("-----------------------------------------------\n")
	fmt.Print("--                                           --\n")
	fmt.Print("--             INICIO DE SESION              --\n")
	fmt.Print("--                                           --\n")
	fmt.Print("--  1. Inicio de sesión como administrador   --\n")
	fmt.Print("--  2. Inicio de sesión como estudiante      --\n")
	fmt.Print("--                                           --\n")
	fmt.Print("-----------------------------------------------\n")
	fmt.Print("                      - Ingrese la opción: ")
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&Op)

	switch Op {

	case 1:
		intentos:=0
		Op1:=0
		var OpM int
		fmt.Print("\n")
		for intentos<3 {
			fmt.Print("---------- Ingreso de adiministrador ----------\n")
			fmt.Print("--                                           --\n")
			fmt.Print("    - Ingrese usuario: ")
			scanner.Scan()
			fmt.Sscan(scanner.Text(),&usuario)
			fmt.Print("    - Ingrese contraseña: ")
			scanner.Scan()
			fmt.Sscan(scanner.Text(),&contraseña)
			fmt.Print("--                                           --\n")
			fmt.Print("-----------------------------------------------\n")

			if usuario=="Admin" && contraseña=="Admin" {
				intentos=3
				// Menú de admin
				fmt.Print("-----------------------------------------------\n")
				fmt.Print("--                                           --\n")
				fmt.Print("--             MENU ADMINISTRADOR            --\n")
				fmt.Print("--                                           --\n")
				fmt.Print("--  1. Ver Estudiantes Pendientes            --\n")
				fmt.Print("--  2. Ver Estudiantes del Sistema           --\n")
				fmt.Print("--  3. Registrar Nuevo Estudiante            --\n")
				fmt.Print("--  4. Carga Masiva de Estudiantes           --\n")
				fmt.Print("--  5. Generar Bítacora                      --\n")
				fmt.Print("--  6. Cerrar Sesión                         --\n")
				fmt.Print("--                                           --\n")
				fmt.Print(" ---------------------------------------------\n")
				fmt.Print("                      - Ingrese la opción: ")
				scanner.Scan()
				fmt.Sscan(scanner.Text(),&OpM)

				switch OpM {
					
				case 1:
					fmt.Print("Lista de estudiantes pendientes.")

				case 2:
					fmt.Print("Lista de estudiantes del sistema.")

				case 3:
					Nombre:=""
					Apellido:=""
					Carnet:=""
					Passw:=""
					fmt.Print("----------- Añadir nuevo estudiante -----------\n")
					fmt.Print("--                                           --\n")
					fmt.Print("    - Ingrese nombre: ")
					scanner.Scan()
					fmt.Sscan(scanner.Text(),&Nombre)
					fmt.Print("    - Ingrese apellido: ")
					scanner.Scan()
					fmt.Sscan(scanner.Text(),&Apellido)
					fmt.Print("    - Ingrese carnet: ")
					scanner.Scan()
					fmt.Sscan(scanner.Text(),&Carnet)
					fmt.Print("    - Ingrese Pass: ")
					scanner.Scan()
					fmt.Sscan(scanner.Text(),&Passw)
					fmt.Print("--                                           --\n")
					fmt.Print("-----------------------------------------------\n")

					fmt.Print("Ingresando estudiante..........")

				case 4:
					fmt.Print("Carga masiva.")

				case 5:
					fmt.Print("Generar bitacora.")

				case 6:
					fmt.Print("\n\n**************** Sesión cerrada... ****************\n")
					usuario=""
					contraseña=""

				default:
					fmt.Print("\n ****** Elija una opción válida. ******\n")
				}

			} else {
				fmt.Print("\n ***** Usuario o contraseña incorrectos. *****\n")
				intentos++
				fmt.Print("**                                           **\n")
				fmt.Print("** Regresar ingrese: 1                       **\n")
				fmt.Print("** Intentar de nuevo: ENTER                  **\n")
				fmt.Print("**                                           **\n")
				fmt.Print(" *********************************************\n")
				fmt.Print("                      - Ingrese la opción: ")
				fmt.Print("\n")
				scanner.Scan()
				fmt.Sscan(scanner.Text(),&Op1)

				if Op1==1 {
					intentos=3
				}
			}
		}
		
	case 2:
		intentos2:=0
		Op2:=0
		for intentos2<3 {
			fmt.Print("\n-------------- Ingreso de alumnos -------------\n")
			fmt.Print("--                                           --\n")
			fmt.Print("    - Ingrese usuario: ")
			scanner.Scan()
			fmt.Sscan(scanner.Text(),&usuario)
			fmt.Print("    - Ingrese contraseña: ")
			scanner.Scan()
			fmt.Sscan(scanner.Text(),&contraseña)
			fmt.Print("--                                           --\n")
			fmt.Print("-----------------------------------------------\n")

			if usuario=="" && contraseña=="" {
				intentos2=3
				now:=time.Now()
				fmt.Print("       - Accediste en la fecha: ")
				fmt.Println(now.Format("2006-01-02 15:04:05"))

			} else {
				fmt.Print("\n ***** Usuario o contraseña incorrectos. *****\n")
				intentos2++
				fmt.Print("**                                           **\n")
				fmt.Print("** Regresar ingrese: 1                       **\n")
				fmt.Print("** Intentar de nuevo: ENTER                  **\n")
				fmt.Print("**                                           **\n")
				fmt.Print(" *********************************************\n")
				fmt.Print("                      - Ingrese la opción: ")
				fmt.Print("\n")
				scanner.Scan()
				fmt.Sscan(scanner.Text(),&Op2)

				if Op2==1 {
					intentos2=3
				}

			}

		}

	default:
		fmt.Print("\n ****** Elija una opción válida. ******\n")


	}
	


}