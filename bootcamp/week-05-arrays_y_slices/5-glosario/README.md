# Glosario — Semana 05: Arrays y Slices

## A

### Aliasing

Situación en la que dos variables (en este caso, dos slice headers) apuntan al mismo backing array. Modificar los datos a través de uno de los slices afecta al otro. Ocurre al hacer `sub := s[lo:hi]`.

```go
s := []int{1, 2, 3}
sub := s[0:2]    // aliasing — comparten backing array
sub[0] = 99
fmt.Println(s)   // [99 2 3] — s también cambia
```

---

## B

### Backing array

El array contiguo en memoria que almacena los datos de un slice. Un slice no "contiene" datos; contiene un puntero al backing array junto con `len` y `cap`. Múltiples slices pueden compartir el mismo backing array.

---

## C

### cap (capacidad)

Número total de elementos que el backing array puede almacenar desde el puntero del slice hasta su final. Siempre `cap >= len`. Se obtiene con la función builtin `cap(s)`.

```go
s := make([]int, 3, 10)
fmt.Println(len(s), cap(s)) // 3 10
```

### copy

Función builtin que copia elementos de un slice fuente (`src`) a uno destino (`dst`). Copia exactamente `min(len(dst), len(src))` elementos y retorna el número copiado. `dst` y `src` no comparten backing array tras la copia.

```go
n := copy(dst, src) // retorna elementos copiados
```

---

## L

### len (longitud)

Número de elementos accesibles en un slice o array. Se obtiene con la función builtin `len(s)`. Para un slice, `len <= cap` siempre.

---

## M

### make

Función builtin para crear slices (y también maps y channels) con tamaño y capacidad explícitos. `make([]T, len, cap)` crea un slice con `len` elementos inicializados a zero value y capacidad `cap`.

```go
s := make([]int, 0, 10) // len=0, cap=10 — preasignado para acumular
```

---

## N

### nil slice

Slice con valor cero: `ptr=nil`, `len=0`, `cap=0`. Es diferente de un slice vacío creado con `make([]T, 0)`. Ambos tienen `len=0` y `append` funciona sobre ambos, pero solo el nil slice es igual a `nil`.

```go
var s []int          // nil slice
t := make([]int, 0)  // slice vacío, no nil
fmt.Println(s == nil) // true
fmt.Println(t == nil) // false
```

---

## R

### Reallocation (realloc)

Proceso por el que `append` crea un nuevo backing array más grande cuando `len == cap`. Go copia todos los elementos al nuevo array y el slice original queda desconectado. Por esto `append` siempre debe reasignarse: `s = append(s, v)`.

---

## S

### Slice

Tipo de datos de Go que representa una ventana sobre un backing array. Su estructura interna tiene tres campos: `ptr` (puntero al primer elemento), `len` (longitud) y `cap` (capacidad). Los slices son referencias, no valores, por lo que asignarlos o pasarlos a funciones copia el header pero comparte el backing array.

### Slice header

La estructura de tres campos que representa internamente un slice: `(ptr *T, len int, cap int)`. Este header se copia al asignar un slice a otra variable o al pasarlo a una función.

### Slicing

Operación `a[lo:hi]` que crea un nuevo slice header apuntando a un rango del backing array existente. No copia datos. La variante `a[lo:hi:max]` permite controlar la capacidad del nuevo slice.

---

## T

### Three-index slice

Expresión `a[lo:hi:max]` que limita la capacidad del slice resultante a `max-lo`. Útil para evitar que `append` sobre el sub-slice sobreescriba elementos del slice original.

```go
s := []int{0, 1, 2, 3, 4}
sub := s[1:3:3] // len=2, cap=2 — append creará nuevo backing array
```

---

## Z

### Zero value de array/slice

Para un array: todos sus elementos tienen el zero value de su tipo (`0`, `""`, `false`, etc.). Para un slice: `nil` (pointer nulo, len=0, cap=0).

```go
var arr [3]int  // [0 0 0]
var s []string  // nil
```
