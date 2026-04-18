package main

// Paquete main: práctica 02 — Slices: make, append, aliasing y copy
//   Qué     → explorar la estructura interna del slice y sus operaciones clave
//   Para qué → escribir código correcto sin bugs de aliasing ni reallocations costosas

import "fmt"

func main() {
	// ============================================
	// PASO 1: make, append y len/cap
	// ============================================
	fmt.Println("--- Paso 1: make, append, len/cap ---")

	// make([]T, len, cap) preasigna capacidad.
	// append crece automáticamente cuando len == cap.

	// Descomenta las siguientes líneas:
	// s := make([]int, 0, 5)
	// fmt.Printf("después de make:   len=%d cap=%d %v\n", len(s), cap(s), s)
	//
	// s = append(s, 1, 2, 3)
	// fmt.Printf("después de append 3: len=%d cap=%d %v\n", len(s), cap(s), s)
	//
	// // Al superar cap, Go alloca nuevo backing array (cap puede doblar)
	// s = append(s, 4, 5, 6)
	// fmt.Printf("después de append 3 más: len=%d cap=%d %v\n", len(s), cap(s), s)
	//
	// // append de otro slice completo con operador ...
	// extra := []int{7, 8, 9}
	// s = append(s, extra...)
	// fmt.Printf("después de append slice: len=%d cap=%d %v\n", len(s), cap(s), s)
	//
	// // nil slice — válido, append funciona sobre él
	// var nulo []int
	// fmt.Printf("nil slice: len=%d cap=%d nil=%v\n", len(nulo), cap(nulo), nulo == nil)
	// nulo = append(nulo, 42)
	// fmt.Println("nil slice tras append:", nulo)

	fmt.Println()

	// ============================================
	// PASO 2: Aliasing — a[lo:hi] comparte backing array
	// ============================================
	fmt.Println("--- Paso 2: Aliasing ---")

	// a[lo:hi] crea un nuevo header pero COMPARTE el backing array.
	// Modificar el sub-slice modifica el slice original.

	// Descomenta las siguientes líneas:
	// nums := []int{10, 20, 30, 40, 50}
	// sub := nums[1:4] // ptr apunta a nums[1], len=3, cap=4
	//
	// fmt.Println("nums antes:", nums)
	// fmt.Println("sub:       ", sub)
	// fmt.Printf("sub: len=%d cap=%d\n", len(sub), cap(sub))
	//
	// // Modificar sub también modifica nums
	// sub[0] = 99
	// fmt.Println("sub[0] = 99 →")
	// fmt.Println("nums después:", nums) // [10 99 30 40 50] — ¡aliasing!
	// fmt.Println("sub después: ", sub)  // [99 30 40]
	//
	// // Three-index slice para limitar cap del sub-slice
	// nums2 := []int{10, 20, 30, 40, 50}
	// seguro := nums2[1:4:4] // cap limitado a 3 — append fuerza nuevo backing array
	// fmt.Printf("seguro: len=%d cap=%d\n", len(seguro), cap(seguro))
	// seguro = append(seguro, 999) // nuevo backing array — nums2 no cambia
	// fmt.Println("nums2 (sin cambios):", nums2)

	fmt.Println()

	// ============================================
	// PASO 3: copy — slice independiente
	// ============================================
	fmt.Println("--- Paso 3: copy ---")

	// copy(dst, src) copia min(len(dst), len(src)) elementos.
	// dst y src no comparten backing array después de copy.

	// Descomenta las siguientes líneas:
	// nums := []int{10, 20, 30, 40, 50}
	//
	// // Clonar un rango del slice de forma independiente
	// independiente := make([]int, 3)
	// n := copy(independiente, nums[1:4])
	// fmt.Printf("copiados: %d elementos\n", n)
	// fmt.Println("independiente:", independiente) // [20 30 40]
	//
	// // Modificar independiente NO afecta a nums
	// independiente[0] = 99
	// fmt.Println("nums (sin cambios):", nums)    // [10 20 30 40 50]
	// fmt.Println("independiente:     ", independiente) // [99 30 40]
	//
	// // Alternativa idiomática para clonar un slice completo
	// clon := append([]int{}, nums...)
	// clon[0] = 777
	// fmt.Println("nums (sin cambios):", nums) // [10 20 30 40 50]
	// fmt.Println("clon:              ", clon) // [777 20 30 40 50]

	fmt.Println()

	// ============================================
	// PASO 4: Filtrado idiomático con nil slice
	// ============================================
	fmt.Println("--- Paso 4: Filtrado ---")

	// nil slice + append: patrón idiomático para filtrar sin alloación previa.

	// Descomenta las siguientes líneas:
	// datos := []int{3, -1, 7, 0, -4, 9, 2, -8, 5}
	//
	// pares := filtrarPares(datos)
	// fmt.Println("datos: ", datos)
	// fmt.Println("pares: ", pares) // [0 -4 2 -8]
	//
	// positivos := filtrarPositivos(datos)
	// fmt.Println("positivos:", positivos) // [3 7 9 2 5]

	fmt.Println()

	// ============================================
	// PASO 5: Slice de slices — matriz 2D
	// ============================================
	fmt.Println("--- Paso 5: Matriz 2D ---")

	// [][]T es la forma idiomática de una matriz en Go.
	// Cada fila es un slice independiente.

	// Descomenta las siguientes líneas:
	// filas, cols := 3, 4
	// matriz := make([][]int, filas)
	// for i := range matriz {
	//     matriz[i] = make([]int, cols)
	// }
	//
	// // Llenar con valores (fila * 10 + columna)
	// for i := range matriz {
	//     for j := range matriz[i] {
	//         matriz[i][j] = i*10 + j
	//     }
	// }
	//
	// imprimirMatriz(matriz)
	//
	// // Cada fila puede tener longitud diferente (slice irregular)
	// triangulo := [][]int{
	//     {1},
	//     {1, 2},
	//     {1, 2, 3},
	//     {1, 2, 3, 4},
	// }
	// fmt.Println("\nTriángulo:")
	// imprimirMatriz(triangulo)

	fmt.Println()
}

// filtrarPares retorna un nuevo slice con solo los números pares.
// func filtrarPares(nums []int) []int {
//     var resultado []int
//     for _, n := range nums {
//         if n%2 == 0 {
//             resultado = append(resultado, n)
//         }
//     }
//     return resultado
// }

// filtrarPositivos retorna un nuevo slice con solo los números positivos.
// func filtrarPositivos(nums []int) []int {
//     var resultado []int
//     for _, n := range nums {
//         if n > 0 {
//             resultado = append(resultado, n)
//         }
//     }
//     return resultado
// }

// imprimirMatriz imprime una matriz 2D fila a fila.
// func imprimirMatriz(m [][]int) {
//     for i, fila := range m {
//         fmt.Printf("fila %d: %v\n", i, fila)
//     }
// }
