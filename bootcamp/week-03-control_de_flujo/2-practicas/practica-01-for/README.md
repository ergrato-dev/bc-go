# Práctica 01 — Patrones de `for`

> **Tipo**: Ejercicio guiado (descomenta el código paso a paso)

## 🎯 Objetivo

Explorar los 4 patrones del bucle `for` de Go, más `break` y `continue`, descomentando el código de cada paso y observando la salida.

## 🚀 Ejecución

```bash
cd starter
go run .
```

---

### Paso 1: `for` clásico (init; condición; post)

El `for` clásico usa tres componentes separados por `;`. Es idéntico al `for` de C, pero sin paréntesis.

```go
// Ejemplo — cuenta de 0 a 4
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

**Abre `starter/main.go`** y descomenta la sección `PASO 1`.

---

### Paso 2: `for` como while (solo condición)

Cuando se omiten `init` y `post`, el `for` se comporta como `while`.

```go
// Mientras el saldo sea positivo, descontar 10
saldo := 30
for saldo > 0 {
    saldo -= 10
    fmt.Println("Saldo restante:", saldo)
}
```

**Descomenta la sección `PASO 2`**.

---

### Paso 3: `for range` sobre slice

`range` es la forma idiomática de iterar sobre slices. Retorna índice y valor en cada iteración.

```go
// Iterar con índice y valor
planetas := []string{"Mercurio", "Venus", "Tierra", "Marte"}
for i, planeta := range planetas {
    fmt.Printf("[%d] %s\n", i, planeta)
}
```

**Descomenta la sección `PASO 3`**.

---

### Paso 4: `for range` sobre string (runes Unicode)

Al usar `range` sobre un `string`, Go itera por **runes** (puntos de código Unicode), no por bytes.

```go
// Cada iteración entrega la posición en bytes y el rune
for posicion, caracter := range "café" {
    fmt.Printf("byte[%d] = %c\n", posicion, caracter)
}
// La 'é' ocupa 2 bytes en UTF-8 → salto de posición
```

**Descomenta la sección `PASO 4`**.

---

### Paso 5: `break` y `continue`

`break` termina el bucle inmediatamente. `continue` salta al inicio de la siguiente iteración.

```go
// Buscar el primer número divisible por 7
for i := 1; i <= 100; i++ {
    if i%7 == 0 {
        fmt.Println("Primero divisible por 7:", i)
        break
    }
}

// Imprimir solo los impares (saltar pares con continue)
for i := 1; i <= 10; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println("Impar:", i)
}
```

**Descomenta la sección `PASO 5`**.

---

## ✅ Verificación

Al completar todos los pasos, la salida debe mostrar:

- Conteo 0–4 (paso 1)
- Saldo decrementando hasta 0 (paso 2)
- Lista de planetas con índice (paso 3)
- Caracteres de "café" con posiciones de byte (paso 4)
- Primer múltiplo de 7 y lista de impares del 1 al 10 (paso 5)

Ejecuta `go vet .` antes de continuar a la Práctica 02.
