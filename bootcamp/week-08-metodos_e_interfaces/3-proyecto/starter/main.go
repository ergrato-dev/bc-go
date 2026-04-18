package main

import "fmt"

// ============================================
// INSTRUCCIONES PARA EL APRENDIZ
//
// Adapta este programa a tu dominio asignado.
// Renombra la interface y los structs a conceptos
// reales de tu dominio. Ejemplos:
//   Biblioteca  → interface Prestable, structs Libro y Revista
//   Farmacia    → interface Vendible, structs Medicamento y Suplemento
//   Gimnasio    → interface Registrable, structs Miembro y Visitante
//   Restaurante → interface Ordenable, structs Platillo y Bebida
// ============================================

// ============================================
// TODO 1: Declarar la interface
// ============================================
// Nombra la interface con un sustantivo o adjetivo
// que describa el comportamiento del dominio.
// Debe tener al menos un método.
//
// Ejemplo:
// type Elemento interface {
// 	Descripcion() string
// 	// Agrega otros métodos que tengan sentido en tu dominio
// }

// ============================================
// TODO 2: Primer struct con sus métodos
// ============================================
// type TipoA struct {
// 	nombre string
// 	// Agrega campos específicos de tu dominio
// }
//
// func (t TipoA) Descripcion() string {
// 	return t.nombre
// }
//
// // TODO: implementar otros métodos de la interface
//
// // TODO: implementar fmt.Stringer — String() string
// func (t TipoA) String() string {
// 	return fmt.Sprintf("[TipoA] %s", t.nombre)
// }

// ============================================
// TODO 3: Segundo struct con sus métodos
// ============================================
// type TipoB struct {
// 	nombre string
// 	// Agrega campos específicos de tu dominio
// }
//
// func (t TipoB) Descripcion() string {
// 	return t.nombre
// }
//
// // TODO: implementar otros métodos de la interface

// ============================================
// TODO 4: Función polimórfica
// ============================================
// Recibe la interface, no el tipo concreto
// func procesarElemento(e Elemento) {
// 	fmt.Println(e.Descripcion())
// }

// ============================================
// TODO 5: Verificación de compilación (fuera de main)
// ============================================
// var _ Elemento = TipoA{}
// var _ Elemento = TipoB{}

func main() {
	// ============================================
	// TODO 6: Crear instancias y usar la función polimórfica
	// ============================================
	// a := TipoA{nombre: "elemento A"}
	// b := TipoB{nombre: "elemento B"}
	//
	// procesarElemento(a)
	// procesarElemento(b)

	// ============================================
	// TODO 7: Slice de interface con valores de distintos tipos
	// ============================================
	// items := []Elemento{a, b}
	// for _, item := range items {
	// 	procesarElemento(item)
	// }

	// ============================================
	// TODO 8: fmt.Stringer — mostrar con fmt.Println
	// ============================================
	// Si TipoA implementa String() string:
	// fmt.Println(a) // usa String() automáticamente

	// Evitar error de compilación por variable no usada durante desarrollo
	_ = fmt.Sprintf
}
