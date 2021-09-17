package main

import "fmt"

func main() {

	paises := map[string]string{
		"RD": "Republica Dominicana",
		"ME": "Mexico",
		"FR": "Francia",
	}
	//imprimir mapa
	fmt.Println(paises)

	if pais, ok := paises["RD"]; ok {
		fmt.Println("La posicion si existe, es: ", pais)
	} else {
		fmt.Println("La posicion NO existe")
	}

}
