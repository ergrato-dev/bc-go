# Práctica 01 — Tipos Básicos en Go

## 🎯 Objetivo

Explorar el sistema de tipos de Go: declarar variables de todos los tipos básicos, observar sus zero values, realizar conversiones explícitas e imprimir con los verbos de formato correctos.

---

## 📋 Pasos

### Paso 1: Declarar variables de tipos básicos

Los tipos básicos de Go cubren enteros, flotantes, booleano y texto. Cada tipo tiene un tamaño fijo y un zero value garantizado.

```go
// Ejemplo: declarar variables con var (zero value automático)
var quantity int
var price float64
var name string
var available bool

fmt.Printf("int:     %v (%T)\n", quantity,  quantity)
fmt.Printf("float64: %v (%T)\n", price,     price)
fmt.Printf("string:  %q (%T)\n", name,      name)
fmt.Printf("bool:    %v (%T)\n", available, available)
```

**Abre `starter/main.go`** y descomenta la sección **PASO 1**.  
Ejecuta con: `go run .`

---

### Paso 2: Zero values — la promesa de Go

En Go, una variable declarada sin inicializar recibe automáticamente su zero value.
Esto elimina los bugs de "variable sin inicializar" comunes en C y C++.

```go
// Ejemplo: zero values de tipos básicos
var i int        // 0
var f float64    // 0
var b bool       // false
var s string     // ""
var p *int       // nil

fmt.Printf("int zero:     %d\n", i)
fmt.Printf("float64 zero: %f\n", f)
fmt.Printf("bool zero:    %t\n", b)
fmt.Printf("string zero:  %q\n", s) // %q muestra las comillas
fmt.Printf("puntero zero: %v\n", p) // nil
```

**Abre `starter/main.go`** y descomenta la sección **PASO 2**.  
Ejecuta con: `go run .`

---

### Paso 3: Conversión explícita de tipos

Go **no permite conversiones implícitas** entre tipos. Debes convertir explícitamente con `TipodDestino(valor)`.

```go
// Ejemplo: conversión de int a float64
items := 7
pricePerItem := 4.99

// ❌ Esto NO compila: mismatched types int and float64
// total := items * pricePerItem

// ✅ Conversión explícita
total := float64(items) * pricePerItem
fmt.Printf("Total: %.2f\n", total) // 34.93

// Conversión float64 → int (trunca, no redondea)
approximate := int(total)
fmt.Printf("Aprox: %d\n", approximate) // 34
```

**Abre `starter/main.go`** y descomenta la sección **PASO 3**.  
Ejecuta con: `go run .`

---

### Paso 4: Verbos de formato `fmt.Printf`

`fmt.Printf` usa verbos para controlar cómo se imprime cada valor.
`%T` es especialmente útil para depurar el tipo real de una variable.

```go
// Ejemplo: verbos de formato principales
x := 42
pi := 3.14159
ch := 'G'
msg := "Hola, Go"

fmt.Printf("%%v  → %v\n",   x)    // 42 (formato por defecto)
fmt.Printf("%%T  → %T\n",   x)    // int
fmt.Printf("%%d  → %d\n",   x)    // 42 (decimal)
fmt.Printf("%%b  → %b\n",   x)    // 101010 (binario)
fmt.Printf("%%x  → %x\n",   x)    // 2a (hexadecimal)
fmt.Printf("%%.2f → %.2f\n", pi)  // 3.14
fmt.Printf("%%c  → %c\n",   ch)   // G
fmt.Printf("%%q  → %q\n",   msg)  // "Hola, Go"
```

**Abre `starter/main.go`** y descomenta la sección **PASO 4**.  
Ejecuta con: `go run .`

---

### Paso 5: string, byte y rune — texto Unicode

Un `string` en Go es una secuencia de bytes (no de caracteres). Para texto Unicode, debes iterar con `range` para obtener `rune` (puntos de código Unicode).

```go
// Ejemplo: diferencia entre bytes y runes en un string Unicode
greeting := "Hola, 世界"

fmt.Printf("String: %q\n", greeting)
fmt.Printf("Bytes:  %d (len cuenta bytes)\n", len(greeting))
fmt.Printf("Runes:  %d ([]rune cuenta caracteres)\n", len([]rune(greeting)))

// Iterar por runes (forma idiomática para texto Unicode)
fmt.Println("Runes:")
for i, r := range greeting {
    fmt.Printf("  [%2d] %c  U+%04X\n", i, r, r)
}
```

**Abre `starter/main.go`** y descomenta la sección **PASO 5**.  
Ejecuta con: `go run .`

---

## ✅ Verificación

Cuando termines, la salida debe mostrar:
- Zero values de cada tipo (incluyendo `""` para string y `nil` para puntero)
- La conversión `int → float64` para calcular el total correctamente
- Los distintos verbos de formato aplicados al mismo valor entero
- La diferencia entre `len(string)` en bytes y `len([]rune(string))` en caracteres
