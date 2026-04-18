# Semana 07 — Punteros

## 📋 Descripción

En Go, los punteros son la herramienta para compartir una variable entre varias partes del programa sin copiarla. Esta semana aprendemos a declarar punteros con `&` y `*`, a distinguir cuándo pasar por valor y cuándo por puntero, y a escribir métodos de puntero que modifican el receptor original.

## 🎯 Objetivos de aprendizaje

Al terminar esta semana serás capaz de:

- Declarar y usar punteros con `&` (dirección) y `*` (derreferencia)
- Entender la diferencia entre pasar por valor y pasar por puntero en funciones
- Escribir métodos de puntero `(t *Tipo)` que modifican el receptor original
- Usar `new(T)` para crear punteros a valores inicializados en zero value
- Reconocer cuándo Go desreferencia automáticamente un puntero (auto-deref)

## 📚 Requisitos previos

- Semana 06: Maps y Structs completada
- Conceptos de semántica de valor de structs

## 🗂️ Estructura de la semana

```
week-07-punteros/
├── README.md                          ← este archivo
├── rubrica-evaluacion.md
├── 0-assets/
│   ├── 01-stack-heap-punteros.svg     ← stack vs heap, & y *
│   └── 02-valor-vs-puntero.svg        ← paso por valor vs puntero
├── 1-teoria/
│   ├── 01-punteros-y-referencias.md   ← &, *, nil, new()
│   └── 02-paso-por-valor-vs-puntero.md ← cuándo usar cada uno
├── 2-practicas/
│   ├── practica-01-punteros/          ← &, *, nil, derreferencia
│   └── practica-02-new/               ← new(), métodos de puntero
├── 3-proyecto/
│   ├── README.md
│   └── starter/
│       └── main.go
├── 4-recursos/
│   ├── ebooks-free/
│   ├── videografia/
│   └── webgrafia/
└── 5-glosario/
    └── README.md
```

## 📝 Contenidos

### Teoría

| Archivo | Tema |
|---------|------|
| [01-punteros-y-referencias.md](1-teoria/01-punteros-y-referencias.md) | Operadores `&` y `*`, nil pointer, `new()`, auto-deref |
| [02-paso-por-valor-vs-puntero.md](1-teoria/02-paso-por-valor-vs-puntero.md) | Cuándo usar puntero, métodos de puntero, reglas de Go |

### Prácticas

| Práctica | Descripción |
|----------|-------------|
| [practica-01-punteros](2-practicas/practica-01-punteros/) | Declarar punteros, derreferir, evitar nil pointer panic |
| [practica-02-new](2-practicas/practica-02-new/) | `new()`, métodos de valor vs puntero, auto-deref |

### Proyecto

[Registro con métodos de puntero](3-proyecto/README.md) — Extiende el registro de la semana anterior agregando métodos de puntero para modificar structs.

## ⏱️ Distribución del tiempo (8 horas)

| Actividad | Tiempo |
|-----------|--------|
| Teoría 01: Punteros y referencias | 1 hora |
| Teoría 02: Paso por valor vs puntero | 1 hora |
| Práctica 01: Punteros básicos | 1.5 horas |
| Práctica 02: new() y métodos de puntero | 1.5 horas |
| Proyecto integrador | 2 horas |
| Revisión y correcciones | 1 hora |

## 📌 Entregables

- **Único entregable obligatorio**: `3-proyecto/starter/main.go` adaptado a tu dominio asignado
- Criterios detallados en [rubrica-evaluacion.md](rubrica-evaluacion.md)

## 🔗 Navegación

← [Semana 06 — Maps y Structs](../week-06-maps_y_structs/README.md) | [Semana 08 — Métodos e Interfaces](../week-08-metodos_e_interfaces/README.md) →
