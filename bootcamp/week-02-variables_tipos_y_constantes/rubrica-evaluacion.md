# Rúbrica de Evaluación — Semana 02
## Variables, Tipos y Constantes

---

## 📊 Distribución de evidencias

| Tipo de evidencia | Peso | Mínimo para aprobar |
|-------------------|------|---------------------|
| 🧠 Conocimiento   | 30 % | 21 / 30 pts         |
| 💪 Desempeño      | 40 % | 28 / 40 pts         |
| 📦 Producto       | 30 % | 21 / 30 pts         |
| **Total**         | **100 %** | **70 / 100 pts** |

---

## 🧠 Conocimiento (30 pts)

Evaluado mediante cuestionario oral o escrito al inicio de la siguiente sesión.

| Pregunta | Pts |
|----------|-----|
| ¿Cuál es la diferencia entre `var x int` y `x := 0`? ¿Cuándo usarías cada uno? | 5 |
| ¿Qué es un zero value en Go? Da el zero value de `int`, `string`, `bool` y un puntero. | 5 |
| ¿Por qué Go no permite conversiones implícitas de tipos? ¿Cómo se convierten explícitamente? | 5 |
| Explica la diferencia entre una constante tipada y una no tipada. ¿Por qué importa? | 5 |
| ¿Qué es `iota` y qué valor toma en cada constante de un bloque `const`? | 5 |
| ¿Cuál es la diferencia entre `byte` y `rune` en Go? ¿Qué tipo de caracteres representa cada uno? | 5 |

---

## 💪 Desempeño (40 pts)

Evaluado durante la sesión práctica en clase.

### Práctica 01 — Tipos básicos (20 pts)

| Criterio | Pts |
|----------|-----|
| Declara correctamente variables de al menos 5 tipos distintos (Paso 1) | 4 |
| Imprime el zero value de cada tipo con `fmt.Printf` y el verbo `%v` (Paso 2) | 4 |
| Realiza al menos una conversión de tipo explícita sin errores de compilación (Paso 3) | 6 |
| Usa `%T` para imprimir el tipo dinámico de una variable (Paso 4) | 4 |
| Manipula una cadena como slice de `rune` para acceder a caracteres Unicode (Paso 5) | 2 |

### Práctica 02 — Constantes e iota (20 pts)

| Criterio | Pts |
|----------|-----|
| Declara constantes tipadas y no tipadas correctamente (Paso 1) | 4 |
| Usa un bloque `const` para agrupar constantes relacionadas (Paso 2) | 4 |
| Implementa una enumeración con `iota` que empieza desde 1 (Paso 3) | 6 |
| Usa `iota` con desplazamiento de bits (`1 << iota`) para flags de permisos (Paso 4) | 6 |

---

## 📦 Producto (30 pts)

Entregable: **proyecto semanal** adaptado al dominio asignado.

### Criterios de evaluación del proyecto

| Criterio | Pts |
|----------|-----|
| El programa compila con `go run .` sin errores | 4 |
| `go vet .` pasa sin advertencias ni errores | 4 |
| Define un tipo personalizado con `iota` para al menos una categoría del dominio | 6 |
| Usa `const` para al menos dos valores fijos del dominio (nombre, capacidad, etc.) | 4 |
| Declara variables con `var` y `:=` en los contextos adecuados | 4 |
| Realiza al menos una conversión de tipo explícita necesaria para el dominio | 4 |
| Imprime información con `fmt.Printf` usando verbos de formato correctos | 2 |
| Comentarios en español explicando las decisiones de implementación | 2 |

### Penalizaciones

| Infracción | Penalización |
|------------|--------------|
| `go vet .` produce errores no corregidos | -5 pts |
| El dominio usado es de la lista restringida (biblioteca, farmacia, etc.) | -10 pts |
| Código copiado de otro aprendiz (dominios distintos con implementación idéntica) | -30 pts (nota 0) |
| Entrega después del plazo sin justificación | -10 pts |

---

## 🔗 Navegación

| Anterior | Siguiente |
|----------|-----------|
| [← Semana 01 — Introducción a Go](../week-01-introduccion_go_y_herramientas/rubrica-evaluacion.md) | [Semana 03 — Control de Flujo →](../week-03-control_de_flujo/rubrica-evaluacion.md) |
