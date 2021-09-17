package main

import "fmt"

func main() {

	//ejemplo 1
	println("Digite 3 numeros, por favor")
	var b int
	var c int
	var d int

	fmt.Scanf("%d \n", &b)
	fmt.Scanf("%d \n", &c)
	fmt.Scanf("%d \n", &d)
	println("Menu de opciones: \n", "1. Sumar \n", "2. Multiplicar \n", "3. Restar \n", "4. Salir \n", "Elija:")
	var a int
	fmt.Scanf("%d \n", &a)
	switch a {
		case 1:
			{
				fmt.Println(b+c+d)
			}
		case 2:
			{
				fmt.Println(b*c*d)
			}
		case 3:
			{
				fmt.Println(b-c-d)
			}
		case 4:
			{
				break
			}
		default:
			fmt.Println("No existe esa operacion")
	}

	//Ejemplo 2
	/*
		  var a uint8
		  a = 3
			switch a {
			case 1:
				fallthrough
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 5:
				fmt.Println("Estas entre semana")
			case 6:
				fallthrough
			case 7:
				fmt.Println("Es fin de semana")
			default:
				fmt.Println("No es un dia valido")
			}
	*/

	//Ejemplo 3
	/*
	switch a := 7; {
	case a >= 0 && a <= 5:
		fmt.Println("Estas entres semana")
	case a >= 6 && a <= 7:
		fmt.Println("Estas en fin de semana")
	default:
		fmt.Println("No es un dia valido")

	}

	 */
}
