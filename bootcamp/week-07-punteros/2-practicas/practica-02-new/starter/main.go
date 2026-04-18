package main

import "fmt"

// Caja representa un contenedor con un string
type Caja struct {
	contenido string
}

// ============================================
// PASO 2: Método de valor — no modifica
// ============================================
// Descomenta este método cuando llegues al Paso 2:
// func (c Caja) vaciar() {
// 	c.contenido = "" // solo afecta la copia local
// }

// ============================================
// PASO 3: Método de puntero — sí modifica
// ============================================
// Descomenta este método cuando llegues al Paso 3:
// func (c *Caja) llenar(v string) {
// 	c.contenido = v // modifica el struct original
// }

// Punto es un struct con dos coordenadas
// (usado en el Paso 4)
type Punto struct {
	X, Y int
}

func main() {
	// ============================================
	// PASO 1: new(T) — puntero al zero value
	// ============================================
	fmt.Println("--- Paso 1: new(T) ---")

	// new(int) asigna un int en zero value (0) y retorna *int
	// Descomenta las siguientes líneas:
	// p := new(int)
	// fmt.Printf("*p == %d al inicio\n", *p)
	// *p = 42
	// fmt.Printf("*p == %d después de asignar\n", *p)

	fmt.Println()

	// ============================================
	// PASO 2: Método de valor — no modifica el original
	// ============================================
	fmt.Println("--- Paso 2: método de valor (no modifica) ---")

	// c Caja en el receptor significa que vaciar recibe una copia
	// Descomenta las siguientes líneas (y el método vaciar arriba):
	// caja := Caja{contenido: "lleno"}
	// fmt.Println("antes de vaciar:", caja.contenido)
	// caja.vaciar() // trabaja sobre una copia — no cambia caja
	// fmt.Println("después de vaciar:", caja.contenido) // sigue "lleno"

	fmt.Println()

	// ============================================
	// PASO 3: Método de puntero — sí modifica el original
	// ============================================
	fmt.Println("--- Paso 3: método de puntero (sí modifica) ---")

	// *Caja en el receptor significa que llenar accede al struct original
	// Go aplica auto-deref: caja.llenar() → (&caja).llenar()
	// Descomenta las siguientes líneas (y el método llenar arriba):
	// caja2 := Caja{}
	// fmt.Println("antes de llenar:", caja2.contenido)
	// caja2.llenar("arroz")
	// fmt.Println("después de llenar:", caja2.contenido) // "arroz"

	fmt.Println()

	// ============================================
	// PASO 4: Auto-deref — sin (*p).campo
	// ============================================
	fmt.Println("--- Paso 4: auto-deref ---")

	// Con un puntero a struct, puedes acceder a campos directamente
	// Go convierte p.X en (*p).X automáticamente
	// Descomenta las siguientes líneas:
	// pt := &Punto{X: 1, Y: 2}
	// pt.X = 10        // equivale a (*pt).X = 10
	// fmt.Println("p.X ==", pt.X)
}
