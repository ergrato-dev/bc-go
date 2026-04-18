# Práctica 02 — CLI con os.Args y go mod

## 🎯 Objetivo

Crear un programa Go que reciba argumentos de línea de comandos, organice código
en una función separada y use correctamente un módulo con `go mod init`.

**Tiempo estimado:** 45 minutos

---

## 📋 Pasos

### Paso 1: Inicializar el módulo

Antes de ejecutar cualquier programa Go en un directorio nuevo, necesitamos un
módulo. El módulo le dice al compilador cómo se llama nuestro proyecto y permite
gestionar dependencias.

```bash
cd starter/
go mod init practica02
```

Esto crea `go.mod` con el nombre del módulo. Verifica el contenido:

```bash
cat go.mod
```

Deberías ver:

```
module practica02

go 1.24
```

**Abre `starter/main.go`** y descomenta la sección del **Paso 1**.

Ejecuta y verifica:

```bash
go run .
```

---

### Paso 2: Leer argumentos con os.Args

`os.Args` es un slice de strings que el sistema operativo pasa al programa
cuando lo ejecutas. El índice `[0]` siempre es el nombre del binario.

```go
// os.Args tiene esta forma cuando ejecutas: go run . Ana
// os.Args[0] = "./practica02"  (el binario)
// os.Args[1] = "Ana"           (primer argumento del usuario)
```

**Abre `starter/main.go`** y descomenta la sección del **Paso 2**.

Ejecuta pasando tu nombre:

```bash
go run . TuNombre
go run .           # sin argumento — debe usar "Visitante"
```

---

### Paso 3: Extraer lógica a una función

Separar la lógica de construcción del mensaje de `main()` es una buena práctica:
facilita el testing, la legibilidad y la reutilización.

```go
// formatGreeting recibe un nombre y retorna el saludo formateado.
// Patrón: lógica en función, main() solo orquesta.
func formatGreeting(name string) string { ... }
```

**Abre `starter/main.go`** y descomenta la sección del **Paso 3**.

```bash
go run . Gopher
```

---

### Paso 4: Información del sistema con runtime

El paquete `runtime` expone metadatos del entorno de ejecución de Go.
Útil para mostrar información de diagnóstico o versión.

```go
import "runtime"

runtime.Version()   // ej: "go1.24.0"
runtime.GOOS        // ej: "linux", "darwin", "windows"
runtime.GOARCH      // ej: "amd64", "arm64"
```

**Abre `starter/main.go`** y descomenta la sección del **Paso 4**.

```bash
go run . Gopher
```

---

### Paso 5: Compilar un binario real

Hasta ahora usamos `go run`. Ahora compilamos un ejecutable:

```bash
go build -o saluda .
./saluda Ana
./saluda              # sin argumento
```

`go build` produce un binario nativo. Comparte el binario con alguien que no
tenga Go instalado — funciona igual.

**Abre `starter/main.go`** y lee los comentarios del **Paso 5** (no hay código que
descomentar aquí, el objetivo es ejecutar los comandos arriba).

---

## 🔍 Verificación final

```bash
go vet .           # debe pasar sin errores ni advertencias
go run . TuNombre  # muestra saludo + info del sistema
go build -o saluda . && ./saluda TuNombre
```

---

## 📚 Conceptos cubiertos

| Concepto         | Dónde aparece                     |
|------------------|-----------------------------------|
| `os.Args`        | Pasos 2, 3                        |
| `go mod init`    | Paso 1                            |
| Funciones        | Paso 3                            |
| `runtime`        | Paso 4                            |
| `go build`       | Paso 5                            |
| Múltiples imports| Pasos 4 y 5                       |
