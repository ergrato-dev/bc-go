// Paquete main: práctica 01 — Múltiples Retornos y Variádicas
//   Qué     → explorar los patrones de declaración de funciones en Go
//   Para qué → dominar múltiples retornos, retornos nombrados y variádicas
//   Impacto  → estos patrones son la base de toda la API estándar de Go

package main

import (
	"fmt"
	"strings"
)

// ============================================
// FUNCIONES DE APOYO
// Descomenta cada bloque de función junto con
// el paso correspondiente en main()
// ============================================

// PASO 1 — Descomenta este bloque:
// dividir retorna el cociente y un error si el divisor es cero.
// Patrón canónico en Go: retornar (resultado, error)
// func dividir(a, b float64) (float64, error) {
// 	if b == 0 {
// 		// fmt.Errorf crea un error con mensaje formateado
// 		return 0, fmt.Errorf("dividir: divisor no puede ser cero")
// 	}
// 	return a / b, nil // nil significa "sin error"
// }

// PASO 2 — Descomenta este bloque:
// estadisticas usa retornos nombrados: suma y promedio son variables locales.
// El naked return al final retorna sus valores actuales sin nombrarlos.
// func estadisticas(nums []float64) (suma, promedio float64) {
// 	for _, n := range nums {
// 		suma += n // suma es una variable local con zero value 0
// 	}
// 	if len(nums) > 0 {
// 		promedio = suma / float64(len(nums))
// 	}
// 	return // naked return: retorna suma y promedio tal como están
// }

// PASO 3 — Descomenta este bloque:
// unir concatena partes con un separador.
// El parámetro variádico ...string es []string dentro de la función.
// func unir(sep string, partes ...string) string {
// 	return strings.Join(partes, sep)
// }

// PASO 5 — Descomenta este bloque:
// transformFn es un tipo de función: toma un int y retorna un int.
// Definir el tipo permite usarlo como parámetro de otras funciones.
// type transformFn func(int) int

// aplicar recibe una función como argumento y la aplica a cada elemento.
// Este patrón se llama HOF (higher-order function).
// func aplicar(nums []int, fn transformFn) []int {
// 	resultado := make([]int, len(nums))
// 	for i, n := range nums {
// 		resultado[i] = fn(n)
// 	}
// 	return resultado
// }

// ============================================
// FUNCIONES NOMBRADAS para PASO 5
// ============================================

// PASO 5 — Descomenta este bloque:
// doble y triple son funciones del tipo transformFn.
// Se pueden pasar como argumentos a aplicar().
// func doble(n int) int  { return n * 2 }
// func cuadrado(n int) int { return n * n }

func main() {
	fmt.Println("=== Práctica 01: Múltiples Retornos y Variádicas ===")
	fmt.Println()

	// ============================================
	// PASO 1: Función con múltiples retornos
	// ============================================
	fmt.Println("--- Paso 1: Múltiples retornos ---")

	// En Go se capturan múltiples retornos con declaración corta :=
	// Descomenta las siguientes líneas:
	// resultado, err := dividir(10, 4)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("10 / 4 = %.2f\n", resultado)
	// }

	// También podemos provocar el error a propósito
	// Descomenta las siguientes líneas:
	// _, err2 := dividir(5, 0)
	// if err2 != nil {
	// 	fmt.Println("Error esperado:", err2)
	// }

	fmt.Println()

	// ============================================
	// PASO 2: Retornos nombrados
	// ============================================
	fmt.Println("--- Paso 2: Retornos nombrados ---")

	// Los retornos nombrados documentan el significado de cada valor.
	// La función usa naked return al final.
	// Descomenta las siguientes líneas:
	// notas := []float64{1, 2, 3, 4, 5}
	// s, prom := estadisticas(notas)
	// fmt.Printf("Suma: %.2f, Promedio: %.2f\n", s, prom)

	fmt.Println()

	// ============================================
	// PASO 3: Función variádica — valores individuales
	// ============================================
	fmt.Println("--- Paso 3: Función variádica ---")

	// Se puede llamar con cualquier cantidad de argumentos
	// Descomenta las siguientes líneas:
	// fmt.Println(unir("-", "a", "b", "c"))
	// fmt.Println(unir(" ", "Go", "es", "potente", "y", "rápido"))

	fmt.Println()

	// ============================================
	// PASO 4: Expansión de slice con ...
	// ============================================
	fmt.Println("--- Paso 4: Expansión de slice ---")

	// Cuando ya tienes un slice, usa ... para pasarlo a una variádica
	// El compilador expande el slice como si fueran argumentos individuales
	// Descomenta las siguientes líneas:
	// palabras := []string{"hola", "mundo", "Go"}
	// fmt.Println(unir(" ", palabras...)) // palabras... expande el slice

	fmt.Println()

	// ============================================
	// PASO 5: Función como valor (first-class)
	// ============================================
	fmt.Println("--- Paso 5: Funciones como valores ---")

	// Las funciones se pueden asignar a variables y pasar como argumentos
	// Descomenta las siguientes líneas:
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println("Originales:", nums)

	// Pasar función nombrada como argumento
	// fmt.Println("Dobles:    ", aplicar(nums, doble))

	// Pasar función nombrada
	// fmt.Println("Cuadrados: ", aplicar(nums, cuadrado))

	// Función anónima pasada directamente (no necesita nombre si se usa una vez)
	// fmt.Println("Triples:   ", aplicar(nums, func(n int) int { return n * 3 }))

	fmt.Println()

	// Evitar "imported and not used" si strings no se usa todavía
	_ = strings.Join
}
