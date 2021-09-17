package main
import (
	"fmt"
	"time"
)
//HACE LA INTERSECCION EN GORUTINA
var es = make(chan int)

func alfabeto() { //FUNCION QUE IMPRIME EL ALFABETO
	for i:='a' ; i<= 'z'; i++ {
		time.Sleep(10*time.Millisecond) //hace una pausa de 10 mil milisegundos
		fmt.Printf("%c\t", i)
	}
	es <- 0
}

func main(){
	//LLAMA LA FUNCION ALFABETO
	print("IMPRIMIENDO EL ALFABETO \n")
	go alfabeto()
	for i := 0; i < 1; i++ {
		<- es
	}
}

