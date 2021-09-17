package main
import (
	"fmt"
	"runtime"
)
//HACE LA INTERSECCION EN GORUTINA
var espera = make(chan int)

func numeros() { //FUNCION QUE IMPRIME LOS NUMEROS
	for i := 0; i < 50; i++ {
		fmt.Printf("%d \t", i)
	}
	espera <- 0
}

func main() {
	println("IMPRIMENDO LOS NUMEROS DEL 1 AL 50 USANDO 2 NUCLEOS")
	runtime.GOMAXPROCS(2) //SE SELECCIONAN 2 NUCLEOS
	go numeros()
	go numeros()
	for i := 0; i < 2; i++ {
		<- espera
	}
}