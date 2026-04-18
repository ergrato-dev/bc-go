# Glosario — Semana 02: Variables, Tipos y Constantes

Términos clave ordenados alfabéticamente.

---

## B

### `bool`
Tipo booleano de Go. Solo puede tomar los valores `true` o `false`. Zero value: `false`.

```go
var activo bool   // false (zero value)
activo = true
```

### `byte`
Alias de `uint8`. Representa un byte de datos (valores 0–255). Se usa frecuentemente con strings y slices de bytes.

```go
var b byte = 'A'  // byte 65 (ASCII de 'A')
```

---

## C

### `const`
Palabra clave para declarar constantes. Los valores de una constante se evalúan en tiempo de compilación y no pueden cambiar en tiempo de ejecución.

```go
const Pi = 3.14159
const AppName = "MiApp"
```

### Conversión de tipos (_type conversion_)
En Go, la conversión entre tipos compatibles es **siempre explícita**. No existe conversión implícita.

```go
var x int = 42
var y float64 = float64(x)  // conversión explícita: int → float64
```

---

## F

### `float32` / `float64`
Tipos de punto flotante de 32 y 64 bits respectivamente. Usan representación IEEE 754. Se recomienda `float64` por defecto; `float32` solo cuando el ahorro de memoria es crítico.

```go
var pi float64 = 3.14159265358979
```

---

## I

### `int`
Tipo entero con signo cuyo tamaño depende de la plataforma (32 bits en sistemas de 32 bits, 64 bits en sistemas de 64 bits). Es el tipo entero recomendado por defecto.

```go
var count int = 0  // zero value: 0
```

### `int8`, `int16`, `int32`, `int64`
Tipos enteros con signo de tamaño fijo. Usar cuando el tamaño específico importa (protocolos de red, archivos binarios). `int32` es alias de `rune`.

### `iota`
Identificador predeclarado de Go que genera constantes enumeradas automáticamente dentro de un bloque `const`. Comienza en `0` y se incrementa en `1` por cada constante del bloque.

```go
type Direction int
const (
    North Direction = iota  // 0
    South                   // 1
    East                    // 2
    West                    // 3
)
```

### Inferencia de tipos (_type inference_)
Capacidad del compilador de deducir el tipo de una variable a partir de su valor inicial. Con `:=` Go infiere el tipo automáticamente.

```go
x := 42        // Go infiere int
y := 3.14      // Go infiere float64
name := "Go"   // Go infiere string
```

---

## R

### `rune`
Alias de `int32`. Representa un punto de código Unicode (un carácter). Usar `rune` en lugar de `byte` cuando el texto puede contener caracteres no-ASCII (acentos, CJK, emojis, etc.).

```go
var r rune = 'ñ'       // rune 241 (U+00F1)
var e rune = '🐹'      // rune del emoji gopher
```

---

## S

### `string`
Tipo inmutable que representa una secuencia de bytes (generalmente UTF-8). Zero value: `""` (cadena vacía).

```go
var nombre string       // "" (zero value)
nombre = "Go"
```

---

## T

### `type`
Palabra clave para definir un nuevo tipo con nombre. Permite crear tipos con semántica propia, incluso basados en tipos primitivos.

```go
type Celsius float64
type UserID int
```

---

## U

### `uint`, `uint8`, `uint16`, `uint32`, `uint64`
Tipos enteros **sin signo** (solo valores ≥ 0). `uint8` es alias de `byte`, `uint` tiene tamaño dependiente de la plataforma.

### Constante no tipada (_untyped constant_)
Constante que no tiene un tipo concreto hasta que se usa en un contexto que requiere uno. Las constantes no tipadas tienen mayor precisión y se adaptan al tipo necesario.

```go
const factor = 10       // no tipada; puede usarse como int, float64, etc.
const pi = 3.14         // no tipada float; más precisa que float64

const typedPi float32 = 3.14  // tipada; solo usable como float32
```

---

## V

### `var`
Palabra clave para declarar variables. Permite declarar variables en el nivel de paquete y dentro de funciones. La variable toma el zero value si no se inicializa.

```go
var nombre string           // zero value: ""
var edad int = 30
var activo bool             // zero value: false
```

---

## Z

### Zero value
El valor que asigna Go automáticamente a cualquier variable no inicializada. Garantiza que no existan variables con valores basura como en C/C++.

| Tipo | Zero value |
|------|-----------|
| `bool` | `false` |
| `int`, `int8` … `int64` | `0` |
| `uint`, `uint8` … `uint64` | `0` |
| `float32`, `float64` | `0.0` |
| `string` | `""` (cadena vacía) |
| `byte` | `0` |
| `rune` | `0` |
| puntero, slice, map, función, interface, channel | `nil` |

---

> ← [Semana 01 — Glosario](../../week-01-introduccion_go_y_herramientas/5-glosario/README.md)
