---
mode: agent
description: Genera el proyecto semanal integrador de una semana del bootcamp de Go.
---

Crea el proyecto semanal integrador para el bootcamp bc-go.

## Datos del proyecto

- **Semana**: [COMPLETAR — ej. 05]
- **Tema**: [COMPLETAR — ej. Arrays y Slices]
- **Ruta destino**: `bootcamp/week-<NN>-<slug>/3-proyecto/`

## Archivos a generar

```
3-proyecto/
├── README.md
└── starter/
    └── main.go
```

> No hay `solution/` en proyectos — cada aprendiz implementa sobre su dominio único.

## Política de dominios (CRÍTICO)

El proyecto usa ejemplos **genéricos** adaptables a cualquier dominio.
**NO usar** como dominio de ejemplo: biblioteca, farmacia, gimnasio, escuela,
tienda de mascotas, restaurante, banco, agencia de taxis, hospital, cine,
hotel, agencia de viajes, concesionario de autos, tienda de ropa, taller mecánico.

Usar dominios alternativos para los ejemplos: museo, planetario, acuario,
zoológico, galería de arte, estación espacial, observatorio, etc.

## Formato del README.md del proyecto

```markdown
# Proyecto Semana XX — Título

## Contexto

Breve descripción del caso de negocio genérico (2-3 líneas).

## Tu dominio asignado

Adapta el código y los tipos al dominio que te asignó el instructor.
Ejemplos de adaptación:

- Biblioteca → Book, Member, Loan
- Farmacia → Medicine, Sale, Inventory
- Gimnasio → Member, Routine, Attendance

## Requisitos

1. **Requisito 1**: descripción concreta y verificable
2. **Requisito 2**: ...
3. **Requisito 3**: ...
   (mínimo 4, máximo 6 requisitos)

## Entregables

- [ ] `main.go` adaptado a tu dominio con tipos propios
- [ ] Programa ejecutable con `go run .` sin errores
- [ ] `go vet .` sin advertencias
- [ ] Cada función con comentario explicando qué hace

## Criterios de evaluación

| Criterio                                   | Puntaje     |
|--------------------------------------------|-------------|
| Programa compila y ejecuta correctamente   | 20 pts      |
| Uso correcto del concepto de la semana     | 40 pts      |
| Adaptación coherente al dominio asignado   | 30 pts      |
| Comentarios en español y código limpio     | 10 pts      |
| **Total**                                  | **100 pts** |
```

## Formato del starter/main.go

```go
package main

import "fmt"

// Item representa un elemento del dominio asignado.
// NOTA PARA EL APRENDIZ:
// Adapta este tipo a tu dominio asignado.
// Ejemplos:
//   - Biblioteca: Book con Author, Available bool
//   - Farmacia: Medicine con Price float64, Stock int
//   - Gimnasio: Member con Plan, Active bool
type Item struct {
	ID   int
	Name string
	// TODO: Agregar campos específicos de tu dominio
}

// processItems aplica el concepto de la semana sobre la colección de items.
// TODO: Implementar usando lo aprendido en la semana
func processItems(items []Item) []Item {
	// 1. [Paso 1 del concepto de la semana]
	// 2. [Paso 2 del concepto de la semana]
	// 3. Retornar resultado
	return nil
}

func main() {
	// TODO: Crear al menos 5 items de tu dominio
	items := []Item{}

	// TODO: Llamar a processItems y mostrar resultados
	result := processItems(items)
	fmt.Println(result)
}
```

## Reglas

- Mínimo 4 `// TODO:` de implementación (uno por requisito)
- Cada TODO debe ser específico y verificable, no genérico
- El starter debe compilar con `go run .` mostrando zero-values o output vacío
- `go vet .` debe pasar en el starter (sin errores)
- Los comentarios del starter deben orientar al alumno sin revelar la implementación
- Semanas 18-20 (Web & APIs): incluir bloque "Cómo ejecutar" con Docker en README
