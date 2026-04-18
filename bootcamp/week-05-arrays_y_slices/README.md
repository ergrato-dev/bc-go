# Semana 05 — Arrays y Slices

## 🎯 Objetivos de aprendizaje

Al finalizar esta semana serás capaz de:

- Declarar y usar arrays de tamaño fijo como tipos de valor en Go
- Comprender la estructura interna de un slice (puntero, len, cap)
- Crear slices con literales, `make` y `append`
- Aplicar operaciones de slicing, `copy` y filtrado idiomático
- Evitar bugs de aliasing al trabajar con slices compartidos

## 📚 Requisitos previos

- Semana 04 — Funciones, múltiples retornos y closures
- Dominio de `for`, `range` y control de flujo (Semana 03)
- Tipos básicos y variables (Semana 02)

## 🗂️ Estructura de la semana

```
week-05-arrays_y_slices/
├── README.md
├── rubrica-evaluacion.md
├── 0-assets/
│   ├── 01-arrays-vs-slices.svg
│   └── 02-slice-operaciones.svg
├── 1-teoria/
│   ├── 01-arrays.md
│   └── 02-slices-internals.md
├── 2-practicas/
│   ├── practica-01-arrays/
│   │   ├── README.md
│   │   └── starter/main.go
│   └── practica-02-slices/
│       ├── README.md
│       └── starter/main.go
├── 3-proyecto/
│   ├── README.md
│   └── starter/main.go
├── 4-recursos/
│   ├── ebooks-free/README.md
│   ├── videografia/README.md
│   └── webgrafia/README.md
└── 5-glosario/
    └── README.md
```

## 📝 Contenidos

### Teoría

- [01 — Arrays: tipos de valor y tamaño fijo](1-teoria/01-arrays.md)
- [02 — Slices: internals, make y append](1-teoria/02-slices-internals.md)

### Prácticas

- [Práctica 01 — Arrays y range](2-practicas/practica-01-arrays/README.md)
- [Práctica 02 — Slices, make y append](2-practicas/practica-02-slices/README.md)

### Proyecto

- [Proyecto semana 05 — Gestor de colección funcional](3-proyecto/README.md)

## ⏱️ Distribución del tiempo (8 horas)

| Actividad | Tiempo |
|-----------|--------|
| Teoría 01 — Arrays | 1 h |
| Teoría 02 — Slices internals | 1 h |
| Práctica 01 — Arrays y range | 1.5 h |
| Práctica 02 — Slices, make y append | 1.5 h |
| Proyecto integrador | 3 h |

## 📌 Entregables

- **Único entregable obligatorio**: Proyecto semana 05
- Implementa el gestor de colección con tu dominio asignado
- Usa slices con `make`, `append`, `copy` y filtrado idiomático
- `go vet ./...` sin errores antes de entregar

## 🔗 Navegación

← [Semana 04 — Funciones](../week-04-funciones/README.md) | [Semana 06 — Maps y Structs](../week-06-maps_y_structs/README.md) →
