# Rúbrica de Evaluación — Semana 01
## Introducción a Go y Herramientas

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
| ¿Qué problema resolvió Go en Google y por qué se diseñó de forma minimalista? | 5 |
| Explica la diferencia entre `go run` y `go build`. ¿Cuándo usarías cada uno? | 5 |
| ¿Qué es un módulo Go? ¿Cómo se diferencia de un paquete? | 5 |
| ¿Qué son los zero values y por qué existen en Go? | 5 |
| ¿Por qué `go.sum` debe commitearse junto con `go.mod`? | 5 |
| ¿Qué regla determina si un identificador es exportado o no en Go? | 5 |

---

## 💪 Desempeño (40 pts)

Evaluado durante la sesión práctica en clase.

### Práctica 01 — Hola Go (20 pts)

| Criterio | Pts |
|----------|-----|
| Descomenta correctamente el Paso 1 y ejecuta con `go run .` | 4 |
| Declara variables con `var` y `:=` correctamente (Paso 2) | 4 |
| Implementa y llama a función con múltiples retornos (Paso 3) | 6 |
| Maneja el error de `dividir` con `if err != nil` (Paso 4) | 6 |

### Práctica 02 — CLI con os.Args (20 pts)

| Criterio | Pts |
|----------|-----|
| Inicializa el módulo con `go mod init` correctamente | 4 |
| Lee `os.Args` con manejo seguro de índice (Paso 2) | 4 |
| Extrae lógica a `formatGreeting()` y la llama desde `main` (Paso 3) | 6 |
| Muestra versión y SO con `runtime.Version()` / `runtime.GOOS` (Paso 4) | 4 |
| Compila con `go build -o saluda .` y ejecuta el binario (Paso 5) | 2 |

---

## 📦 Producto (30 pts)

Entregable: **proyecto semanal** adaptado al dominio asignado.

### Criterios de evaluación del proyecto

| Criterio | Pts |
|----------|-----|
| El programa compila con `go run .` sin errores | 5 |
| `go vet .` pasa sin advertencias ni errores | 5 |
| `const systemName` reemplazado con el dominio asignado (no "Sistema de Gestión") | 5 |
| `os.Args` implementado: lee nombre del usuario o usa valor por defecto | 5 |
| `buildWelcome()` implementado: valida entradas y retorna mensaje formateado | 5 |
| Información adicional del dominio impresa (el TODO de `fmt.Println`) | 3 |
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
| — (primera semana) | [Semana 02 — Variables, Tipos y Constantes →](../week-02-variables_tipos_y_constantes/README.md) |
