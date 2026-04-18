package main

// Paquete main: proyecto semana 05 — gestor de colección funcional
//   Qué     → implementar operaciones CRUD sobre una colección usando slices
//   Para qué → aplicar make, append, copy, filtrado y eliminación idiomáticos
//   Impacto  → adapta este starter a tu dominio asignado; los TODOs marcan los puntos de cambio

import "fmt"

// ============================================
// TIPOS — adapta a tu dominio asignado
// ============================================

// Exhibicion representa un elemento de la colección del acuario.
//
// NOTA PARA EL APRENDIZ:
// Renombra este tipo y sus campos según tu dominio asignado.
// Ejemplos:
//   - Biblioteca:  Libro{ID, Titulo, Autor, Disponible bool}
//   - Farmacia:    Medicamento{ID, Nombre, Precio float64, Stock int}
//   - Gimnasio:    Miembro{ID, Nombre, Plan string, Activo bool}
//   - Restaurante: Plato{ID, Nombre, Categoria, Precio float64}
type Exhibicion struct {
	ID        int
	Nombre    string
	Categoria string
	// TODO: Agrega campos específicos de tu dominio
	// Ejemplo (Biblioteca): Autor string, Disponible bool
	// Ejemplo (Farmacia):   Precio float64, Stock int
}

// ============================================
// COLECCIÓN — slice con make
// ============================================

// nuevaColeccion crea una colección vacía con capacidad inicial.
// TODO: Ajusta la capacidad según el volumen esperado de tu dominio.
func nuevaColeccion(capacidad int) []Exhibicion {
	// TODO: Implementar usando make([]Exhibicion, 0, capacidad)
	return nil
}

// ============================================
// OPERACIONES CRUD
// ============================================

// agregar añade una nueva exhibicion a la colección.
// Retorna la colección actualizada — siempre captura el retorno.
// TODO: Implementar usando append.
func agregar(coleccion []Exhibicion, e Exhibicion) []Exhibicion {
	// TODO: Usar append y retornar el resultado
	return coleccion
}

// buscarPorID busca una exhibicion por su ID.
// Retorna (exhibicion, true) si la encuentra, o (zero value, false) si no.
// TODO: Implementar con range y comparación de IDs.
func buscarPorID(coleccion []Exhibicion, id int) (Exhibicion, bool) {
	// 1. Iterar con range sobre la colección
	// 2. Si e.ID == id, retornar (e, true)
	// 3. Si no se encontró, retornar (Exhibicion{}, false)
	return Exhibicion{}, false
}

// filtrarPorCategoria retorna un nuevo slice con las exhibiciones de la categoría dada.
// Usa nil slice + append para el filtrado idiomático.
// TODO: Implementar el filtro.
func filtrarPorCategoria(coleccion []Exhibicion, categoria string) []Exhibicion {
	// 1. Declarar var resultado []Exhibicion (nil slice)
	// 2. Iterar con range
	// 3. Si e.Categoria == categoria, append a resultado
	// 4. Retornar resultado
	return nil
}

// eliminar elimina la exhibicion con el ID dado, preservando el orden.
// Retorna la colección actualizada.
// TODO: Implementar preservando el orden de los elementos restantes.
func eliminar(coleccion []Exhibicion, id int) []Exhibicion {
	// 1. Buscar el índice del elemento con el ID dado
	// 2. Si no existe, retornar coleccion sin cambios
	// 3. Eliminar con append(coleccion[:i], coleccion[i+1:]...)
	// NOTA: esta operación modifica el backing array original — úsala con cuidado
	return coleccion
}

// clonar retorna una copia independiente de la colección usando copy.
// Modificar el clon NO afecta a la colección original.
// TODO: Implementar con make + copy.
func clonar(coleccion []Exhibicion) []Exhibicion {
	// 1. dst := make([]Exhibicion, len(coleccion))
	// 2. copy(dst, coleccion)
	// 3. Retornar dst
	return nil
}

// ============================================
// VISUALIZACIÓN
// ============================================

// imprimirColeccion muestra todos los elementos de la colección.
// TODO: Adapta el formato de impresión a tu dominio.
func imprimirColeccion(coleccion []Exhibicion) {
	if len(coleccion) == 0 {
		fmt.Println("  (colección vacía)")
		return
	}
	for _, e := range coleccion {
		// TODO: Ajusta el formato según los campos de tu dominio
		fmt.Printf("  [%d] %-20s | %-14s|\n", e.ID, e.Nombre, e.Categoria)
	}
}

// ============================================
// MAIN — programa principal
// ============================================

func main() {
	fmt.Println("=== Gestor de Colección ===")
	fmt.Println()

	// TODO: Ajusta el tamaño inicial según el volumen esperado de tu dominio
	coleccion := nuevaColeccion(10)

	// --- Agregar elementos ---
	fmt.Println("[1] Agregar elementos...")

	// TODO: Reemplaza estos datos con elementos de tu dominio asignado
	coleccion = agregar(coleccion, Exhibicion{ID: 1, Nombre: "Tiburón Toro", Categoria: "Peces"})
	coleccion = agregar(coleccion, Exhibicion{ID: 2, Nombre: "Pulpo Gigante", Categoria: "Invertebrados"})
	coleccion = agregar(coleccion, Exhibicion{ID: 3, Nombre: "Tortuga de Mar", Categoria: "Reptiles"})
	coleccion = agregar(coleccion, Exhibicion{ID: 4, Nombre: "Raya Manta", Categoria: "Peces"})

	fmt.Printf("Colección (%d elementos):\n", len(coleccion))
	imprimirColeccion(coleccion)
	fmt.Println()

	// --- Buscar por ID ---
	fmt.Println("[2] Buscar por ID...")

	// TODO: Busca un ID representativo de tu dominio
	if encontrado, ok := buscarPorID(coleccion, 2); ok {
		fmt.Printf("Encontrado: [%d] %s (%s)\n", encontrado.ID, encontrado.Nombre, encontrado.Categoria)
	} else {
		fmt.Println("No encontrado")
	}
	fmt.Println()

	// --- Filtrar por categoría ---
	fmt.Println("[3] Filtrar por categoría...")

	// TODO: Reemplaza "Peces" con una categoría de tu dominio
	categoria := "Peces"
	filtrados := filtrarPorCategoria(coleccion, categoria)
	fmt.Printf("Categoría '%s' (%d resultado/s):\n", categoria, len(filtrados))
	imprimirColeccion(filtrados)
	fmt.Println()

	// --- Eliminar un elemento ---
	fmt.Println("[4] Eliminar ID 2...")

	// TODO: Elimina un ID representativo de tu dominio
	coleccion = eliminar(coleccion, 2)
	fmt.Printf("Colección tras eliminar (%d elementos):\n", len(coleccion))
	imprimirColeccion(coleccion)
	fmt.Println()

	// --- Clonar colección ---
	fmt.Println("[5] Clonar colección...")

	copia := clonar(coleccion)

	// Modificar el clon NO debe afectar al original
	if len(copia) > 0 {
		copia[0].Nombre = "MODIFICADO"
	}

	fmt.Println("Original (sin cambios):")
	imprimirColeccion(coleccion)
	fmt.Println("Clon (modificado):")
	imprimirColeccion(copia)
}
