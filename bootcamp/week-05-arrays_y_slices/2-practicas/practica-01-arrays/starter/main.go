package main

// Paquete main: práctica 01 — Arrays en Go
//   Qué     → declarar arrays de tamaño fijo y entender su semántica de valor
//   Para qué → diferenciar arrays de slices y saber cuándo usar cada uno

import "fmt"

func main() {
	// ============================================
	// PASO 1: Declaración e inicialización
	// ============================================
	fmt.Println("--- Paso 1: Declaración ---")

	// Go tiene tres formas de inicializar un array.
	// El tamaño es parte del tipo: [3]int ≠ [4]int.

	// Descomenta las siguientes líneas:
	// var puntuaciones [5]int
	// fmt.Println("Puntuaciones iniciales:", puntuaciones) // [0 0 0 0 0]
	//
	// puntuaciones[0] = 95
	// puntuaciones[1] = 87
	// puntuaciones[2] = 92
	// puntuaciones[3] = 78
	// puntuaciones[4] = 88
	// fmt.Println("Puntuaciones:", puntuaciones)
	//
	// colores := [3]string{"rojo", "verde", "azul"}
	// fmt.Println("Colores:", colores)
	//
	// planetas := [...]string{"Mercurio", "Venus", "Tierra", "Marte"}
	// fmt.Printf("Planetas (%d): %v\n", len(planetas), planetas)

	fmt.Println()

	// ============================================
	// PASO 2: range y estadísticas
	// ============================================
	fmt.Println("--- Paso 2: range y estadísticas ---")

	// range sobre un array devuelve (índice, copia del valor).
	// Usamos _ para ignorar el índice cuando no lo necesitamos.

	// Descomenta las siguientes líneas:
	// temps := [7]float64{22.5, 24.0, 19.3, 21.8, 25.1, 23.4, 20.7}
	//
	// var suma float64
	// for i, temp := range temps {
	//     fmt.Printf("Día %d: %.1f°C\n", i+1, temp)
	//     suma += temp
	// }
	// promedio := suma / float64(len(temps))
	// fmt.Printf("Promedio semanal: %.2f°C\n", promedio)
	//
	// // Ignorar índice con _ cuando solo necesitas el valor
	// var maximo float64
	// for _, temp := range temps {
	//     if temp > maximo {
	//         maximo = temp
	//     }
	// }
	// fmt.Printf("Temperatura máxima: %.1f°C\n", maximo)

	fmt.Println()

	// ============================================
	// PASO 3: Semántica de valor — copia al asignar
	// ============================================
	fmt.Println("--- Paso 3: Semántica de valor ---")

	// Asignar un array crea una copia independiente.
	// Pasar a una función también copia el array completo.
	// Para evitar la copia usa un puntero: *[n]T

	// Descomenta las siguientes líneas:
	// original := [3]int{10, 20, 30}
	// copia := original // copia total — no aliasing
	// copia[0] = 99
	//
	// fmt.Println("original:", original) // [10 20 30] — sin cambios
	// fmt.Println("copia:   ", copia)    // [99 20 30]
	//
	// // Pasar con puntero para modificar el original
	// triplicar(&original)
	// fmt.Println("original triplicado:", original) // [30 60 90]

	fmt.Println()

	// ============================================
	// PASO 4: Comparación con == y claves de mapa
	// ============================================
	fmt.Println("--- Paso 4: Comparación y claves de mapa ---")

	// Los arrays son comparables con == si sus elementos lo son.
	// Los slices NO son comparables — úsalos como claves de mapa.

	// Descomenta las siguientes líneas:
	// a := [3]int{1, 2, 3}
	// b := [3]int{1, 2, 3}
	// c := [3]int{1, 2, 4}
	//
	// fmt.Println("a == b:", a == b) // true
	// fmt.Println("a == c:", a == c) // false
	//
	// // Arrays como claves de mapa (slices no son comparables — no pueden ser claves)
	// type Coordenada [2]float64
	//
	// ciudades := map[Coordenada]string{
	//     {40.4168, -3.7038}: "Madrid",
	//     {48.8566, 2.3522}:  "París",
	//     {51.5074, -0.1278}: "Londres",
	// }
	//
	// buscar := Coordenada{40.4168, -3.7038}
	// if ciudad, ok := ciudades[buscar]; ok {
	//     fmt.Println("Ciudad encontrada:", ciudad) // Madrid
	// }

	fmt.Println()
}

// triplicar multiplica cada elemento del array por 3.
// Recibe un puntero para modificar el array original sin copiarlo.
// func triplicar(arr *[3]int) {
//     for i := range arr {
//         arr[i] *= 3
//     }
// }
