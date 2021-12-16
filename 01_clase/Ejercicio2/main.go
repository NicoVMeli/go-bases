/*
Ejercicio 2 - Clima

Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares.
Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
Imprime los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?
*/

package main

import "fmt"

func main() {
	var temperatura float32
	var humedad float32
	var presionAtmosferica float32
	temperatura = 17.5432
	humedad = 43.2222
	presionAtmosferica = 32.23231
	fmt.Printf("La temperatura es de %.2f \n", temperatura)
	fmt.Printf("La humedad es de %.2f \n", humedad)
	fmt.Printf("Y la presion de %.2f\n", presionAtmosferica)

}
