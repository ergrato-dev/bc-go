# Semana 02 — Variables, Tipos y Constantes

## 📋 Descripción

Segunda semana del bootcamp. Exploramos el sistema de tipos de Go: todos los tipos básicos, sus zero values, conversión explícita entre tipos y el poderoso mecanismo de `const` con `iota` para crear enumeraciones idiomáticas.

## 🎯 Objetivos de Aprendizaje

Al finalizar esta semana, serás capaz de:

- ✅ Declarar variables con `var` y `:=` en los contextos correctos
- ✅ Conocer todos los tipos básicos de Go: `bool`, `int`, `float64`, `string`, `byte`, `rune`
- ✅ Explicar el zero value de cada tipo y por qué Go los garantiza
- ✅ Convertir entre tipos de manera explícita (sin conversiones implícitas)
- ✅ Declarar constantes con `const` y entender la diferencia entre tipadas y no tipadas
- ✅ Usar `iota` para crear enumeraciones legibles y mantenibles
- ✅ Imprimir variables con los verbos de formato correctos (`%T`, `%v`, `%d`, `%f`, `%q`)

## 📚 Requisitos Previos

- Semana 01 completada (`go run`, `go mod init`, `package main` dominados)
- VS Code con la extensión oficial de Go instalada
- Familiaridad con el concepto de tipo de dato

## 🗂️ Estructura de la Semana

```
week-02-variables_tipos_y_constantes/
├── README.md                              # Este archivo
├── rubrica-evaluacion.md                  # Criterios de evaluación
├── 0-assets/
│   ├── 01-tipos-basicos-zero-values.svg   # Diagrama de tipos y zero values
│   └── 02-constantes-iota.svg             # Diagrama de const e iota
├── 1-teoria/
│   ├── 01-tipos-basicos-y-zero-values.md  # Teoría: tipos, var, :=, conversión
│   └── 02-constantes-e-iota.md            # Teoría: const, iota, enums
├── 2-practicas/
│   ├── practica-01-tipos/                 # Explorar todos los tipos básicos
│   │   ├── README.md
│   │   └── starter/
│   │       └── main.go
│   └── practica-02-constantes-iota/       # const, iota y patrones
│       ├── README.md
│       └── starter/
│           └── main.go
├── 3-proyecto/
│   ├── README.md                          # Proyecto: catálogo con tipos y constantes
│   └── starter/
│       └── main.go
├── 4-recursos/
│   ├── ebooks-free/
│   │   └── README.md
│   ├── videografia/
│   │   └── README.md
│   └── webgrafia/
│       └── README.md
└── 5-glosario/
    └── README.md
```

## 📝 Contenidos

### Teoría

- [01 — Tipos básicos y zero values: var, :=, conversión explícita](1-teoria/01-tipos-basicos-y-zero-values.md)
- [02 — Constantes e iota: const, enumeraciones idiomáticas](1-teoria/02-constantes-e-iota.md)

### Prácticas

- [Práctica 01 — Tipos básicos: explorar el sistema de tipos de Go](2-practicas/practica-01-tipos/README.md)
- [Práctica 02 — Constantes e iota: enumeraciones con const](2-practicas/practica-02-constantes-iota/README.md)

### Proyecto

- [Proyecto Semana 02 — Catálogo con tipos y constantes](3-proyecto/README.md)

## ⏱️ Distribución del Tiempo (8 horas)

| Actividad | Tiempo |
|-----------|--------|
| Teoría 01: Tipos básicos y zero values | 1.5 h |
| Teoría 02: Constantes e iota | 1 h |
| Práctica 01: Tipos básicos | 1.5 h |
| Práctica 02: Constantes e iota | 1.5 h |
| Proyecto: Catálogo | 2.5 h |

## 📌 Entregables

> **El único entregable obligatorio** para aprobar la semana es el **Proyecto**.

- [ ] Proyecto semana 02 completo y funcional (`go vet .` sin errores)

## 🔗 Navegación

| Anterior | Siguiente |
|----------|-----------|
| [← Semana 01 — Introducción a Go](../week-01-introduccion_go_y_herramientas/README.md) | [Semana 03 — Control de Flujo →](../week-03-control_de_flujo/README.md) |
