package main
import (
	"fmt"
	"time"
)
//HACE LA INTERSECCION EN GORUTINA
var k = make(chan int)

//FUNCION QUE IMPRIME MI NOMBRE
func nombre(num int) {
	fmt.Println("Yunisse", num)
	time.Sleep(100*time.Millisecond) //HACE UNA PAUSA DE 100 MIL MILISEGUNDOS
	k <- 0
}

func main(){
	println("Imprime mi nombre")
	for i:=1; i<10; i++{
		go nombre(i)
		for i:=1; i<1; i++{
			<- k
		}
	}
}

