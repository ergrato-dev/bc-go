// Paquete main: punto de entrada de la práctica 01 — Tipos Básicos en Go
//   Qué     → programa guiado para explorar el sistema de tipos de Go
//   Para qué → aprender a declarar variables, observar zero values y convertir tipos
//   Impacto  → descomenta cada sección siguiendo el README; go run . después de cada paso

package main

import "fmt"

func main() {
	// ============================================
	// PASO 1: Declarar variables de tipos básicos
	// ============================================
	fmt.Println("--- Paso 1: Tipos básicos ---")

	// Go tiene un sistema de tipos estático: cada variable tiene un tipo fijo.
	// var declara la variable sin inicializar — recibe su zero value automáticamente.

	// Descomenta las siguientes líneas:
	// var quantity  int
	// var price     float64
	// var name      string
	// var available bool
	//
	// fmt.Printf("int:     %v (%T)\n", quantity,  quantity)
	// fmt.Printf("float64: %v (%T)\n", price,     price)
	// fmt.Printf("string:  %q (%T)\n", name,      name)
	// fmt.Printf("bool:    %v (%T)\n", available, available)

	fmt.Println()

	// ============================================
	// PASO 2: Zero values — la promesa de Go
	// ============================================
	fmt.Println("--- Paso 2: Zero values ---")

	// Toda variable declarada sin inicializar recibe su zero value.
	// Esto garantiza que no haya "garbage values" como en C.
	//
	// Tipos de valor      → zero value concreto (0, false, "")
	// Tipos de referencia → zero value es nil (punteros, slices, maps, canales)

	// Descomenta las siguientes líneas:
	// var i int
	// var f float64
	// var b bool
	// var s string
	// var p *int
	//
	// fmt.Printf("int zero:     %d\n",  i)
	// fmt.Printf("float64 zero: %f\n",  f)
	// fmt.Printf("bool zero:    %t\n",  b)
	// fmt.Printf("string zero:  %q\n",  s)
	// fmt.Printf("puntero zero: %v\n",  p)

	fmt.Println()

	// ============================================
	// PASO 3: Conversión explícita de tipos
	// ============================================
	fmt.Println("--- Paso 3: Conversión de tipos ---")

	// Go NO permite conversiones implícitas.
	// Debes usar la sintaxis: TipoDestino(valor)
	// La conversión float64 → int trunca (no redondea).

	// Descomenta las siguientes líneas:
	// items := 7
	// pricePerItem := 4.99
	//
	// // float64(items) convierte int a float64 para la multiplicación
	// total := float64(items) * pricePerItem
	// fmt.Printf("Items: %d, Precio: %.2f, Total: %.2f\n", items, pricePerItem, total)
	//
	// // int(total) trunca: 34.93 → 34 (no redondea a 35)
	// approximate := int(total)
	// fmt.Printf("Aproximado (truncado): %d\n", approximate)

	fmt.Println()

	// ============================================
	// PASO 4: Verbos de formato fmt.Printf
	// ============================================
	fmt.Println("--- Paso 4: Verbos de formato ---")

	// fmt.Printf usa verbos para controlar la representación del valor.
	// %T es especialmente útil para depurar el tipo real de una variable.

	// Descomenta las siguientes líneas:
	// x  := 42
	// pi := 3.14159
	// ch := 'G'
	// msg := "Hola, Go"
	//
	// fmt.Printf("%%v  (default)  → %v\n",   x)
	// fmt.Printf("%%T  (tipo)     → %T\n",   x)
	// fmt.Printf("%%d  (decimal)  → %d\n",   x)
	// fmt.Printf("%%b  (binario)  → %b\n",   x)
	// fmt.Printf("%%x  (hex)      → %x\n",   x)
	// fmt.Printf("%%.2f (flotante) → %.2f\n", pi)
	// fmt.Printf("%%c  (carácter) → %c\n",   ch)
	// fmt.Printf("%%q  (quoted)   → %q\n",   msg)

	fmt.Println()

	// ============================================
	// PASO 5: string, byte y rune — texto Unicode
	// ============================================
	fmt.Println("--- Paso 5: string, byte y rune ---")

	// Un string en Go es una secuencia inmutable de bytes (no de caracteres).
	// Para texto Unicode, len() cuenta bytes; len([]rune(s)) cuenta caracteres.
	// range sobre un string itera por runes, no por bytes.

	// Descomenta las siguientes líneas:
	// greeting := "Hola, 世界"
	//
	// fmt.Printf("String:         %q\n", greeting)
	// fmt.Printf("Bytes  (len):   %d\n", len(greeting))
	// fmt.Printf("Runes  ([]rune): %d\n", len([]rune(greeting)))
	//
	// fmt.Println("\nIteración por runes:")
	// for i, r := range greeting {
	//     fmt.Printf("  [%2d] %c  U+%04X\n", i, r, r)
	// }
}
