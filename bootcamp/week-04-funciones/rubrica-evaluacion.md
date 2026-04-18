# Rúbrica de Evaluación — Semana 04: Funciones

> Distribución: Conocimiento 30% · Desempeño 40% · Producto 30%

---

## 🧠 Conocimiento (30 pts)

Evaluación teórica — 6 preguntas × 5 puntos

| # | Pregunta | Pts |
|---|----------|-----|
| 1 | ¿Cuál es la sintaxis para retornar múltiples valores en Go y cómo se capturan en el sitio de llamada? | 5 |
| 2 | ¿Qué son los retornos nombrados y en qué casos conviene usarlos? | 5 |
| 3 | ¿Cómo se declara una función variádica en Go? ¿Qué tipo tiene el parámetro dentro de la función? | 5 |
| 4 | ¿Qué significa que las funciones sean "valores de primera clase" en Go? Da un ejemplo concreto. | 5 |
| 5 | ¿Qué es un closure y qué variable(s) captura? Explica el entorno léxico. | 5 |
| 6 | ¿Cuál es la diferencia entre pasar `slice...` y pasar `slice` a una función variádica? | 5 |

**Total Conocimiento**: 30 pts

---

## 💪 Desempeño (40 pts)

### Práctica 01 — Múltiples retornos (20 pts)

| Paso | Criterio | Pts |
|------|----------|-----|
| 1 | Función con dos retornos (`(float64, error)`) correctamente definida y capturada | 4 |
| 2 | Retornos nombrados: función con `result float64` + `naked return` | 4 |
| 3 | Función variádica `sum(nums ...int)` con iteración correcta | 4 |
| 4 | Expansión de slice con `...` al llamar función variádica | 4 |
| 5 | Función como valor: asignación a variable, llamada desde variable | 4 |

### Práctica 02 — Closures (20 pts)

| Paso | Criterio | Pts |
|------|----------|-----|
| 1 | Closure contador: estado encapsulado incrementa correctamente | 5 |
| 2 | Closure generador de secuencia: retorna función que recuerda estado | 5 |
| 3 | HOF `applyToAll`: aplica función a cada elemento de un slice | 5 |
| 4 | Closure como middleware: función que envuelve otra con lógica adicional | 5 |

**Total Desempeño**: 40 pts

---

## 📦 Producto (30 pts)

### Proyecto — Procesador funcional del dominio asignado

| Criterio | Pts |
|----------|-----|
| Función con múltiples retornos (`(resultado, error)`) correctamente usada | 5 |
| Al menos una función variádica con sentido dentro del dominio | 4 |
| Función asignada a variable o pasada como argumento | 4 |
| Al menos un closure que encapsule estado del dominio | 5 |
| HOF que procese la colección principal del dominio | 4 |
| Manejo de error idiomático (early return, `fmt.Errorf` con `%w`) | 4 |
| `go vet ./...` sin errores | 2 |
| El dominio es coherente y reconocible (no se mezclan conceptos) | 2 |

**Total Producto**: 30 pts

---

## ⚠️ Penalizaciones

| Situación | Descuento |
|-----------|-----------|
| Uso de `_` para ignorar errores sin justificación | -3 pts por ocurrencia |
| Closure que no captura ninguna variable (no es closure real) | -3 pts |
| Función variádica llamada sin `...` cuando debería usarlo | -2 pts |
| Retornos nombrados con `return` explícito con valores (inconsistente) | -2 pts |
| Entrega sin adaptar al dominio asignado (usa ejemplos del starter) | -10 pts |

---

## ✅ Criterios de aprobación

- Mínimo **21/30** en Conocimiento
- Mínimo **28/40** en Desempeño
- Mínimo **21/30** en Producto
- `go vet ./...` sin errores en el proyecto
