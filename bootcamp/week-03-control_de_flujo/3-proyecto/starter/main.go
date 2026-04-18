// Paquete main: proyecto semana 03 — procesador de registros
//   Qué     → clasificar y reportar elementos de un dominio usando control de flujo
//   Para qué → integrar for range, switch, defer e if con inicialización
//   Impacto  → adapta este programa a tu dominio asignado; los TODOs indican los puntos de cambio

package main

import "fmt"

// ============================================
// TIPO DE DATO
// Qué     → estructura que representa un elemento del dominio
// Para qué → agrupar los datos de cada registro en una sola unidad
// Impacto  → añade o cambia campos según las necesidades de tu dominio
// ============================================

// Item representa un elemento del catálogo del dominio asignado.
// TODO: Renombra Item y sus campos según tu dominio
// Ejemplos:
//   Biblioteca:  type Book struct { Title, Author string; Pages int }
//   Farmacia:    type Drug struct { Name string; Price float64; RequiresPrescription bool }
//   Gimnasio:    type Equipment struct { Name string; Zone string; Available bool }
type Item struct {
	Name     string
	Category string
	// TODO: Agrega campos relevantes para tu dominio
	// Ejemplo: Value float64, Available bool, Count int
}

// ============================================
// FUNCIÓN: classify
// Qué     → determina la clasificación de un item con switch
// Para qué → separar la lógica de clasificación del bucle principal
// Impacto  → añade más casos según las categorías de tu dominio
// ============================================

// classify retorna una etiqueta de clasificación para el item.
// TODO: Reemplaza los casos con las categorías de tu dominio (mínimo 3 + default)
// El switch puede ser sobre item.Category, item.Name, un campo numérico, etc.
func classify(item Item) string {
	switch item.Category {
	case "A":
		// TODO: Reemplaza "A" con el nombre de tu primera categoría
		// Ejemplo (Museo): case "Pintura": return "Arte visual"
		return "Categoría A"
	case "B":
		// TODO: Segunda categoría de tu dominio
		return "Categoría B"
	case "C":
		// TODO: Tercera categoría de tu dominio
		return "Categoría C"
	default:
		return "Sin categoría"
	}
}

func main() {
	// ============================================
	// DEFER DE CIERRE
	// Qué     → mensaje que se imprime cuando main retorna
	// Para qué → simular el cierre de un recurso o sesión
	// Impacto  → debe declararse al inicio de main, después de cualquier "apertura"
	// ============================================

	// TODO: Cambia el mensaje para que sea coherente con tu dominio
	// Ejemplos: "Cerrando catálogo de libros...", "Cerrando sistema de inventario..."
	defer fmt.Println("\nCerrando sistema de registros. ¡Hasta luego!")

	// ============================================
	// DATOS DEL DOMINIO
	// Qué     → slice con al menos 5 elementos del dominio
	// Para qué → iterar sobre ellos con for range
	// Impacto  → añade todos los campos necesarios para tu dominio
	// ============================================

	// TODO: Reemplaza estos items con elementos de tu dominio asignado
	// Usa al menos 5 elementos con categorías variadas para que el switch tenga utilidad
	items := []Item{
		{Name: "Elemento 1", Category: "A"},
		{Name: "Elemento 2", Category: "B"},
		{Name: "Elemento 3", Category: "A"},
		{Name: "Elemento 4", Category: "C"},
		{Name: "Elemento 5", Category: "B"},
		{Name: "Elemento 6", Category: "D"}, // caerá en default
	}

	// ============================================
	// PROCESAMIENTO CON for range + switch
	// Qué     → iterar sobre cada item y clasificarlo
	// Para qué → generar el reporte y acumular estadísticas
	// Impacto  → usar for range es más idiomático que indexado manual
	// ============================================

	// contadores por categoría para el resumen final
	// TODO: Ajusta los contadores según las categorías de tu dominio
	counts := make(map[string]int)

	fmt.Println("=== Reporte de Registros ===")
	fmt.Println()

	for _, item := range items {
		// TODO: Usa if con inicialización de variable al menos una vez aquí
		// Ejemplo: si necesitas validar un campo antes de clasificar
		// if label := classify(item); label != "Sin categoría" {
		//     fmt.Printf("  [✓] %s → %s\n", item.Name, label)
		// } else {
		//     fmt.Printf("  [?] %s → sin clasificar\n", item.Name)
		// }
		label := classify(item)
		counts[label]++
		fmt.Printf("  %-20s → %s\n", item.Name, label)
	}

	// ============================================
	// RESUMEN ESTADÍSTICO
	// Qué     → imprime el conteo total y por categoría
	// Para qué → dar una vista resumida del catálogo procesado
	// Impacto  → usa for range sobre el map de contadores
	// ============================================
	fmt.Println()
	fmt.Println("=== Resumen ===")
	fmt.Printf("Total de elementos: %d\n", len(items))

	// TODO: Itera sobre counts con for range para imprimir cada categoría
	// for categoria, cantidad := range counts {
	//     fmt.Printf("  %-20s: %d\n", categoria, cantidad)
	// }
	_ = counts // elimina esta línea cuando implementes el for range anterior
}
