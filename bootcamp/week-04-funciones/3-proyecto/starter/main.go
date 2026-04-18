package main
// Paquete main: proyecto semana 04 — procesador funcional
//   Qué     → aplicar múltiples retornos, variádicas, closures y HOF
//   Para qué → consolidar los patrones de funciones de Go en un programa real
//   Impacto  → adapta este starter a tu dominio asignado; los TODOs marcan los puntos de cambio


import (
	"fmt"
)

// ============================================
// TIPO PRINCIPAL
// TODO: Renombrar 'Obra' con el elemento de tu dominio
//       Ej: Libro, Medicamento, Miembro, Estudiante, Paciente
// ============================================

// Obra representa un elemento del dominio (ejemplo: museo).
// TODO: Cambiar nombre y campos según tu dominio asignado.
type Obra struct {
	ID        int
	Nombre    string
	Categoria string
	// TODO: Agregar campos relevantes de tu dominio
	// Ejemplo (Biblioteca): Autor string, Disponible bool
	// Ejemplo (Farmacia): Precio float64, Stock int
	// Ejemplo (Gimnasio): Nivel string, FechaIngreso string
}

// ============================================
// GENERADOR DE IDs — closure
// TODO: Renombrar para que refleje tu dominio
// Ej: nuevoGeneradorPaciente, nuevoGeneradorSocio
// ============================================

// nuevoGeneradorID retorna un closure que genera IDs secuenciales.
// Cada closure tiene su propio contador independiente en el heap.
func nuevoGeneradorID() func() int {
	siguiente := 0
	return func() int {
		siguiente++
		return siguiente
	}
}

// ============================================
// BÚSQUEDA — múltiples retornos (resultado, error)
// TODO: Adaptar para buscar elementos de tu dominio
// ============================================

// buscarPorID busca una Obra en el catálogo por su ID.
// Retorna la obra encontrada y nil, o zero value y error si no existe.
// TODO: Renombrar la función y el tipo de retorno según tu dominio
func buscarPorID(catalogo []Obra, id int) (Obra, error) {
	for _, obra := range catalogo {
		if obra.ID == id {
			return obra, nil
		}
	}
	// fmt.Errorf con %w permite envolver errores para errors.Is/As
	return Obra{}, fmt.Errorf("buscarPorID: obra con ID %d no encontrada", id)
}

// ============================================
// FUNCIÓN VARIÁDICA
// TODO: Adaptar para agregar atributos de tu dominio
// Ej: agregarEtiquetas(libro, etiquetas ...string)
//     agregarIngredientes(med, ingredientes ...string)
// ============================================

// agregarNotas agrega notas de curador a una obra.
// El parámetro variádico ...string es []string dentro de la función.
// TODO: Renombrar y cambiar semántica según tu dominio
func agregarNotas(obra Obra, notas ...string) string {
	if len(notas) == 0 {
		return obra.Nombre + " (sin notas)"
	}
	resultado := obra.Nombre + " — notas: "
	for i, nota := range notas {
		if i > 0 {
			resultado += ", "
		}
		resultado += nota
	}
	return resultado
}

// ============================================
// CLASIFICACIÓN — switch sobre categorías
// TODO: Reemplazar categorías con las de tu dominio
// ============================================

// clasificar determina el tipo de una obra según su categoría.
// TODO: Cambiar los casos del switch para tu dominio
//
//	Ej (Biblioteca): "ficcion", "no-ficcion", "referencia"
//	Ej (Farmacia): "analgesico", "antibiotico", "vitamina"
func clasificar(obra Obra) string {
	switch obra.Categoria {
	case "pintura", "escultura":
		// TODO: reemplazar con categorías de tu dominio
		return "Arte clásico"
	case "fotografía", "instalación":
		// TODO: reemplazar con categorías de tu dominio
		return "Arte contemporáneo"
	case "grabado", "dibujo":
		// TODO: reemplazar con categorías de tu dominio
		return "Arte gráfico"
	default:
		return "Sin clasificar"
	}
}

// ============================================
// HOF — filtrar por tipo
// TODO: Adaptar el predicado a tu dominio
// ============================================

