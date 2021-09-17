package main
import (
	"fmt"
	"time"
)
//HACE LA INTERSECCION EN GORUTINA
var e = make(chan int)
func color(num int) { //FUNCION QUE IMPRIME EL COLOR NEGRO
	for i := 0; i < num; i++ {
		fmt.Println("NEGRO")
		time.Sleep(10*time.Millisecond) //HACE UNA PAUSA DE 10 MIL MILISEGUNDOS
	}
	e <- 0
}
func main() {
	//LLAMA LA FUNCION COLOR
	println("IMPRIMIENDO COLOR NEGRO")
	for i := 0; i < 20; i++ {
		go color(i)
		for i := 0; i < 1; i++ {
			<- e
		}
	}
}
