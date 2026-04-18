# Rúbrica de Evaluación — Semana 07: Punteros

## 📊 Distribución de Evidencias

| Tipo | Porcentaje |
|------|-----------|
| 🧠 Conocimiento | 30% |
| 💪 Desempeño | 40% |
| 📦 Producto | 30% |

---

## 🧠 Conocimiento (30%) — Cuestionario teórico

### Preguntas de evaluación

1. ¿Qué retorna el operador `&` aplicado a una variable?
2. ¿Qué hace el operador `*` aplicado a un puntero?
3. ¿Qué pasa si derreferencias un puntero `nil`?
4. ¿Cuál es la diferencia entre un método de valor y uno de puntero?
5. ¿Qué tipo retorna `new(int)` y cuál es su valor inicial?
6. ¿En qué caso Go desreferencia un puntero automáticamente (auto-deref)?

### Criterios de evaluación (30 pts total)

| Pregunta | Puntos |
|----------|--------|
| `&` retorna la dirección de memoria de la variable | 5 pts |
| `*` derreferencia: accede al valor en la dirección apuntada | 5 pts |
| Derreferencia de nil → pánico en tiempo de ejecución | 5 pts |
| Método de valor: copia; método de puntero: modifica el original | 5 pts |
| `new(int)` retorna `*int` con valor `0` | 5 pts |
| Auto-deref: al llamar un método de puntero sobre un valor direccionable | 5 pts |

---

## 💪 Desempeño (40%) — Prácticas en clase

### Práctica 01 — Punteros básicos (20 pts)

| Paso | Criterio | Puntos |
|------|----------|--------|
| Paso 1 | Declara un puntero con `&`, imprime dirección y valor derreferenciado | 5 pts |
| Paso 2 | Modifica el valor original a través del puntero con `*p = valor` | 5 pts |
| Paso 3 | Pasa un puntero a una función y verifica que el original cambia | 5 pts |
| Paso 4 | Detecta nil pointer y evita el pánico con verificación `if p != nil` | 5 pts |

### Práctica 02 — new() y métodos de puntero (20 pts)

| Paso | Criterio | Puntos |
|------|----------|--------|
| Paso 1 | Usa `new(T)` para crear un puntero a zero value y asigna mediante derreferencia | 5 pts |
| Paso 2 | Define método de valor y verifica que NO modifica el original | 5 pts |
| Paso 3 | Define método de puntero y verifica que SÍ modifica el original | 5 pts |
| Paso 4 | Demuestra auto-deref: llamar método de puntero sobre variable (no puntero) | 5 pts |

---

## 📦 Producto (30%) — Proyecto integrador

### Criterios de evaluación (30 pts)

| Criterio | Puntos |
|----------|--------|
| Struct principal con al menos 4 campos | 3 pts |
| Función `agregar` con receptor de puntero o parámetro puntero que modifica el original | 4 pts |
| Al menos un método de puntero `(t *Tipo)` implementado | 5 pts |
| Al menos un método de valor `(t Tipo)` implementado | 4 pts |
| Función que recibe `*Tipo` para modificar el struct (no retorna copia) | 5 pts |
| Verificación de nil pointer antes de derreferencia cuando aplique | 3 pts |
| `go vet ./...` sin errores | 3 pts |
| Código adaptado coherentemente al dominio asignado | 3 pts |

---

## ✅ Criterios de Aprobación

- Mínimo **21/30** en conocimiento
- Mínimo **28/40** en desempeño
- Mínimo **21/30** en producto
- `go vet ./...` sin errores
- Proyecto funcional con el dominio asignado

---

## 📝 Notas de evaluación

- Se descuenta **3 puntos** si hay un posible nil pointer dereference sin verificación
- Se descuenta **2 puntos** si el método de puntero no produce ningún cambio observable en el original (indica que no fue implementado correctamente)
- Se descuenta **1 punto** por usar `*&variable` cuando se puede acceder directamente a la variable
- **Bonus +2 pts**: demostrar que `new(T)` y `&T{}` son equivalentes en el proyecto
