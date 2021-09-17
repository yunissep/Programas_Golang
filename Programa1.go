package main

import "fmt"
//HACE LA INTERSECCION EN GORUTINA
var l = make(chan int)

//FUNCION QUE CREAR LA TABLA
func mostrartabla(num int) {
	for i:=1 ; i<=12; i++{
		fmt.Println(num, "*",i, "=", i*num)
	}
	fmt.Println("\n")
	l <- 0
}

func main(){
	// AQUI SE LLAMA LA FUNCION PARA MOSTRAR LAS TABLAS
	println("Imprimir las tablas del 1 al 3")
	for i:=1; i<4; i++{
		go mostrartabla(i)
		for i:=0; i<1; i++{   <- l 		}
	}
}
