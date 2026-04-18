# Rúbrica de Evaluación — Semana 05: Arrays y Slices

## 📊 Distribución de Evidencias

| Tipo | Porcentaje |
|------|-----------|
| 🧠 Conocimiento | 30% |
| 💪 Desempeño | 40% |
| 📦 Producto | 30% |

---

## 🧠 Conocimiento (30%) — Cuestionario teórico

### Preguntas de evaluación

1. ¿Qué diferencia fundamental existe entre un array y un slice en Go?
2. ¿Cuáles son los tres campos que componen el header de un slice?
3. ¿Qué ocurre con la capacidad cuando `append` supera el cap actual?
4. ¿Por qué modificar un sub-slice puede afectar al slice original?
5. ¿Qué función usas para copiar un slice sin compartir el backing array?
6. ¿Cuál es la diferencia entre `len(s)` y `cap(s)` en un slice?

### Criterios de evaluación (30 pts total)

| Pregunta | Puntos |
|----------|--------|
| Diferencia array vs slice | 5 pts |
| Campos del slice header | 5 pts |
| Comportamiento de append al superar cap | 5 pts |
| Aliasing en sub-slices | 5 pts |
| Uso correcto de copy | 5 pts |
| Diferencia len vs cap | 5 pts |

---

## 💪 Desempeño (40%) — Prácticas en clase

### Práctica 01 — Arrays (20 pts)

| Paso | Criterio | Puntos |
|------|----------|--------|
| Paso 1 | Declara array con literal y accede por índice correctamente | 5 pts |
| Paso 2 | Itera con `range` y acumula correctamente | 5 pts |
| Paso 3 | Pasa array a función y demuestra semántica de copia | 5 pts |
| Paso 4 | Compara arrays con `==` y los usa como claves de mapa | 5 pts |

### Práctica 02 — Slices (20 pts)

| Paso | Criterio | Puntos |
|------|----------|--------|
| Paso 1 | Crea slice con `make` y usa `append` correctamente | 4 pts |
| Paso 2 | Demuestra aliasing al hacer slicing y explica el porqué | 4 pts |
| Paso 3 | Usa `copy` para obtener slice independiente | 4 pts |
| Paso 4 | Implementa filtrado idiomático con nil slice | 4 pts |
| Paso 5 | Construye una matriz 2D con slice de slices | 4 pts |

---

## 📦 Producto (30%) — Proyecto integrador

### Criterios de evaluación (30 pts total)

| Criterio | Puntos |
|----------|--------|
| Slice inicializado con `make` y capacidad apropiada | 4 pts |
| Función `agregar` usando `append` correctamente | 4 pts |
| Función `buscar` recorriendo con `range` | 3 pts |
| Función `filtrar` retornando nuevo slice | 4 pts |
| Función `eliminar` preservando orden | 4 pts |
| Uso correcto de `copy` al clonar la colección | 4 pts |
| `go vet ./...` sin errores | 4 pts |
| Código adaptado coherentemente al dominio asignado | 3 pts |

---

## ✅ Criterios de Aprobación

- Mínimo **21/30** en conocimiento
- Mínimo **28/40** en desempeño
- Mínimo **21/30** en producto
- `go vet ./...` sin errores en el proyecto
- Proyecto funcional con el dominio asignado

---

## 📝 Notas de evaluación

- Se descuenta **1 punto** por cada `_` usado para ignorar errores donde no corresponde
- Se descuenta **2 puntos** si el proyecto presenta aliasing no intencional (bug de backing array)
- Se descuenta **2 puntos** si `append` se usa sin reasignar: `append(s, v)` sin capturar retorno
- **Bonus +2 pts**: implementar función `ordenar` con un algoritmo simple (sin `sort` de stdlib)
