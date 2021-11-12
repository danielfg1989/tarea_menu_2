package main

import "fmt"

func main() {
	menu_principal := `	Menu Principal
	1-Cambiar contrasena
	2-Editar Perfil
	3-Foro
	4-Finalizar programa
	`
	menu_perfil := `	1-Cambiar foto de perfil
	2-Cambiar nombre de usuario
	3-Editar biografia
	4-Volver al menu principal`

	menu_foro := `	1-Publicar nueva entrada en el foro
	2-Editar nueva entrada en el foro
	3-Eliminar entrada del foro
	4-Volver al menu principal`

	var opc_menu_principal int
	var opc_menu_perfil int
	var opc_menu_foro int

	menu := 0

	for {
		if menu == 0 {
			fmt.Println(menu_principal)
			fmt.Scanf("%d", &opc_menu_principal)
		}
		switch opc_menu_principal {
		case 1:
			fmt.Println("Su contrasena ha sido cambiado con exito")
			break
		case 2:
			fmt.Println(menu_perfil)
			fmt.Scanf("%d", &opc_menu_perfil)
			switch opc_menu_perfil {
			case 1:
				fmt.Println("Foto de perfil cambiada")
				menu = 2
			case 2:
				fmt.Println("Nombre de usuarioo cambiado")
				menu = 2
			case 3:
				fmt.Println("Biografia editada")
				menu = 2
			case 4:
				menu = 0
				break
			default:
				fmt.Println("Ingreso una opcion invalida")

			}
		case 3:
			fmt.Println(menu_foro)
			fmt.Scanf("%d", &opc_menu_foro)
			switch opc_menu_foro {
			case 1:
				fmt.Println("Entrada del foro publicada")
				menu = 2
			case 2:
				fmt.Println("Entrada del foro editada")
				menu = 2
			case 3:
				fmt.Println("Entrada del foro eliminada")
				menu = 2
			case 4:
				menu = 0
				break
			default:
				fmt.Println("Ingreso una opcion invalida")

			}
		case 4:
			menu = 4
			break
		default:
			fmt.Println("Ingreso un valor incorrecto")
		}

		if menu == 4 {
			fmt.Println("Hasta luego, esperamos volver a verte pronto")
			break
		}
	}
}