// filtrar retorna los elementos que satisfacen el predicado.
// Es una función de orden superior: recibe una función como argumento.
func filtrar(obras []Obra, pred func(Obra) bool) []Obra {
	resultado := make([]Obra, 0, len(obras))
	for _, o := range obras {
		if pred(o) {
			resultado = append(resultado, o)
		}
	}
	return resultado
}

func main() {
	// TODO: Cambiar el mensaje para reflejar tu dominio
	defer fmt.Println("\n[Sistema cerrado correctamente]")

	fmt.Println("=== Proyecto: Procesador Funcional ===")
	fmt.Println()

	// TODO: Renombrar la variable y reemplazar los datos con tu dominio
	catalogo := []Obra{
		{ID: 1, Nombre: "La Gioconda", Categoria: "pintura"},
		{ID: 2, Nombre: "El Pensador", Categoria: "escultura"},
		{ID: 3, Nombre: "Migración", Categoria: "fotografía"},
		{ID: 4, Nombre: "Gravedad", Categoria: "instalación"},
		{ID: 5, Nombre: "Boceto #7", Categoria: "dibujo"},
		{ID: 6, Nombre: "Xilografía", Categoria: "grabado"},
	}

	// ============================================
	// Closure: generador de IDs
	// TODO: Usar para generar IDs de nuevos elementos al registrarlos
	// ============================================
	generarID := nuevoGeneradorID()

	// Simular registro de una nueva obra con ID generado
	nuevaObra := Obra{
		ID:        generarID(), // closure genera el siguiente ID
		Nombre:    "Nueva Adquisición",
		Categoria: "escultura",
	}
	// TODO: Reemplazar con lógica de registro de tu dominio
	fmt.Printf("Nueva obra registrada: ID=%d, Nombre=%s\n", nuevaObra.ID, nuevaObra.Nombre)
	fmt.Println()

	// ============================================
	// Múltiples retornos: buscar por ID
	// ============================================
	fmt.Println("--- Búsqueda por ID ---")
	obra, err := buscarPorID(catalogo, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// TODO: Adaptar el mensaje de resultado a tu dominio
		fmt.Printf("Encontrado: %s (categoría: %s)\n", obra.Nombre, obra.Categoria)
	}

	// Búsqueda que falla — el error se maneja con early return
	_, err = buscarPorID(catalogo, 99)
	if err != nil {
		fmt.Println("Error esperado:", err)
	}
	fmt.Println()

	// ============================================
	// Variádica: agregar notas
	// ============================================
	fmt.Println("--- Función variádica ---")
	// TODO: Reemplazar agregarNotas con la función variádica de tu dominio
	fmt.Println(agregarNotas(catalogo[0], "obra maestra", "restaurada 2021"))
	fmt.Println(agregarNotas(catalogo[1])) // cero argumentos es válido
	fmt.Println()

	// ============================================
	// Clasificación + reporte
	// ============================================
	fmt.Println("--- Clasificación del catálogo ---")
	conteos := make(map[string]int)
	for _, o := range catalogo {
		tipo := clasificar(o)
		conteos[tipo]++
		// TODO: Adaptar el formato de salida a tu dominio
		fmt.Printf("  %-20s → %s\n", o.Nombre, tipo)
	}
	fmt.Println()

	// Reporte de conteos
	fmt.Println("--- Resumen por tipo ---")
	// TODO: Iterar sobre conteos e imprimir cada tipo con su cantidad
	_ = conteos // reemplazar con: for tipo, cantidad := range conteos { ... }

	// ============================================
	// HOF: filtrar obras contemporáneas
	// ============================================
	fmt.Println("--- HOF: filtrar ---")
	// TODO: Cambiar el predicado para filtrar elementos de tu dominio
	contemporaneas := filtrar(catalogo, func(o Obra) bool {
		return o.Categoria == "fotografía" || o.Categoria == "instalación"
	})
	fmt.Printf("Arte contemporáneo (%d obras):\n", len(contemporaneas))
	for _, o := range contemporaneas {
		fmt.Printf("  - %s\n", o.Nombre)
	}
}
