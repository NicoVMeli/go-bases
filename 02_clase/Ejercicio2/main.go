/*
Ejercicio 2 - Descuento

Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos,
para ello necesitan una aplicación que les permita calcular el descuento en base a 2 variables,
su precio y el descuento en porcentaje. Espera obtener como resultado el valor con el descuento
aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos.
*/
package main

import "fmt"

func main() {
	var descuento int
	var precioProducto int
	descuento = 15
	precioProducto = 100
	aplicarDescuento(descuento, precioProducto)
}

func aplicarDescuento(descuento int, precio int) {
	//var precioTotal float32
	precioTotal := float32((precio - (precio * descuento / 100)))
	fmt.Printf("El precio a pagar es de %.2f, \nya que al precio original $%d\nle aplicamos un descuento de %d\n", precioTotal, precio, descuento)
}
