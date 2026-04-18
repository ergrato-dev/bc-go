package main
// Paquete main: punto de entrada del proyecto semana 02
//   Qué     → catálogo de inventario con tipos, constantes e iota
//   Para qué → aplicar el sistema de tipos de Go a un dominio real asignado
//   Impacto  → adapta este archivo a tu dominio; los TODOs son los puntos de implementación


import "fmt"

// ============================================
// CONSTANTES DEL DOMINIO
// Qué     → valores fijos que identifican el sistema
// Para qué → centralizar la configuración en un solo lugar
// Impacto  → cambiar el nombre aquí lo actualiza en todo el programa
// ============================================

// TODO: Reemplaza "Acuario Municipal" con el nombre de tu dominio asignado
const systemName = "Acuario Municipal"

// TODO: Ajusta esta constante según tu dominio
// Ejemplos: maxCapacity (biblioteca), maxStock (farmacia), maxMembers (gimnasio)
const maxCapacity = 500

// TODO: Agrega al menos una constante numérica adicional relevante para tu dominio
// Ejemplo: const foundedYear = 1998

// ============================================
// TIPO PERSONALIZADO CON IOTA
// Qué     → enum de categorías del dominio usando iota
// Para qué → clasificar los elementos con seguridad de tipos
// Impacto  → el compilador rechaza usar un int arbitrario donde se espera Category
// ============================================

// Category representa el tipo de elemento en el catálogo.
// TODO: Renombra Category y sus valores según tu dominio
// Ejemplos:
//
//	Biblioteca:  type Genre int   — Fiction, NonFiction, Reference, Periodical
//	Farmacia:    type DrugType int — OTC, Prescription, Controlled
//	Gimnasio:    type Equipment int — Cardio, Strength, Flexibility, Recovery
type Category int

// TODO: Reemplaza los valores con las categorías de tu dominio (mínimo 3)
const (
	CategoryFish    Category = iota // 0
	CategoryMammal                  // 1
	CategoryReptile                 // 2
	CategoryInvert                  // 3
)

// ============================================
// FUNCIÓN: describeCategory
// Qué     → convierte un valor Category a su nombre legible
// Para qué → poder imprimir la categoría sin saber el número
// Impacto  → si agregas una categoría nueva, añade el caso aquí también
// ============================================

// TODO: Implementa esta función con un switch sobre el valor de cat
// Debe retornar el nombre de la categoría como string
// Si el valor es desconocido, retornar "desconocido"
func describeCategory(cat Category) string {
	// TODO: Usar switch cat { case CategoryFish: return "Pez" ... }
	return "desconocido"
}

func main() {
	// ============================================
	// 1. PRESENTACIÓN DEL SISTEMA
	// Usa las constantes para presentar el sistema
	// ============================================

	// TODO: Imprime el nombre del sistema y su capacidad máxima
	// Ejemplo: fmt.Printf("Sistema: %s | Capacidad: %d\n", systemName, maxCapacity)
	fmt.Printf("Sistema: %s\n", systemName)

	// ============================================
	// 2. DECLARAR UNA ENTRADA DEL CATÁLOGO
	// Declara variables para representar un elemento de tu dominio
	// ============================================

	// TODO: Renombra estas variables y sus valores según tu dominio
	// Usa := para la declaración dentro de main
	itemName := "Delfín mular"

	// TODO: Agrega más variables relevantes para tu dominio
	// Ejemplo (biblioteca): var available bool (zero value = false)
	// Ejemplo (farmacia):   price float64 = 12.50
	var itemCount int // zero value: 0

	// TODO: Asigna una categoría usando el tipo personalizado
	var itemCategory Category // zero value: 0 (primera categoría)

	// TODO: Realiza al menos una conversión de tipo explícita
	// Ejemplo: calcular ocupación como porcentaje
	//   occupancy := float64(itemCount) / float64(maxCapacity) * 100
	_ = float64(itemCount) // elimina este _ cuando implementes la conversión real

	// ============================================
	// 3. IMPRIMIR EL ELEMENTO CON fmt.Printf
	// Usa al menos %s (o %q), %d y %T (o %v)
	// ============================================

	// TODO: Imprime la información del elemento con los verbos correctos
	fmt.Printf("Nombre:    %s\n", itemName)
	fmt.Printf("Cantidad:  %d\n", itemCount)
	fmt.Printf("Categoría: %s (valor=%d, tipo=%T)\n",
		describeCategory(itemCategory), itemCategory, itemCategory)

	// ============================================
	// 4. INFORMACIÓN ADICIONAL DEL DOMINIO
	// ============================================

	// TODO: Imprime al menos una pieza de información adicional relevante
	// Ejemplos:
	//   fmt.Printf("Géneros disponibles: %d\n", int(CategoryInvert)+1)
	//   fmt.Printf("Año de fundación: %d\n", foundedYear)
	fmt.Println("---")
	fmt.Println("Programa adaptado a: [tu dominio asignado]")
}
