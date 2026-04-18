# Semana 08 — Métodos e Interfaces

## Descripción

Esta semana introduces el sistema de tipos de Go en su dimensión más poderosa: los **métodos** como comportamiento adjunto a tipos, y las **interfaces** como contratos implícitos. Aprenderás el duck typing de Go —si un tipo tiene los métodos requeridos, satisface la interface automáticamente, sin declararlo explícitamente.

---

## 🎯 Objetivos de aprendizaje

Al finalizar esta semana serás capaz de:

- Definir métodos sobre cualquier tipo definido en el paquete
- Entender el **method set** de valores vs punteros
- Declarar interfaces con uno o más métodos
- Implementar interfaces de forma implícita (duck typing)
- Usar la interface `any` (`interface{}`) y sus implicaciones
- Pasar valores de interface a funciones polimórficas
- Usar la interface `fmt.Stringer` y `error` como ejemplos canónicos

---

## 📚 Requisitos previos

- Semana 06: structs y métodos básicos
- Semana 07: punteros y métodos de puntero
- Entender la diferencia entre receptor de valor y de puntero

---

## 🗂️ Estructura de la semana

```
week-08-metodos_e_interfaces/
├── README.md
├── rubrica-evaluacion.md
├── 0-assets/
│   ├── 01-method-sets.svg          # method set: T vs *T, interface satisfaction
│   └── 02-duck-typing.svg          # duck typing: tipos diferentes, misma interface
├── 1-teoria/
│   ├── 01-metodos-y-method-sets.md
│   └── 02-interfaces-implicitas.md
├── 2-practicas/
│   ├── practica-01-metodos/
│   └── practica-02-interfaces/
├── 3-proyecto/
├── 4-recursos/
└── 5-glosario/
```

---

## 📝 Contenidos

### Teoría

| Archivo | Tema |
|---------|------|
| [01-metodos-y-method-sets.md](1-teoria/01-metodos-y-method-sets.md) | Métodos sobre tipos, method sets, `fmt.Stringer` |
| [02-interfaces-implicitas.md](1-teoria/02-interfaces-implicitas.md) | Interfaces, duck typing, `any`, interface `error` |

### Prácticas

| Práctica | Descripción |
|----------|-------------|
| [practica-01-metodos](2-practicas/practica-01-metodos/) | Definir métodos, usar `fmt.Stringer`, method set |
| [practica-02-interfaces](2-practicas/practica-02-interfaces/) | Declarar interfaces, implementar implícitamente, polimorfismo |

### Proyecto

| Archivo | Descripción |
|---------|-------------|
| [3-proyecto/README.md](3-proyecto/README.md) | Instrucciones del proyecto integrador |
| [3-proyecto/starter/main.go](3-proyecto/starter/main.go) | Código inicial con TODOs |

---

## ⏱️ Distribución del tiempo (8 horas)

| Actividad | Tiempo |
|-----------|--------|
| Teoría 01: Métodos y method sets | 1h |
| Teoría 02: Interfaces implícitas | 1h |
| Práctica 01: Métodos | 1.5h |
| Práctica 02: Interfaces | 1.5h |
| Proyecto integrador | 2h |
| Revisión y autoevaluación | 1h |

---

## 📌 Entregables

- `3-proyecto/starter/main.go` con todos los TODOs resueltos
- Código compilando con `go run main.go` sin errores
- `go vet ./...` sin advertencias

---

## 🔗 Navegación

← [Semana 07 — Punteros](../week-07-punteros/README.md) | [Semana 09 — Composición y Embedding](../week-09-composicion_y_embedding/README.md) →
