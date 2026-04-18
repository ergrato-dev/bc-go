package main

import "fmt"

func main() {
	// ============================================
	// PASO 1: Declarar y escribir en un map
	// ============================================
	fmt.Println("--- Paso 1: Declarar y escribir ---")

	// make(map[K]V) crea un map vacío listo para escritura.
	// Descomenta las siguientes líneas:
	// stock := make(map[string]int)
	// stock["lapiz"]    = 50
	// stock["cuaderno"] = 20
	// stock["goma"]     = 0 // asignado explícitamente como 0
	//
	// fmt.Printf("Stock lapiz:    %d\n", stock["lapiz"])
	// fmt.Printf("Stock goma:     %d\n", stock["goma"])
	// fmt.Printf("Stock regla:    %d\n", stock["regla"]) // clave inexistente → 0

	fmt.Println()

	// ============================================
	// PASO 2: Patrón v, ok para verificar existencia
	// ============================================
	fmt.Println("--- Paso 2: Patrón v, ok ---")

	// El patrón v, ok distingue "clave inexistente" de "clave con zero value".
	// Descomenta las siguientes líneas (requiere stock del paso anterior):
	// cantidad, ok := stock["regla"]
	// if !ok {
	// 	fmt.Println("regla no está registrada")
	// } else {
	// 	fmt.Printf("regla tiene %d unidades\n", cantidad)
	// }
	//
	// if cant, ok := stock["goma"]; ok {
	// 	fmt.Printf("goma encontrada: %d unidades\n", cant)
	// }

	fmt.Println()

	// ============================================
	// PASO 3: Iterar y contar frecuencias
	// ============================================
	fmt.Println("--- Paso 3: Frecuencias ---")

	// freq[k]++ funciona aunque la clave no exista: empieza desde el zero value 0.
	// Descomenta las siguientes líneas:
	// palabras := []string{"go", "rust", "go", "python", "go", "rust"}
	// freq := make(map[string]int)
	// for _, p := range palabras {
	// 	freq[p]++
	// }
	//
	// for lenguaje, veces := range freq {
	// 	fmt.Printf("%s: %d veces\n", lenguaje, veces)
	// }

	fmt.Println()

	// ============================================
	// PASO 4: delete y map como conjunto (set)
	// ============================================
	fmt.Println("--- Paso 4: delete y set ---")

	// delete(m, k) elimina la clave k del map. Es seguro si k no existe.
	// Descomenta las siguientes líneas (requiere stock del paso 1):
	// delete(stock, "goma")
	// if _, ok := stock["goma"]; !ok {
	// 	fmt.Println("goma eliminada correctamente")
	// }
	//
	// procesados := make(map[string]bool)
	// procesados["orden-001"] = true
	// procesados["orden-002"] = true
	//
	// if procesados["orden-001"] {
	// 	fmt.Println("orden-001 ya fue procesada")
	// }
	// if !procesados["orden-003"] {
	// 	fmt.Println("orden-003 pendiente de procesar")
	// }

	fmt.Println()
}
