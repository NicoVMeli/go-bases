/*
Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego
tener cada una de las letras por separado para deletrearla.
Crear una aplicación que tenga una variable con la palabra e imprimir la
cantidad de letras que tiene la misma.
Luego imprimí cada una de las letras.
*/
package main

import "fmt"

func main() {
	var palabra []string
	palabra = append(palabra, "Mortadela")

	fmt.Println("La palabra ", palabra[0], " tiene ", len(palabra[0]), " letras.")
	contadorDeLetras(palabra[0])
}

//imprime las letras contenidas en el string pasado
//se parsea el string xq al acceder a la letra se transforma en byte
func contadorDeLetras(palabra string) {
	for i := 0; i < len(palabra); i++ {
		fmt.Println(string(palabra[i]))
	}
}
