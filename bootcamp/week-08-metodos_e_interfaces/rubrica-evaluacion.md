# Rúbrica de Evaluación — Semana 08: Métodos e Interfaces

## Distribución por tipo de evidencia

| Tipo | Peso | Puntos |
|------|------|--------|
| Conocimiento 🧠 | 30% | 30 pts |
| Desempeño 💪 | 40% | 40 pts |
| Producto 📦 | 30% | 30 pts |

---

## 🧠 Conocimiento (30 pts)

Evaluación teórica — 6 preguntas de 5 puntos cada una:

| # | Pregunta | Concepto evaluado |
|---|----------|-------------------|
| 1 | ¿Cuál es la diferencia entre definir un método con receptor `T` vs `*T` en cuanto al method set? | Method set |
| 2 | ¿Puede Go saber en tiempo de compilación si un tipo satisface una interface? ¿Cómo? | Satisfacción implícita |
| 3 | ¿Qué ocurre si una variable de tipo interface contiene un puntero nil? ¿Es la interface nil? | Interface con nil interno |
| 4 | ¿Por qué la interface `fmt.Stringer` es útil? ¿Qué método requiere? | Stringer |
| 5 | ¿Qué significa que las interfaces en Go son implícitas? ¿Qué ventaja aporta? | Duck typing |
| 6 | ¿Cuál es la diferencia entre `any` y un tipo concreto al almacenar en una variable? | `any` / `interface{}` |

---

## 💪 Desempeño (40 pts)

### Práctica 01 — Métodos (20 pts)

| Paso | Descripción | Puntos |
|------|-------------|--------|
| 1 | Define método con receptor de valor que retorna información | 5 pts |
| 2 | Define método con receptor de puntero que modifica estado | 5 pts |
| 3 | Implementa `fmt.Stringer` con método `String() string` | 5 pts |
| 4 | Demuestra que `T` no satisface interface que requiere `*T` en el method set | 5 pts |

### Práctica 02 — Interfaces (20 pts)

| Paso | Descripción | Puntos |
|------|-------------|--------|
| 1 | Declara interface con al menos un método | 5 pts |
| 2 | Implementa la interface con dos tipos distintos (implícitamente) | 5 pts |
| 3 | Función polimórfica que acepta la interface | 5 pts |
| 4 | Usa `any` para almacenar distintos tipos y muestra las limitaciones | 5 pts |

---

## 📦 Producto (30 pts)

### Proyecto integrador — Tipos con interface común

| Criterio | Puntos |
|----------|--------|
| Define al menos una interface coherente con el dominio | 8 pts |
| Al menos dos tipos distintos implementan la interface implícitamente | 8 pts |
| Función o loop que trabaja con valores de la interface (polimorfismo) | 7 pts |
| `fmt.Stringer` implementado en al menos un tipo | 4 pts |
| `go vet ./...` sin errores, código compilando | 3 pts |
| **Total** | **30 pts** |

---

## Criterios de aprobación

- Mínimo **21/30** en conocimiento (70%)
- Mínimo **28/40** en desempeño (70%)
- Mínimo **21/30** en producto (70%)
- Entrega puntual
- Implementación coherente con el dominio asignado
