# Práctica 02 — `switch` y `defer`

> **Tipo**: Ejercicio guiado (descomenta el código paso a paso)

## 🎯 Objetivo

Explorar los patrones de `switch` (con expresión, sin expresión, múltiples valores por caso) y el comportamiento de `defer` (stack LIFO, evaluación de argumentos).

## 🚀 Ejecución

```bash
cd starter
go run .
```

---

### Paso 1: `switch` con expresión y múltiples valores por caso

En Go no se necesita `break` al final de cada `case` — el fall-through es opt-in con `fallthrough`.
Un `case` puede listar varios valores separados por coma.

```go
// Clasificar vocales y consonantes
letra := 'a'
switch letra {
case 'a', 'e', 'i', 'o', 'u':
    fmt.Println("Vocal")
default:
    fmt.Println("Consonante")
}
```

**Abre `starter/main.go`** y descomenta la sección `PASO 1`.

---

### Paso 2: `switch` sin expresión (como if/else)

Cuando se omite la expresión, cada `case` evalúa una condición booleana. Es más legible que una cadena de `if/else if`.

```go
// Clasificar temperatura
t := 38.2
switch {
case t >= 40.0:
    fmt.Println("Fiebre alta")
case t >= 37.5:
    fmt.Println("Fiebre moderada")
default:
    fmt.Println("Normal")
}
```

**Descomenta la sección `PASO 2`**.

---

### Paso 3: `defer` — orden LIFO

`defer` pospone la ejecución hasta que la función actual retorne. Múltiples `defer` se ejecutan en orden LIFO (último en declararse, primero en ejecutarse).

```go
func demostrarLIFO() {
    defer fmt.Println("primero declarado → último en ejecutar")
    defer fmt.Println("segundo declarado → segundo en ejecutar")
    defer fmt.Println("tercero declarado → primero en ejecutar")
    fmt.Println("ejecutando función...")
}
```

**Descomenta la sección `PASO 3`**.

---

### Paso 4: `defer` — evaluación de argumentos

Los argumentos de la función diferida se evalúan **en el momento de la declaración**, no al ejecutarse. Para capturar el valor final, se usa un closure.

```go
x := 10
defer fmt.Println("arg evaluado al declarar:", x)  // imprime 10
x = 99
// ↑ imprime 10, aunque x es 99 cuando se ejecuta el defer
```

**Descomenta la sección `PASO 4`**.

---

## ✅ Verificación

Al completar todos los pasos, la salida debe mostrar:

- Clasificación de varias letras como vocal/consonante (paso 1)
- Clasificación de temperaturas con `switch` sin expresión (paso 2)
- Tres mensajes de defer en orden inverso (paso 3)
- Demostración de captura de argumento vs closure (paso 4)

Ejecuta `go vet .` antes de continuar al proyecto.
