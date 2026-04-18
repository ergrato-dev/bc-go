package main

import "fmt"

// ============================================
// PASO 3: función auxiliar para triplicar
// ============================================
// Descomenta esta función cuando llegues al Paso 3:
// func triplicar(n *int) {
// 	*n = *n * 3
// }

// ============================================
// PASO 4: función con nil check
// ============================================
// Descomenta esta función cuando llegues al Paso 4:
// func imprimirValor(p *int) {
// 	if p == nil {
// 		fmt.Println("puntero nil — sin valor")
// 		return
// 	}
// 	fmt.Println("valor:", *p)
// }

func main() {
	// ============================================
	// PASO 1: Obtener la dirección con &
	// ============================================
	fmt.Println("--- Paso 1: & dirección ---")

	// & retorna la dirección de memoria de n
	// El tipo de p es *int (puntero a int)
	// Descomenta las siguientes líneas:
	// n := 42
	// p := &n
	// fmt.Printf("valor de n: %d\n", n)
	// fmt.Printf("p apunta a la misma dirección que n: %v\n", p == &n)

	fmt.Println()

	// ============================================
	// PASO 2: Derreferir con *
	// ============================================
	fmt.Println("--- Paso 2: * derreferencia ---")

	// *p accede al valor en la dirección que guarda p
	// Modificar *p modifica n directamente
	// Descomenta las siguientes líneas:
	// v := 100
	// ptr := &v
	// fmt.Println("antes:", v)
	// *ptr = 200
	// fmt.Println("después:", v) // v cambió porque ptr apunta a v

	fmt.Println()

	// ============================================
	// PASO 3: Modificar el original con función
	// ============================================
	fmt.Println("--- Paso 3: modificar con función ---")

	// triplicar recibe *int y modifica el original
	// Descomenta las siguientes líneas (y la función triplicar arriba):
	// x := 5
	// fmt.Println("x antes:", x)
	// triplicar(&x) // pasamos la dirección de x
	// fmt.Println("x después:", x) // 15

	fmt.Println()

	// ============================================
	// PASO 4: Nil check antes de derreferir
	// ============================================
	fmt.Println("--- Paso 4: nil check ---")

	// var nulo *int declara un puntero nil (zero value de *int)
	// Descomenta las siguientes líneas (y la función imprimirValor arriba):
	// var nulo *int // nil
	// vv := 77
	// imprimirValor(nulo) // puntero nil — sin valor
	// imprimirValor(&vv)  // valor: 77
}
