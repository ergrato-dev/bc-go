# Práctica 02 — Closures y Funciones de Orden Superior

## Objetivo

Construir y usar closures para encapsular estado, crear generadores y componer comportamiento mediante funciones de orden superior.

---

## Instrucciones

Abre `starter/main.go` y descomenta cada sección siguiendo los pasos en orden. Ejecuta con `go run main.go` al finalizar cada paso.

---

### Paso 1: Closure como contador

Un closure captura la variable de su entorno y la recuerda entre llamadas. Cada llamada a `contadorNuevo()` crea un entorno independiente con su propia variable `n`.

```go
func contadorNuevo() func() int {
    n := 0 // esta variable vive en el heap
    return func() int {
        n++
        return n
    }
}
```

**Descomenta la sección `PASO 1`.**

Resultado esperado:
```
contarA: 1
contarA: 2
contarB: 1   ← entorno independiente
contarA: 3   ← A continúa desde donde estaba
```

---

### Paso 2: Fábrica de funciones (closure que retorna closure)

Una función puede retornar una función con comportamiento parametrizado. El parámetro queda "atrapado" en el closure.

```go
func multiplicadorPor(factor int) func(int) int {
    return func(n int) int {
        return n * factor // 'factor' capturado
    }
}
```

**Descomenta la sección `PASO 2`.**

Resultado esperado:
```
triplicar(4) = 12
triplicar(7) = 21
prefijo WARN: [WARN] disco lleno
prefijo INFO: [INFO] inicio ok
```

---

### Paso 3: HOF — filtrar, transformar, reducir

Las funciones de orden superior reciben funciones como argumentos. Esto permite abstraer el _patrón_ de iteración separándolo de la _lógica_.

```go
func filtrar(nums []int, pred func(int) bool) []int { ... }
func transformar(nums []int, fn func(int) int) []int { ... }
func reducir(nums []int, inicial int, fn func(int, int) int) int { ... }
```

**Descomenta la sección `PASO 3`.**

Resultado esperado:
```
Originales: [1 2 3 4 5 6 7 8 9 10]
Pares:      [2 4 6 8 10]
Cuadrados:  [4 16 36 64 100]
Suma total: 220
```

---

### Paso 4: Closure como middleware

Un closure puede envolver cualquier función para agregar comportamiento antes y/o después sin modificarla. Este patrón es la base del middleware en servidores HTTP.

```go
func conLogger(nombre string, fn func() error) func() error {
    return func() error {
        fmt.Printf("→ %s\n", nombre)
        err := fn()
        // ...
        return err
    }
}
```

**Descomenta la sección `PASO 4`.**

Resultado esperado:
```
→ procesarPedido
✓ procesarPedido completado
→ validarStock
✗ validarStock falló: stock insuficiente
```
