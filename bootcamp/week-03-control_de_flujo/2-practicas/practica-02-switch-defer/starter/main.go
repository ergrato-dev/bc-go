// Paquete main: práctica 02 — switch y defer
//   Qué     → explorar switch con múltiples valores por caso y el modelo LIFO de defer
//   Para qué → dominar la selección de casos y la gestión idiomática de recursos
//   Impacto  → switch no necesita break; defer garantiza limpieza incluso con errores

package main

import "fmt"

// ============================================
// FUNCIÓN auxiliar: clasificarLetra
// Qué     → clasifica un rune como vocal, dígito o consonante
// Para qué → mostrar switch con múltiples valores por caso
// Impacto  → un case puede listar varios valores separados por coma
// ============================================

// Descomenta la función completa:
// func clasificarLetra(r rune) string {
// 	switch r {
// 	case 'a', 'e', 'i', 'o', 'u',
// 		'A', 'E', 'I', 'O', 'U':
// 		return "vocal"
// 	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
// 		return "dígito"
// 	default:
// 		return "consonante u otro"
// 	}
// }

// ============================================
// FUNCIÓN auxiliar: demostrarLIFO
// Qué     → declara tres defer y muestra el orden de ejecución
// Para qué → visualizar el stack LIFO de defer
// Impacto  → último declarado = primero en ejecutar
// ============================================

// Descomenta la función completa:
// func demostrarLIFO() {
// 	fmt.Println("Iniciando función...")
// 	defer fmt.Println("  [defer 1] primero declarado → último en ejecutar")
// 	defer fmt.Println("  [defer 2] segundo declarado → segundo en ejecutar")
// 	defer fmt.Println("  [defer 3] tercero declarado → primero en ejecutar")
// 	fmt.Println("Función a punto de retornar")
// }

func main() {
	// ============================================
	// PASO 1: switch con expresión y múltiples valores por caso
	// Qué     → case puede listar varios valores separados por coma
	// Para qué → reemplazar cadenas largas de if/else con código más claro
	// Impacto  → no se necesita break; Go no hace fall-through implícito
	// ============================================
	fmt.Println("--- Paso 1: switch con múltiples valores por caso ---")

	// Descomenta las siguientes líneas:
	// letras := []rune{'a', 'b', 'E', '5', 'z', 'O'}
	// for _, l := range letras {
	// 	fmt.Printf("'%c' → %s\n", l, clasificarLetra(l))
	// }

	fmt.Println()

	// ============================================
	// PASO 2: switch sin expresión (como if/else)
	// Qué     → switch evaluando condiciones booleanas por caso
	// Para qué → reemplazar if/else if largas con una estructura más legible
	// Impacto  → cada case es una condición booleana independiente
	// ============================================
	fmt.Println("--- Paso 2: switch sin expresión ---")

	// Descomenta las siguientes líneas:
	// temperaturas := []float64{36.5, 37.6, 38.1, 39.0, 40.5}
	// for _, t := range temperaturas {
	// 	var estado string
	// 	switch {
	// 	case t >= 40.0:
	// 		estado = "fiebre alta"
	// 	case t >= 38.0:
	// 		estado = "fiebre moderada"
	// 	case t >= 37.5:
	// 		estado = "temperatura elevada"
	// 	default:
	// 		estado = "normal"
	// 	}
	// 	fmt.Printf("%.1f°C → %s\n", t, estado)
	// }

	fmt.Println()

	// ============================================
	// PASO 3: defer — orden LIFO
	// Qué     → múltiples defer se ejecutan en orden inverso al de declaración
	// Para qué → garantiza que los recursos se liberen en orden correcto
	// Impacto  → el último defer declarado es el primero en ejecutarse
	// ============================================
	fmt.Println("--- Paso 3: defer LIFO ---")

	// Descomenta la siguiente línea:
	// demostrarLIFO()

	fmt.Println()

	// ============================================
	// PASO 4: defer — evaluación de argumentos
	// Qué     → los argumentos de defer se evalúan al declararlo, no al ejecutarlo
	// Para qué → entender qué valor captura defer vs un closure
	// Impacto  → para capturar el valor final, usar defer func() { ... }()
	// ============================================
	fmt.Println("--- Paso 4: defer y evaluación de argumentos ---")

	// Argumento evaluado en el momento de la declaración.
	// Descomenta las siguientes líneas:
	// x := 10
	// defer fmt.Printf("  arg evaluado al declarar defer: %d\n", x) // captura 10
	// x = 99
	// fmt.Printf("  valor de x al retornar: %d\n", x)
	// // ↑ El defer imprimirá 10, no 99

	fmt.Println()

	// Closure captura la variable, no el valor inmediato.
	// Descomenta las siguientes líneas:
	// y := 10
	// defer func() {
	// 	fmt.Printf("  closure lee y al ejecutar defer: %d\n", y) // lee 99
	// }()
	// y = 99
	// fmt.Printf("  valor de y al retornar: %d\n", y)
	// // ↑ El defer del closure imprimirá 99
}
