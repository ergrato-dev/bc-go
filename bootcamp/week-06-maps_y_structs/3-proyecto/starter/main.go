package main

import "fmt"

// ============================================
// NOTA PARA EL APRENDIZ:
// Adapta este programa a tu dominio asignado.
// Ejemplos:
//   - Biblioteca:  Libro  (autor, ISBN, prestado)
//   - Farmacia:    Medicamento (dosis, stock, controlado)
//   - Gimnasio:    Miembro (plan, activo, fechaInscripcion)
//   - Escuela:     Estudiante (grado, promedio, activo)
// ============================================

// Pieza representa un elemento del dominio asignado.
// TODO: Renombra el tipo y adapta los campos a tu dominio.
type Pieza struct {
	ID           int
	Nombre       string
	Categoria    string
	EnExhibicion bool
	// TODO: Agrega campos específicos de tu dominio
}

// Resumen retorna una representación legible de la pieza.
// Es un método de valor: trabaja sobre una copia, no modifica el original.
// TODO: Adapta el formato a los campos de tu dominio.
func (p Pieza) Resumen() string {
	// 1. Determinar el estado (en exhibición / en bodega)
	// 2. Retornar un string formateado con los campos principales
	// Ejemplo: "[1] Vasija Maya — Cerámica (en exhibición)"
	return "" // TODO: implementar
}

// agregar añade p al slice (mantiene orden) y al map (búsqueda O(1)).
// Recibe puntero al slice para poder hacer append y que el cambio sea visible fuera.
// TODO: Renombra la función y el tipo si cambiaste el nombre del struct.
func agregar(p Pieza, registro *[]Pieza, indice map[int]Pieza) {
	// 1. Hacer append de p al slice apuntado por registro
	// 2. Asignar p en indice usando p.ID como clave
	// TODO: implementar
}

// buscar retorna la pieza con el id dado usando el map (no range sobre el slice).
// Retorna (Pieza, true) si existe, (Pieza{}, false) si no.
// TODO: Renombra la función y el tipo si es necesario.
func buscar(id int, indice map[int]Pieza) (Pieza, bool) {
	// Usar el patrón v, ok := indice[id]
	// TODO: implementar
	return Pieza{}, false
}

// listar imprime todas las piezas del slice en orden de inserción.
// TODO: Renombra la función y el tipo si es necesario.
func listar(registro []Pieza) {
	// Usar range sobre el slice para imprimir cada pieza con su Resumen()
	// TODO: implementar
}

func main() {
	// Inicializar las dos estructuras de datos
	registro := []Pieza{}
	indice := make(map[int]Pieza)
	// Suprimir error de "declared and not used" mientras el código está en TODOs.
	// Elimina esta línea cuando implementes las funciones.
	_ = indice

	// TODO: Crear al menos 4 instancias de tu struct usando literales con nombre de campo
	// Ejemplo (museo):
	// p1 := Pieza{ID: 1, Nombre: "Vasija Maya", Categoria: "Cerámica", EnExhibicion: true}
	// p2 := Pieza{ID: 2, Nombre: "Espada Medieval", Categoria: "Armamento", EnExhibicion: false}
	// p3 := Pieza{ID: 3, Nombre: "Papiro Egipcio", Categoria: "Documentos", EnExhibicion: true}
	// p4 := Pieza{ID: 4, Nombre: "Fósil Trilobite", Categoria: "Paleontología", EnExhibicion: true}

	// TODO: Agregar cada instancia con la función agregar
	// agregar(p1, &registro, indice)
	// agregar(p2, &registro, indice)
	// agregar(p3, &registro, indice)
	// agregar(p4, &registro, indice)

	// TODO: Listar todas las piezas
	fmt.Println("=== Catálogo completo ===")
	listar(registro)

	// TODO: Buscar una pieza existente
	fmt.Println("\n=== Búsqueda por ID ===")
	// if p, ok := buscar(1, indice); ok {
	// 	fmt.Println("Encontrada:", p.Resumen())
	// }

	// TODO: Buscar una pieza inexistente
	// if _, ok := buscar(99, indice); !ok {
	// 	fmt.Println("ID 99: no encontrado")
	// }
}
