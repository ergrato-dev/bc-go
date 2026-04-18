# Semana 06 — Maps y Structs

## 🎯 Objetivos de aprendizaje

Al finalizar esta semana serás capaz de:

- Declarar, inicializar y operar maps con claves y valores arbitrarios
- Detectar si una clave existe con el patrón `v, ok := m[k]`
- Definir structs con campos tipados y valores zero
- Usar structs como tipos de valor y comparar instancias con `==`
- Combinar maps y structs para modelar dominios del mundo real

## 📚 Requisitos previos

- Semana 05 — Arrays y Slices
- Tipos básicos, variables y control de flujo (Semanas 02-03)
- Funciones con múltiples retornos (Semana 04)

## 🗂️ Estructura de la semana

```
week-06-maps_y_structs/
├── README.md
├── rubrica-evaluacion.md
├── 0-assets/
│   ├── 01-maps-internals.svg
│   └── 02-structs-composicion.svg
├── 1-teoria/
│   ├── 01-maps.md
│   └── 02-structs-y-metodos-de-valor.md
├── 2-practicas/
│   ├── practica-01-maps/
│   │   ├── README.md
│   │   └── starter/main.go
│   └── practica-02-structs/
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

- [01 — Maps: tabla hash, acceso y patrones](1-teoria/01-maps.md)
- [02 — Structs: tipos de valor y métodos](1-teoria/02-structs-y-metodos-de-valor.md)

### Prácticas

- [Práctica 01 — Maps: CRUD y patrones idiomáticos](2-practicas/practica-01-maps/README.md)
- [Práctica 02 — Structs: definición, métodos de valor y composición](2-practicas/practica-02-structs/README.md)

### Proyecto

- [Proyecto semana 06 — Registro con maps y structs](3-proyecto/README.md)

## ⏱️ Distribución del tiempo (8 horas)

| Actividad | Tiempo |
|-----------|--------|
| Teoría 01 — Maps | 1 h |
| Teoría 02 — Structs | 1 h |
| Práctica 01 — Maps | 1.5 h |
| Práctica 02 — Structs | 1.5 h |
| Proyecto integrador | 3 h |

## 📌 Entregables

- **Único entregable obligatorio**: Proyecto semana 06
- Usa un map como índice de búsqueda rápida sobre structs
- Implementa al menos un método de valor sobre el struct principal
- `go vet ./...` sin errores antes de entregar

## 🔗 Navegación

← [Semana 05 — Arrays y Slices](../week-05-arrays_y_slices/README.md) | [Semana 07 — Punteros](../week-07-punteros/README.md) →
