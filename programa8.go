package main

import (
	"fmt"
	"math/rand"
)

func numerosAleatorios() int {

	return rand.Intn(100)
}

func main() {
	var vec [10]int
	for i:=0; i<10; i++{
		vec[i]=numerosAleatorios()

		//fmt.Println(randomNum())
	}

	for j:=0; j<10; j++{
		fmt.Print(vec[j], "\t")
	}
}