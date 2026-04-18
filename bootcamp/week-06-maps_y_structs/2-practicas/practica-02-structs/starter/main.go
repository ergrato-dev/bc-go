package main

import "fmt"

// Producto representa un artículo del catálogo.
// La definición del tipo va fuera de main — nivel de paquete.
type Producto struct {
	ID      int
	Nombre  string
	Precio  float64
	EnStock bool
}

// String implementa fmt.Stringer.
// Go llamará este método automáticamente al imprimir con fmt.Println.
// Descomenta cuando llegues al Paso 3:
// func (p Producto) String() string {
// 	disponible := "sin stock"
// 	if p.EnStock {
// 		disponible = "disponible"
// 	}
// 	return fmt.Sprintf("[%d] %-20s $%.2f (%s)",
// 		p.ID, p.Nombre, p.Precio, disponible)
// }

func main() {
	// ============================================
	// PASO 1: Struct con zero value
	// ============================================
	fmt.Println("--- Paso 1: Zero value ---")

	// var declara el struct con todos los campos en su zero value.
	// Descomenta las siguientes líneas:
	// var p Producto
	// fmt.Printf("Nombre vacío: %v\n", p.Nombre == "")
	// fmt.Printf("Precio cero:  %v\n", p.Precio == 0)
	// fmt.Printf("En stock:     %v\n", p.EnStock)

	fmt.Println()

	// ============================================
	// PASO 2: Literal con nombre de campo
	// ============================================
	fmt.Println("--- Paso 2: Literal ---")

	// Siempre nombrar los campos al inicializar un struct.
	// Los campos omitidos quedan en su zero value.
	// Descomenta las siguientes líneas:
	// p1 := Producto{
	// 	ID:      1,
	// 	Nombre:  "Cuaderno A4",
	// 	Precio:  2.50,
	// 	EnStock: true,
	// }
	// p2 := Producto{
	// 	ID:     2,
	// 	Nombre: "Lápiz",
	// 	// Precio y EnStock quedan en 0.0 y false
	// }
	//
	// fmt.Printf("p1: %s $%.2f\n", p1.Nombre, p1.Precio)
	// fmt.Printf("p2 en stock: %v\n", p2.EnStock)

	fmt.Println()

	// ============================================
	// PASO 3: Método de valor (String)
	// ============================================
	fmt.Println("--- Paso 3: Método de valor ---")

	// Descomenta primero el método String() arriba del main.
	// Luego descomenta las siguientes líneas:
	// p1 := Producto{ID: 1, Nombre: "Cuaderno A4", Precio: 2.50, EnStock: true}
	// p2 := Producto{ID: 2, Nombre: "Lápiz"}
	//
	// fmt.Println(p1) // Go llama p1.String() automáticamente
	// fmt.Println(p2)

	fmt.Println()

	// ============================================
	// PASO 4: Semántica de valor y comparación ==
	// ============================================
	fmt.Println("--- Paso 4: Semántica de valor ---")

	// Asignar un struct crea una copia — las variables son independientes.
	// Descomenta las siguientes líneas:
	// a := Producto{ID: 1, Nombre: "Cuaderno A4", Precio: 2.50, EnStock: true}
	// b := a // copia completa de todos los campos
	// b.Nombre = "Cuaderno A5"
	// b.Precio = 3.00
	//
	// fmt.Println("a.Nombre:", a.Nombre) // Cuaderno A4
	// fmt.Println("b.Nombre:", b.Nombre) // Cuaderno A5
	//
	// c := Producto{ID: 1, Nombre: "Cuaderno A4", Precio: 2.50, EnStock: true}
	// fmt.Println("a == c:", a == c) // true
	// fmt.Println("a == b:", a == b) // false

	fmt.Println()
}
