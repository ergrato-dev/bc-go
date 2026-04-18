# Rúbrica de Evaluación — Semana 06: Maps y Structs

## 📊 Distribución de Evidencias

| Tipo | Porcentaje |
|------|-----------|
| 🧠 Conocimiento | 30% |
| 💪 Desempeño | 40% |
| 📦 Producto | 30% |

---

## 🧠 Conocimiento (30%) — Cuestionario teórico

### Preguntas de evaluación

1. ¿Qué retorna `m[k]` cuando la clave `k` no existe en el map?
2. ¿Cuál es el patrón idiomático para verificar si una clave existe en un map?
3. ¿Por qué los mapas no son seguros para uso concurrente en Go?
4. ¿En qué se diferencia un struct de un mapa a nivel de tipos?
5. ¿Qué significa que los structs son tipos de valor en Go?
6. ¿Cuándo son comparables dos structs con `==`?

### Criterios de evaluación (30 pts total)

| Pregunta | Puntos |
|----------|--------|
| Zero value al acceder a clave inexistente | 5 pts |
| Patrón `v, ok := m[k]` | 5 pts |
| Seguridad concurrente de maps | 5 pts |
| Struct vs map: tipos estáticos vs dinámicos | 5 pts |
| Semántica de valor de structs | 5 pts |
| Comparabilidad de structs con == | 5 pts |

---

## 💪 Desempeño (40%) — Prácticas en clase

### Práctica 01 — Maps (20 pts)

| Paso | Criterio | Puntos |
|------|----------|--------|
| Paso 1 | Declara map con `make` y asigna valores correctamente | 5 pts |
| Paso 2 | Usa el patrón `v, ok` para distinguir clave inexistente de zero value | 5 pts |
| Paso 3 | Itera el map con `range` y construye frecuencias | 5 pts |
| Paso 4 | Elimina entradas con `delete` y usa map como conjunto (set) | 5 pts |

### Práctica 02 — Structs (20 pts)

| Paso | Criterio | Puntos |
|------|----------|--------|
| Paso 1 | Define struct con campos tipados y usa zero value correctamente | 5 pts |
| Paso 2 | Declara e inicializa struct con literal con nombre de campos | 5 pts |
| Paso 3 | Implementa método de valor con receptor `(t Tipo)` | 5 pts |
| Paso 4 | Compara dos structs con `==` y demuestra semántica de valor | 5 pts |

---

## 📦 Producto (30%) — Proyecto integrador

### Criterios de evaluación (30 pts total)

| Criterio | Puntos |
|----------|--------|
| Struct principal con al menos 4 campos tipados | 4 pts |
| Map usado como índice `map[id]Struct` para búsqueda O(1) | 4 pts |
| Al menos un método de valor implementado sobre el struct | 4 pts |
| Función `agregar` que actualiza tanto slice como map | 4 pts |
| Función `buscar` usando el map (no range sobre slice) | 4 pts |
| Patrón `v, ok := m[k]` para manejar clave inexistente | 3 pts |
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

- Se descuenta **2 puntos** si se accede a un mapa sin verificar existencia donde el zero value produciría un resultado incorrecto
- Se descuenta **2 puntos** si el proyecto usa `range` sobre slice para buscar cuando hay un map indexado disponible
- Se descuenta **1 punto** por usar literal sin nombres de campo en structs con más de 2 campos: `Tipo{v1, v2, v3}` es frágil
- **Bonus +2 pts**: implementar un segundo método de valor para formatear el struct como string (`String() string`)
