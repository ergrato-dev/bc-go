package main

import (
	"fmt"
)

func main() {
	// ============================================
	// PASO 1: El programa mínimo
	// ============================================
	// Todo programa Go necesita package main y func main().
	// fmt.Println imprime con salto de línea al final.
	// Descomenta las siguientes líneas:
	// fmt.Println("¡Hola, Go!")
	// fmt.Println()

	// ============================================
	// PASO 2: Declarar variables
	// ============================================
	// Go tiene dos formas de declarar variables:
	// - var: forma explícita, tipo opcional cuando se asigna valor
	// - :=  forma corta con inferencia (solo dentro de funciones)
	// Descomenta las siguientes líneas:
	// var lenguaje string = "Go"
	// version := "1.24"
	// fmt.Printf("Lenguaje: %s | Versión: %s\n", lenguaje, version)
	// fmt.Println()

	// ============================================
	// PASO 3: Múltiples valores de retorno
	// ============================================
	// Go permite retornar múltiples valores de una función.
	// Es el mecanismo base del manejo de errores idiomático.
	// Descomenta las siguientes líneas:
	// saludo, longitud := saludar("Gopher")
	// fmt.Printf("Saludo: %q (longitud: %d)\n", saludo, longitud)
	// fmt.Println()

	// ============================================
	// PASO 4: Manejo de errores
	// ============================================
	// En Go los errores son valores, no excepciones.
	// Una función que puede fallar retorna (resultado, error).
	// El error SIEMPRE debe verificarse — ignorarlo es un bug.
	// Descomenta las siguientes líneas:
	// resultado, err := dividir(10, 2)
	// if err != nil {
	//     fmt.Println("Error:", err)
	//     return
	// }
	// fmt.Printf("10 / 2 = %.2f\n", resultado)
	//
	// _, err = dividir(5, 0)
	// if err != nil {
	//     fmt.Println("División inválida:", err)
	// }
	// fmt.Println()

	// Evitar error "declared and not used" mientras comentas pasos
	_ = fmt.Sprintf("")
}

// saludar construye un saludo y retorna también su longitud.
// Muestra cómo Go permite múltiples valores de retorno.
func saludar(nombre string) (string, int) {
	saludo := "Hola, " + nombre + "!"
	// len() retorna el número de bytes, no de caracteres Unicode
	return saludo, len(saludo)
}

// dividir divide a entre b y retorna error si b es cero.
// Esta es la firma idiomática de Go: (resultado, error).
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// fmt.Errorf crea un error con mensaje formateado
		return 0, fmt.Errorf("no se puede dividir %v entre cero", a)
	}
	return a / b, nil
}
