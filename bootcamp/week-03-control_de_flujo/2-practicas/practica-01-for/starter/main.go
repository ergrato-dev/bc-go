package main
// Paquete main: práctica 01 — Patrones de for
//   Qué     → explorar los 4 patrones del bucle for en Go
//   Para qué → dominar la única estructura de bucle del lenguaje
//   Impacto  → Go no tiene while ni foreach; todo se expresa con for


import "fmt"

func main() {
	// ============================================
	// PASO 1: for clásico (init; condición; post)
	// Qué     → bucle con tres componentes explícitos
	// Para qué → control total sobre inicio, condición y avance
	// Impacto  → no requiere paréntesis alrededor de la condición
	// ============================================
	fmt.Println("--- Paso 1: for clásico ---")

	// El for clásico tiene la misma forma que en C/Java pero sin paréntesis.
	// i se declara dentro del for y solo existe durante el bucle.
	// Descomenta las siguientes líneas:
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("i = %d\n", i)
	// }

	fmt.Println()

	// ============================================
	// PASO 2: for como while (solo condición)
	// Qué     → for sin init ni post
	// Para qué → expresar bucles con condición dinámica
	// Impacto  → equivalente exacto a while en otros lenguajes
	// ============================================
	fmt.Println("--- Paso 2: for como while ---")

	// Cuando se omiten init y post, for actúa como while.
	// La variable se declara fuera porque su valor importa después del bucle.
	// Descomenta las siguientes líneas:
	// saldo := 30
	// for saldo > 0 {
	// 	saldo -= 10
	// 	fmt.Printf("Saldo restante: %d\n", saldo)
	// }
	// fmt.Println("Saldo agotado")

	fmt.Println()

	// ============================================
	// PASO 3: for range sobre slice
	// Qué     → iteración idiomática con range
	// Para qué → obtener índice y valor sin gestionar el contador
	// Impacto  → usar _ para ignorar el índice o el valor no necesario
	// ============================================
	fmt.Println("--- Paso 3: for range sobre slice ---")

	// range retorna (índice, valor) en cada iteración.
	// Descomenta las siguientes líneas:
	// planetas := []string{"Mercurio", "Venus", "Tierra", "Marte"}
	//
	// // con índice y valor
	// for i, planeta := range planetas {
	// 	fmt.Printf("[%d] %s\n", i, planeta)
	// }
	//
	// fmt.Println()
	//
	// // solo valor (índice ignorado con _)
	// for _, planeta := range planetas {
	// 	fmt.Println("Planeta:", planeta)
	// }

	fmt.Println()

	// ============================================
	// PASO 4: for range sobre string (runes Unicode)
	// Qué     → range sobre string itera por runes, no bytes
	// Para qué → manejar texto con caracteres no-ASCII (tildes, CJK, emojis)
	// Impacto  → la posición es en bytes; para 'é' el salto es de 2 bytes
	// ============================================
	fmt.Println("--- Paso 4: for range sobre string ---")

	// Al usar range sobre un string, Go entrega:
	//   - posición: índice en bytes dentro del string
	//   - caracter: el rune (int32) en esa posición
	// Descomenta las siguientes líneas:
	// for posicion, caracter := range "café" {
	// 	fmt.Printf("byte[%d] = %c  (rune %d)\n", posicion, caracter, caracter)
	// }
	// fmt.Println()
	// // Comparar longitud en bytes vs número de runes:
	// palabra := "café"
	// fmt.Printf("len(bytes) = %d  |  runes = %d\n", len(palabra), len([]rune(palabra)))

	fmt.Println()

	// ============================================
	// PASO 5: break y continue
	// Qué     → control de flujo dentro del bucle
	// Para qué → break termina el bucle; continue salta a la siguiente iteración
	// Impacto  → evitar código innecesario después de encontrar lo que se busca
	// ============================================
	fmt.Println("--- Paso 5: break y continue ---")

	// break sale del bucle en cuanto se cumple la condición.
	// Descomenta las siguientes líneas:
	// fmt.Print("Primer múltiplo de 7 entre 1 y 100: ")
	// for i := 1; i <= 100; i++ {
	// 	if i%7 == 0 {
	// 		fmt.Println(i)
	// 		break // sale del bucle; no sigue buscando
	// 	}
	// }
	//
	// // continue salta al inicio de la siguiente iteración.
	// fmt.Print("Impares del 1 al 10: ")
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue // saltar los pares
	// 	}
	// 	fmt.Printf("%d ", i)
	// }
	// fmt.Println()
}
