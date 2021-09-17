package main

import "fmt"

type Persona struct {
	nombre string
	apellido string
	edad   int
}
// Firma del metodp
func (p *Persona) agregarNombre(nom string, ape string ) {
	p.nombre=nom
	p.apellido=ape
}

func (p *Persona) EdadPersona(i int) {
	if i < 0 {
		fmt.Println("Error")
	} else {
		p.edad = i
	}
}

func main() {
	var persona Persona
	var l1, l2 string
	var l3 int

	fmt.Println("HOLA \n", "Digite su nombre y apellido por favor")
	fmt.Scanf("%s\n",  &l1)
	fmt.Scanf("%s\n",  &l2)
	fmt.Println( "Digite su edad por favor")
	fmt.Scanf("%d",  &l3)
	persona.agregarNombre(l1,l2)
	persona.EdadPersona(l3)
	fmt.Println(persona)
}
