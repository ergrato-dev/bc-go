// Paquete main: práctica 02 — Closures y Funciones de Orden Superior
//   Qué     → construir closures y HOF en Go
//   Para qué → encapsular estado, componer comportamiento y evitar repetición
//   Impacto  → los closures son la base de middleware, callbacks y generadores

package main

import (
	"fmt"
)

// ============================================
// FUNCIONES DE APOYO
// Descomenta cada bloque junto con el paso
// correspondiente en main()
// ============================================

// PASO 1 — Descomenta este bloque:
// contadorNuevo retorna un closure que incrementa n en cada llamada.
// Cada llamada a contadorNuevo crea un entorno heap independiente.
// func contadorNuevo() func() int {
// 	n := 0 // n escapa al heap porque el closure la captura
// 	return func() int {
// 		n++ // modifica n del entorno padre
// 		return n
// 	}
// }

// PASO 2 — Descomenta este bloque:
// multiplicadorPor es una fábrica de funciones.
// 'factor' queda capturado en cada closure que retorna.
// func multiplicadorPor(factor int) func(int) int {
// 	return func(n int) int {
// 		return n * factor
// 	}
// }

// prefijador retorna una función que antepone prefijo a strings.
// func prefijador(prefijo string) func(string) string {
// 	return func(s string) string {
// 		return prefijo + s
// 	}
// }

// PASO 3 — Descomenta este bloque:
// filtrar retorna los elementos de nums que satisfacen pred.
// pred es una función: recibe int, retorna bool.
// func filtrar(nums []int, pred func(int) bool) []int {
// 	resultado := make([]int, 0, len(nums))
// 	for _, n := range nums {
// 		if pred(n) {
// 			resultado = append(resultado, n)
// 		}
// 	}
// 	return resultado
// }

// transformar aplica fn a cada elemento y retorna el nuevo slice.
// func transformar(nums []int, fn func(int) int) []int {
// 	resultado := make([]int, len(nums))
// 	for i, n := range nums {
// 		resultado[i] = fn(n)
// 	}
// 	return resultado
// }

// reducir combina todos los elementos con fn, partiendo de inicial.
// func reducir(nums []int, inicial int, fn func(int, int) int) int {
// 	acc := inicial
// 	for _, n := range nums {
// 		acc = fn(acc, n)
// 	}
// 	return acc
// }

// PASO 4 — Descomenta este bloque:
// conLogger envuelve fn con logging antes y después.
// Retorna una función con la misma firma — el middleware pattern.
// func conLogger(nombre string, fn func() error) func() error {
// 	return func() error {
// 		fmt.Printf("→ %s\n", nombre)
// 		err := fn()
// 		if err != nil {
// 			fmt.Printf("✗ %s falló: %v\n", nombre, err)
// 			return err
// 		}
// 		fmt.Printf("✓ %s completado\n", nombre)
// 		return nil
// 	}
// }

func main() {
	fmt.Println("=== Práctica 02: Closures y HOF ===")
	fmt.Println()

	// ============================================
	// PASO 1: Closure como contador
	// ============================================
	fmt.Println("--- Paso 1: Closure contador ---")

	// Cada llamada a contadorNuevo() crea un entorno independiente
	// Descomenta las siguientes líneas:
	// contarA := contadorNuevo()
	// contarB := contadorNuevo() // su propio 'n', no comparte con A

	// fmt.Println("contarA:", contarA()) // 1
	// fmt.Println("contarA:", contarA()) // 2
	// fmt.Println("contarB:", contarB()) // 1 — entorno independiente
	// fmt.Println("contarA:", contarA()) // 3 — A continúa desde donde estaba

	fmt.Println()

	// ============================================
	// PASO 2: Fábrica de funciones
	// ============================================
	fmt.Println("--- Paso 2: Fábrica de funciones ---")

	// multiplicadorPor retorna un closure con 'factor' capturado
	// Descomenta las siguientes líneas:
	// triplicar := multiplicadorPor(3)
	// fmt.Printf("triplicar(4) = %d\n", triplicar(4))  // 12
	// fmt.Printf("triplicar(7) = %d\n", triplicar(7))  // 21

	// prefijador retorna un closure con 'prefijo' capturado
	// Descomenta las siguientes líneas:
	// advertir := prefijador("[WARN] ")
	// informar := prefijador("[INFO] ")
	// fmt.Println("prefijo WARN:", advertir("disco lleno"))
	// fmt.Println("prefijo INFO:", informar("inicio ok"))

	fmt.Println()

	// ============================================
	// PASO 3: HOF — filtrar, transformar, reducir
	// ============================================
	fmt.Println("--- Paso 3: HOF ---")

	// Las funciones anónimas se pasan directamente como predicados
	// Descomenta las siguientes líneas:
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println("Originales:", nums)

	// Filtrar pares — predicado: n%2 == 0
	// pares := filtrar(nums, func(n int) bool { return n%2 == 0 })
	// fmt.Println("Pares:     ", pares)

	// Transformar: elevar al cuadrado — fn: n*n
	// cuadrados := transformar(pares, func(n int) int { return n * n })
	// fmt.Println("Cuadrados: ", cuadrados)

	// Reducir: sumar todos — fn: acumulador + n
	// suma := reducir(cuadrados, 0, func(acc, n int) int { return acc + n })
	// fmt.Printf("Suma total: %d\n", suma)

	fmt.Println()

	// ============================================
	// PASO 4: Closure como middleware
	// ============================================
	fmt.Println("--- Paso 4: Middleware con closure ---")

	// conLogger envuelve cualquier func() error con logging
	// Descomenta las siguientes líneas:
	// procesarPedido := conLogger("procesarPedido", func() error {
	// 	// simulación de operación exitosa
	// 	return nil
	// })
	// procesarPedido()

	// validarStock simula un error
	// validarStock := conLogger("validarStock", func() error {
	// 	return fmt.Errorf("stock insuficiente")
	// })
	// validarStock()

	fmt.Println()
}
