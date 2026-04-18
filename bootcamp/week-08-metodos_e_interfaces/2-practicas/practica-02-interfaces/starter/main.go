package main

import "fmt"

// ============================================
// PASO 1: Declarar la interface Medible
// ============================================
// Descomenta cuando llegues al Paso 1:
// type Medible interface {
// 	Medir() string
// }

// ============================================
// PASO 2: Dos tipos que satisfacen Medible
// ============================================
// Descomenta los dos tipos y sus métodos (Paso 2):

// type Distancia struct{ metros float64 }
// type Masa struct{ kg float64 }

// func (d Distancia) Medir() string { return fmt.Sprintf("%.2f m", d.metros) }
// func (m Masa) Medir() string      { return fmt.Sprintf("%.2f kg", m.kg) }

// ============================================
// PASO 3: Función polimórfica
// ============================================
// Descomenta esta función cuando llegues al Paso 3:
// func mostrar(m Medible) {
// 	fmt.Println("Medición:", m.Medir())
// }

func main() {
	// ============================================
	// PASOS 1 + 2: interface y tipos
	// ============================================
	fmt.Println("--- Paso 1 + 2: tipos implementan Medible ---")

	// Distancia y Masa satisfacen Medible implícitamente
	// Descomenta (con los tipos y la interface arriba):
	// d := Distancia{metros: 100}
	// masa := Masa{kg: 75}
	// fmt.Println("Medición:", d.Medir())
	// fmt.Println("Medición:", masa.Medir())

	fmt.Println()

	// ============================================
	// PASO 3: Función polimórfica y slice de interface
	// ============================================
	fmt.Println("--- Paso 3: slice de interface ---")

	// Descomenta (con mostrar y los tipos arriba):
	// elementos := []Medible{Distancia{100}, Masa{75}}
	// for _, e := range elementos {
	// 	mostrar(e)
	// }

	fmt.Println()

	// ============================================
	// PASO 4: any y type assertion
	// ============================================
	fmt.Println("--- Paso 4: any y type assertion ---")

	// any acepta cualquier valor pero pierde el tipo concreto
	// Descomenta (con los tipos arriba):
	// var v any = Distancia{metros: 50}
	// // v.Medir() no compila — any no tiene Medir()
	// if d, ok := v.(Distancia); ok {
	// 	fmt.Println("Type assertion exitosa:", d.Medir())
	// }
}
