# Semana 03 — Control de Flujo

> **Prerequisitos**: Semana 02 — Variables, Tipos y Constantes

## 🎯 Objetivos de aprendizaje

Al finalizar esta semana serás capaz de:

- Usar `if` con inicialización de variable en la condición (`if v, err := ...; err != nil`)
- Implementar todos los patrones del único bucle de Go: `for` (clásico, while, infinito, range)
- Aplicar `switch` sin `break` explícito y con múltiples casos por rama
- Entender y usar `defer` correctamente: stack LIFO, evaluación de argumentos y su uso con recursos
- Elegir la estructura de control adecuada según el caso de uso
- Evitar los errores comunes: bucles infinitos, variables de `defer` capturadas por referencia

## 📚 Requisitos previos

- Variables con `var` y `:=` (Semana 02)
- Tipos básicos: `int`, `bool`, `string` (Semana 02)
- Constantes e `iota` (Semana 02)

## 🗂️ Estructura de la semana

```
week-03-control_de_flujo/
├── README.md
├── rubrica-evaluacion.md
├── 0-assets/
│   ├── 01-if-for-flow.svg
│   └── 02-switch-defer-flow.svg
├── 1-teoria/
│   ├── 01-if-y-for.md
│   └── 02-switch-y-defer.md
├── 2-practicas/
│   ├── practica-01-for/
│   │   ├── README.md
│   │   └── starter/main.go
│   └── practica-02-switch-defer/
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

| Archivo | Tema |
|---------|------|
| [01-if-y-for.md](1-teoria/01-if-y-for.md) | `if` con inicialización, todos los patrones de `for` |
| [02-switch-y-defer.md](1-teoria/02-switch-y-defer.md) | `switch` (expresión, tipo, sin condición), `defer` LIFO |

### Prácticas

| Práctica | Tema |
|----------|------|
| [practica-01-for](2-practicas/practica-01-for/) | Patrones de `for`: clásico, while, range, infinito + break/continue |
| [practica-02-switch-defer](2-practicas/practica-02-switch-defer/) | `switch` con múltiples casos, `defer` con recursos |

### Proyecto

| Entregable | Descripción |
|------------|-------------|
| [3-proyecto/](3-proyecto/) | Procesador de registros con clasificación y estadísticas |

## ⏱️ Distribución del tiempo (8 horas)

| Actividad | Tiempo |
|-----------|--------|
| Teoría: `if` y `for` | 1 h |
| Teoría: `switch` y `defer` | 1 h |
| Práctica 01: `for` | 1.5 h |
| Práctica 02: `switch` y `defer` | 1.5 h |
| Proyecto integrador | 2.5 h |
| Revisión y `go vet` | 30 min |

## 📌 Entregables

- **Único entregable obligatorio**: `3-proyecto/starter/main.go` adaptado a tu dominio asignado
- Ejecutar `go vet .` antes de entregar; cero warnings

## 🔗 Navegación

← [Semana 02 — Variables, Tipos y Constantes](../week-02-variables_tipos_y_constantes/README.md)  
→ [Semana 04 — Funciones](../week-04-funciones/README.md)
