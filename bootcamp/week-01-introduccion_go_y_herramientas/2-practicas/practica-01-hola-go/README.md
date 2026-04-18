# Práctica 01 — Hola Go: Primer Programa

## 🎯 Objetivo

Escribir, ejecutar y entender la estructura de un programa Go completo, usando el toolchain oficial paso a paso.

## 📋 Instrucciones

Cada paso te guía a descomentar una sección en `starter/main.go`. Lee la explicación, entiende el concepto, luego descomenta y ejecuta con `go run .`.

---

### Paso 1: El programa mínimo

Todo programa Go ejecutable necesita `package main` y `func main()`. El compilador impide compilar sin ellos.

```go
// Un programa Go mínimo y completo
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, Go!")
}
```

**Abre `starter/main.go`** y descomenta la sección **PASO 1**.

Ejecuta:

```bash
go run .
```

Deberías ver: `¡Hola, Go!`

---

### Paso 2: Declarar variables

Go tiene dos formas de declarar variables. Aprenderemos ambas y cuándo usar cada una.

```go
// var — forma explícita, útil cuando el tipo no es obvio
var nombre string = "Gopher"

// := — forma corta con inferencia, la más común dentro de funciones
version := "1.24"
```

**Descomenta la sección PASO 2** en `starter/main.go` y ejecuta.

---

### Paso 3: Múltiples valores de retorno

Go permite que las funciones retornen múltiples valores. Es el mecanismo base para el manejo de errores.

```go
// saludar retorna el saludo y su longitud
func saludar(nombre string) (string, int) {
    saludo := "Hola, " + nombre + "!"
    return saludo, len(saludo)
}
```

**Descomenta la sección PASO 3** y ejecuta.

---

### Paso 4: Manejo de errores

Los errores son valores en Go. Siempre se retornan y siempre se verifican.

```go
// La función retorna (resultado, error)
// Si algo falla, el error no es nil
resultado, err := operacion()
if err != nil {
    fmt.Println("Error:", err)
    return
}
```

**Descomenta la sección PASO 4** y ejecuta.

---

### Paso 5: `go fmt` y `go vet`

Aplica las herramientas de calidad al código:

```bash
# Desde el directorio starter/
go fmt .
go vet .
```

Ambos deberían ejecutarse sin errores ni salida (silencio = éxito).
