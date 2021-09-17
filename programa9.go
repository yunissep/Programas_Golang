package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var vec [10]int
	var vec2 [10]int
	for i:=0; i<10; i++{
		vec[i]=rand.Intn(100)
	}
	for l:=0; l<10; l++{
		fmt.Print(vec[l], "\t")
	}
    println("\n")
	for k:=0; k<10; k++{
		vec2[k]=vec[k]*3
	}

	for j:=0; j<10; j++{
		fmt.Print(vec2[j], "\t")
	}
}
