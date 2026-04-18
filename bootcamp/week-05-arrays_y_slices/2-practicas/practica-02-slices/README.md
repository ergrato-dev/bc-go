# Práctica 02 — Slices: make, append y copy

## Objetivo

Explorar la estructura interna de los slices: crear con `make` y literales, entender el aliasing de `a[lo:hi]`, desconectar slices con `copy`, filtrar con nil slice y construir matrices 2D.

## Instrucciones

Abre `starter/main.go` y descomenta las secciones en orden. Observa cómo `len` y `cap` cambian en cada operación.

---

### Paso 1: make, append y len/cap

`make([]T, len, cap)` preasigna capacidad. `append` crece el slice automáticamente.

```go
// Preasignar capacidad conocida evita reallocations innecesarias
s := make([]int, 0, 5)
fmt.Println(len(s), cap(s)) // 0 5

s = append(s, 1, 2, 3)
fmt.Println(len(s), cap(s)) // 3 5

// Al superar cap, Go crea nuevo backing array
s = append(s, 4, 5, 6)
fmt.Println(len(s), cap(s)) // 6 10 (cap aumentó)
```

**Descomenta la sección PASO 1** en `starter/main.go`.

---

### Paso 2: Aliasing — a[lo:hi] comparte backing array

`a[lo:hi]` crea un nuevo header que apunta al mismo backing array. Modificar el sub-slice modifica el original.

```go
nums := []int{10, 20, 30, 40, 50}
sub := nums[1:4] // ptr apunta a nums[1], len=3, cap=4

sub[0] = 99
fmt.Println(nums) // [10 99 30 40 50] — ¡modificado!
```

**Descomenta la sección PASO 2** y verifica que nums cambia al modificar sub.

---

### Paso 3: copy — slice verdaderamente independiente

`copy(dst, src)` copia los elementos sin compartir backing array. Retorna el número de elementos copiados.

```go
nums := []int{10, 20, 30, 40, 50}
independiente := make([]int, 3)
copy(independiente, nums[1:4])

independiente[0] = 99
fmt.Println(nums) // [10 20 30 40 50] — sin cambios
```

**Descomenta la sección PASO 3** y verifica la independencia.

---

### Paso 4: Filtrado idiomático con nil slice

El patrón de filtrado más eficiente usa `var resultado []T` (nil slice) y acumula con `append`.

```go
func filtrarPares(nums []int) []int {
    var resultado []int
    for _, n := range nums {
        if n%2 == 0 {
            resultado = append(resultado, n)
        }
    }
    return resultado
}
```

**Descomenta la sección PASO 4** y la función `filtrarPares` al final del archivo.

---

### Paso 5: Slice de slices — matriz 2D

Go no tiene arrays multidimensionales idiomáticos. La forma correcta es usar `[][]T`.

```go
filas := 3
cols := 4
matriz := make([][]int, filas)
for i := range matriz {
    matriz[i] = make([]int, cols)
}
```

**Descomenta la sección PASO 5** y la función `imprimirMatriz` al final del archivo.
