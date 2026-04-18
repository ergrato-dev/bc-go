package main

import "fmt"

// Rectangulo tiene ancho y alto en float64
type Rectangulo struct {
	ancho, alto float64
}

// ============================================
// PASO 1: Método de valor — Area
// ============================================
// Descomenta este método cuando llegues al Paso 1:
// func (r Rectangulo) Area() float64 {
// 	return r.ancho * r.alto
// }

// ============================================
// PASO 2: Método de puntero — Escalar
// ============================================
// Descomenta este método cuando llegues al Paso 2:
// func (r *Rectangulo) Escalar(factor float64) {
// 	r.ancho *= factor
// 	r.alto *= factor
// }

// ============================================
// PASO 3: fmt.Stringer — String()
// ============================================
// Descomenta este método cuando llegues al Paso 3:
// func (r Rectangulo) String() string {
// 	return fmt.Sprintf("Rect(%.1f x %.1f)", r.ancho, r.alto)
// }

// ============================================
// PASO 4: Verificación de interface en compilación
// ============================================
// Descomenta cuando llegues al Paso 4 (y tengas String() descomentado):
// var _ fmt.Stringer = Rectangulo{}

func main() {
	// ============================================
	// PASO 1: Método con receptor de valor
	// ============================================
	fmt.Println("--- Paso 1: método de valor ---")

	// Descomenta las siguientes líneas (y el método Area arriba):
	// r := Rectangulo{ancho: 3, alto: 4}
	// fmt.Printf("Área: %.2f\n", r.Area())

	fmt.Println()

	// ============================================
	// PASO 2: Método con receptor de puntero
	// ============================================
	fmt.Println("--- Paso 2: método de puntero ---")

	// Escalar modifica r directamente — usa receptor *Rectangulo
	// Descomenta (y el método Escalar arriba):
	// r2 := Rectangulo{ancho: 3, alto: 4}
	// r2.Escalar(2) // Go convierte a (&r2).Escalar(2)
	// fmt.Printf("Después de escalar x2: %.1f x %.1f\n", r2.ancho, r2.alto)

	fmt.Println()

	// ============================================
	// PASO 3: fmt.Stringer
	// ============================================
	fmt.Println("--- Paso 3: fmt.Stringer ---")

	// Si Rectangulo tiene String() string, fmt.Println lo usa
	// Descomenta (y el método String y Escalar arriba):
	// r3 := Rectangulo{ancho: 3, alto: 4}
	// r3.Escalar(2)
	// fmt.Println(r3) // usa String() automáticamente

	fmt.Println()

	// ============================================
	// PASO 4: Verificación de interface
	// ============================================
	fmt.Println("--- Paso 4: verificación de interface ---")

	// La verificación ocurre al compilar, no aquí
	// Descomenta la línea var _ fmt.Stringer = Rectangulo{} arriba
	// Si String() está implementada, compilará sin error
	// fmt.Println("Stringer verificado en compilación")
}
