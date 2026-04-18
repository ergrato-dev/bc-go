# Proyecto Semana 04 — Procesador Funcional

## Descripción

Implementa un **procesador funcional** para tu dominio asignado usando los patrones de funciones de la semana: múltiples retornos, funciones variádicas, closures y funciones de orden superior.

El ejemplo en el starter usa el dominio de un **museo** (colección de obras de arte). Tú debes adaptarlo completamente a tu dominio asignado.

---

## Dominio de ejemplo (NO usar en tu entrega)

El starter usa un museo con obras de arte clasificadas por categoría. Sirve solo como guía estructural.

---

## Ejemplos de adaptación por dominio

| Dominio | Elemento | Clasificar por | Variádica | Closure |
|---------|----------|----------------|-----------|---------|
| Biblioteca | Libro | género | agregarEtiquetas | contadorPréstamos |
| Farmacia | Medicamento | tipo | ingredientesActivos | controlStock |
| Gimnasio | Miembro | nivel | agregarRutinas | generadorID |
| Escuela | Estudiante | grado | agregarMaterias | promediador |
| Hospital | Paciente | especialidad | agregarDiagnósticos | generadorCodigo |

---

## Requisitos funcionales

Tu programa debe incluir:

1. **Struct del dominio** con al menos 3 campos relevantes
2. **Función con múltiples retornos** `(resultado, error)` — por ejemplo, buscar un elemento por ID
3. **Función variádica** que agregue atributos a un elemento del dominio
4. **Función de clasificación** (`switch` o `if`) que categorice elementos
5. **Closure** que encapsule estado del dominio (ej: contador, generador de ID, acumulador)
6. **HOF** (`filtrar` o `transformar`) que procese la colección principal

---

## Estructura esperada del programa

```
main()
├── defer — mensaje de cierre del sistema
├── colección de elementos (slice del struct del dominio)
├── generadorID := nuevoGeneradorID()   ← closure
├── for range colección
│   ├── buscarElemento() → (resultado, error)   ← múltiples retornos
│   └── clasificar(elemento)                     ← switch/if
├── HOF: filtrar o transformar sobre la colección
└── reporte final
```

---

## Entregable

Archivo `starter/main.go` modificado con:
- Struct renombrado al dominio asignado
- Todas las funciones implementadas (sin TODOs sin resolver)
- `go vet ./...` sin errores
- Comentarios en español explicando las decisiones

---

## Criterios de evaluación

Ver [rubrica-evaluacion.md](../../rubrica-evaluacion.md) — sección Producto (30 pts).
