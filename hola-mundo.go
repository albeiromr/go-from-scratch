package main

import "fmt"

var numero int = 124
var texto string = "hola mundo"

func main() {

	nuevoNumero := 34
	nuevoTexto := "Albeiro"

	if numero == 125 {

		fmt.Printf("para todos un %v, soy %v", texto, nuevoTexto)
	} else {
		fmt.Printf("para todos un %v, con %d dedos", texto, nuevoNumero)
	}
	
}
