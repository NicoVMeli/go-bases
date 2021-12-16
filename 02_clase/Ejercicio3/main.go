/*
Ejercicio 3 - Préstamo

Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
- Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años,
- Se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
- Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje
de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var age uint8
	var isWorking string
	var cash int
	if evaluarEdad(age) {
		if estaTrabajando(isWorking) {
			evaluarSueldo(cash)
		}
	}
}

//Evalua la edad del usuario y retorna un bool para continuar con el proceso
func evaluarEdad(age uint8) bool {
	fmt.Println("Ingrese su edad")
	fmt.Scanln(&age) //tomo el valor ingresado en la terminal
	switch {         //como evaluo la edad en el case, aca no lleva condicion ni nada
	case age < 22:
		negacionDelPrestamos()
		return false
	default:
		fmt.Println("Por ahora vamos bien, unas preguntas más..")
		return true
	}
}

//Evalua si el string ingresado en una Y o N y retorna una respuesta
//a mejorar seria evaluar que solo se pueda ingresar esos valores
func estaTrabajando(isWorking string) bool {
	fmt.Println("Esta usted trabajando como esclavo????\nIngrese (Y o N)")
	fmt.Scanln(&isWorking)
	if strings.ToLower(isWorking) == "y" {
		fmt.Println("Vamos bien")
		return true
	} else {
		negacionDelPrestamos()
		return false
	}
}

func evaluarSueldo(cash int) {
	fmt.Println("De cuanto es tu sueldo actual titan??")
	fmt.Scanln(&cash)
	switch {
	case cash < 100000:
		fmt.Println("Te vamos a dar el prestamo, pero con interes.")
	default:
		fmt.Println("Crack, Groso aca esta tu prestamos y sin interes, vivilo")
	}
}

//Devulve un mensaje negando el prestamo
func negacionDelPrestamos() {
	fmt.Println("No podes obtener el beneficios del prestamo")
}
