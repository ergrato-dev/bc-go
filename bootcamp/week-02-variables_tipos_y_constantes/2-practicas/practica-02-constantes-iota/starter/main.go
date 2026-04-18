package main
// Paquete main: punto de entrada de la práctica 02 — Constantes e iota en Go
//   Qué     → programa guiado para explorar const, iota y patrones de flags
//   Para qué → aprender a crear enumeraciones idiomáticas y flags combinables
//   Impacto  → descomenta cada sección siguiendo el README; go run . después de cada paso


import "fmt"

// ============================================
// Declaraciones de tipos e iota (scope de paquete)
// ============================================
// Los tipos personalizados y sus constantes deben declararse fuera de main
// para que sean accesibles en todo el paquete.

// Season representa las estaciones del año.
// Qué     → tipo personalizado sobre int para el enum Season
// Para qué → el compilador distingue Season de int; no puedes mezclar accidentalmente
// Impacto  → seguridad de tipos en compilación: no puedes pasar un int arbitrario como Season

// Descomenta el bloque de tipo Season para el Paso 3:
// type Season int
//
// const (
//     Spring Season = iota // 0
//     Summer               // 1
//     Autumn               // 2
//     Winter               // 3
// )

// AccessFlag representa permisos de acceso como flags de bits.
// Qué     → tipo uint8 para permisos combinables con OR a nivel de bits
// Para qué → representar múltiples permisos en un solo byte (hasta 8 permisos)
// Impacto  → usar & para verificar si un permiso está activo sin condiciones complejas

// Descomenta el bloque de tipo AccessFlag para el Paso 4:
// type AccessFlag uint8
//
// const (
//     CanRead    AccessFlag = 1 << iota // 1 << 0 = 0b00000001 = 1
//     CanWrite                          // 1 << 1 = 0b00000010 = 2
//     CanDelete                         // 1 << 2 = 0b00000100 = 4
//     CanPublish                        // 1 << 3 = 0b00001000 = 8
// )

func main() {
	// ============================================
	// PASO 1: Constantes simples y bloque const
	// ============================================
	fmt.Println("--- Paso 1: Bloque const ---")

	// Un bloque const agrupa constantes relacionadas.
	// Son evaluadas en tiempo de compilación: no pueden ser el resultado de una función.
	// El compilador puede incrustarlas directamente en el código (sin memoria en heap).

	// Descomenta las siguientes líneas:
	// const (
	//     appVersion  = "2.0.0"
	//     maxRetries  = 3
	//     defaultPort = 8080
	//     timeoutSecs = 30
	// )
	//
	// fmt.Printf("App: %s\n", appVersion)
	// fmt.Printf("Puerto por defecto: %d\n", defaultPort)
	// fmt.Printf("Reintentos máximos: %d\n", maxRetries)
	// fmt.Printf("Timeout: %d segundos\n", timeoutSecs)

	fmt.Println()

	// ============================================
	// PASO 2: Constantes tipadas vs no tipadas
	// ============================================
	fmt.Println("--- Paso 2: Tipada vs no tipada ---")

	// Constante no tipada: se adapta al contexto de uso (flexible).
	// Constante tipada: solo compatible con su tipo exacto (segura).

	// Descomenta las siguientes líneas:
	// const maxItems = 100 // no tipada
	//
	// var a int32 = maxItems // ✅ compatible directamente
	// var b int64 = maxItems // ✅ también compatible
	// fmt.Printf("int32: %d | int64: %d\n", a, b)
	//
	// const typedLimit int32 = 500 // tipada
	// var c int64 = int64(typedLimit) // requiere conversión explícita
	// fmt.Printf("tipada como int64 (conversión): %d\n", c)

	fmt.Println()

	// ============================================
	// PASO 3: iota — enumeración con tipo personalizado
	// ============================================
	fmt.Println("--- Paso 3: iota básico ---")

	// iota empieza en 0 en cada bloque const y se incrementa en 1.
	// El tipo personalizado (type Season int) añade seguridad en compilación.
	// Recuerda descomenta también la declaración de tipo arriba (fuera de main).

	// Descomenta las siguientes líneas:
	// current := Autumn
	// fmt.Printf("Estación:   %d\n",  current)  // 2
	// fmt.Printf("Tipo:       %T\n",  current)  // main.Season
	// fmt.Printf("Spring=%d Summer=%d Autumn=%d Winter=%d\n",
	//     Spring, Summer, Autumn, Winter)

	fmt.Println()

	// ============================================
	// PASO 4: iota con desplazamiento de bits — flags
	// ============================================
	fmt.Println("--- Paso 4: Flags de permisos ---")

	// 1 << iota genera 1, 2, 4, 8 (potencias de 2).
	// Cada constante tiene exactamente un bit activo → permisos combinables con |.
	// Recuerda descomenta también la declaración de tipo arriba (fuera de main).

	// Descomenta las siguientes líneas:
	// fmt.Printf("CanRead:    %04b = %d\n", CanRead,    CanRead)
	// fmt.Printf("CanWrite:   %04b = %d\n", CanWrite,   CanWrite)
	// fmt.Printf("CanDelete:  %04b = %d\n", CanDelete,  CanDelete)
	// fmt.Printf("CanPublish: %04b = %d\n", CanPublish, CanPublish)
	//
	// // Editor: puede leer, escribir y publicar (pero no borrar)
	// editor := CanRead | CanWrite | CanPublish
	// fmt.Printf("\nEditor flags: %04b\n", editor)
	// fmt.Printf("¿Puede leer?:    %v\n", editor&CanRead    != 0)
	// fmt.Printf("¿Puede escribir? %v\n", editor&CanWrite   != 0)
	// fmt.Printf("¿Puede borrar?:  %v\n", editor&CanDelete  != 0)
	// fmt.Printf("¿Puede publicar? %v\n", editor&CanPublish != 0)
}
