package main

import "fmt"

// ============================================
// INSTRUCCIONES PARA EL APRENDIZ
//
// Adapta este programa a tu dominio asignado.
// Ejemplos de adaptación:
//   Biblioteca  → Item = Libro,  campo1=titulo, campo2=autor, campo3=disponible (bool)
//   Farmacia    → Item = Medicamento, campo1=nombre, campo2=proveedor, campo3=activo (bool)
//   Gimnasio    → Item = Miembro, campo1=nombre, campo2=plan, campo3=activo (bool)
//   Restaurante → Item = Platillo, campo1=nombre, campo2=categoria, campo3=disponible (bool)
//
// Reglas:
//   1. Renombra Item y sus campos al dominio
//   2. Completa todos los TODO
//   3. go run main.go debe ejecutar sin errores
//   4. go vet ./...  debe pasar sin advertencias
// ============================================

// Item representa un elemento de tu dominio asignado.
// TODO: Renombra este struct y sus campos según tu dominio.
// Ejemplo (Biblioteca): type Libro struct { titulo, autor string; disponible bool }
type Item struct {
	nombre    string
	categoria string
	activo    bool
}

// descripcion retorna un resumen del item (receptor de VALOR — solo lectura).
// TODO: Adapta el mensaje al vocabulario de tu dominio.
// Ejemplo (Biblioteca): "El Señor de los Anillos de Tolkien — disponible"
func (i Item) descripcion() string {
	// TODO: Construir y retornar una cadena descriptiva del item.
	// Usa fmt.Sprintf y los campos del struct.
	// Ejemplo: return fmt.Sprintf("%s (%s) — activo: %v", i.nombre, i.categoria, i.activo)
	return ""
}

// activar cambia el estado del item a activo (receptor de PUNTERO — modifica).
// TODO: Implementar. Debe cambiar el campo booleano de estado a true.
func (i *Item) activar() {
	// TODO: i.activo = true (o el nombre del campo de estado de tu dominio)
}

// desactivar cambia el estado del item a inactivo (receptor de PUNTERO — modifica).
// TODO: Implementar. Debe cambiar el campo booleano de estado a false.
func (i *Item) desactivar() {
	// TODO: i.activo = false
}

// actualizarCategoria cambia la categoría del item (función, no método).
// Recibe un puntero para poder modificar el struct original.
// TODO: Implementar. Debe asignar el parámetro nuevaCategoria al campo correspondiente.
func actualizarCategoria(i *Item, nuevaCategoria string) {
	// TODO: i.categoria = nuevaCategoria
}

func main() {
	// ============================================
	// 1. Crear un item usando &Item{} o new(Item)
	// ============================================
	// TODO: Crea un item con valores iniciales usando &Item{...} o new(Item).
	// Ejemplo:
	// item := &Item{nombre: "Elemento de prueba", categoria: "General", activo: false}
	// fmt.Println("Creado:", item.descripcion())

	// ============================================
	// 2. Usar métodos de puntero para modificar
	// ============================================
	// TODO: Llama a item.activar() y muestra el resultado con item.descripcion().
	// TODO: Llama a item.desactivar() y muestra el resultado.

	// ============================================
	// 3. Modificar con función que recibe puntero
	// ============================================
	// TODO: Llama a actualizarCategoria(item, "nueva categoría") y muestra el resultado.

	// ============================================
	// 4. Nil check antes de usar un puntero
	// ============================================
	// TODO: Declara un *Item nil y verifica antes de usarlo.
	// var vacio *Item
	// if vacio != nil {
	//     fmt.Println(vacio.descripcion())
	// } else {
	//     fmt.Println("sin item asignado")
	// }

	// Evitar error de compilación por variable no usada durante desarrollo
	_ = fmt.Sprintf
}
