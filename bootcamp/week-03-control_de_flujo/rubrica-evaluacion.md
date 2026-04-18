# Rúbrica de Evaluación — Semana 03: Control de Flujo

## Distribución de puntos

| Tipo de evidencia | Porcentaje | Puntos |
|-------------------|-----------|--------|
| Conocimiento 🧠 | 30% | 30 pts |
| Desempeño 💪 | 40% | 40 pts |
| Producto 📦 | 30% | 30 pts |
| **Total** | **100%** | **100 pts** |

> **Mínimo para aprobar**: 70 puntos totales y al menos 70% en cada tipo de evidencia.

---

## 🧠 Conocimiento (30 pts)

Evaluación oral o escrita al finalizar la semana. 6 preguntas × 5 puntos c/u.

| # | Pregunta | Pts |
|---|----------|-----|
| 1 | ¿Cuál es la diferencia entre `if cond {}` y `if v := expr; cond {}`? ¿Cuándo conviene usar la segunda forma? | 5 |
| 2 | Go tiene un único constructor de bucle. Describe los 4 patrones que puede adoptar `for` con código de ejemplo de cada uno. | 5 |
| 3 | ¿Por qué en Go el `switch` no necesita `break`? ¿Cómo se logra el comportamiento opuesto (fall-through)? | 5 |
| 4 | ¿En qué orden se ejecutan múltiples llamadas a `defer` dentro de una función? Da un ejemplo con salida esperada. | 5 |
| 5 | ¿Cuándo se evalúan los argumentos de una función llamada con `defer`? Demuéstralo con código. | 5 |
| 6 | Explica el patrón `defer f.Close()` y por qué es preferible usarlo inmediatamente después de abrir un recurso. | 5 |

---

## 💪 Desempeño (40 pts)

### Práctica 01 — `for` (20 pts)

| Criterio | Pts |
|----------|-----|
| Paso 1: `for` clásico con índice, condición y post (produce salida correcta) | 4 |
| Paso 2: `for` como while (condición booleana, sin inicialización ni post) | 4 |
| Paso 3: `for range` sobre slice de strings (índice y valor correctos) | 4 |
| Paso 4: `for range` sobre string con runes Unicode (imprime cada rune y su valor) | 4 |
| Paso 5: `break` y `continue` (salir al encontrar condición, saltar elementos) | 4 |

### Práctica 02 — `switch` y `defer` (20 pts)

| Criterio | Pts |
|----------|-----|
| Paso 1: `switch` con expresión — múltiples valores por caso (`case "a", "e", "i":`) | 5 |
| Paso 2: `switch` sin expresión — funciona como if-else encadenado | 5 |
| Paso 3: `defer` simple — orden LIFO verificable en la salida | 5 |
| Paso 4: `defer` con argumento evaluado en el momento de la llamada | 5 |

---

## 📦 Producto (30 pts)

El proyecto integrador debe procesar una lista de elementos del dominio asignado usando control de flujo.

| Criterio | Pts |
|----------|-----|
| El programa compila y ejecuta sin errores (`go run .`) | 5 |
| `go vet .` sin advertencias | 3 |
| Usa `for range` para iterar sobre los elementos del dominio | 4 |
| Usa `switch` para clasificar o categorizar elementos (mínimo 3 casos) | 4 |
| Usa `defer` al menos una vez de forma idiomática | 4 |
| Usa `if` con inicialización de variable al menos una vez | 4 |
| La salida describe correctamente el dominio asignado | 3 |
| Código limpio: indentación correcta, sin dead code, comentarios educativos | 3 |

---

## ⚠️ Penalizaciones

| Infracción | Penalización |
|------------|-------------|
| Uso de `fallthrough` sin justificación documentada | -3 pts |
| `defer` dentro de un bucle sin explicación (puede causar resource leak) | -5 pts |
| Variable de loop capturada incorrectamente en closure | -5 pts |
| Código copiado de otro aprendiz (mismo dominio o lógica idéntica) | -100% producto |
| Entrega tardía (por cada día hábil) | -5 pts |
