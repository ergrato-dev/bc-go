# Práctica 01 — Múltiples Retornos y Variádicas

## Objetivo

Explorar los cuatro patrones más importantes de declaración de funciones en Go: múltiples retornos, retornos nombrados, funciones variádicas y funciones como valores.

---

## Instrucciones

Abre `starter/main.go` y descomenta cada sección siguiendo los pasos en orden. Ejecuta el programa con `go run main.go` al finalizar cada paso para verificar el resultado.

---

### Paso 1: Función con múltiples retornos

Go permite retornar más de un valor, lo que elimina la necesidad de excepciones para el manejo de errores. El patrón `(resultado, error)` es el más común en la librería estándar.

```go
// dividir retorna el resultado de la división y un error si el divisor es cero.
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("dividir: divisor no puede ser cero")
    }
    return a / b, nil
}
```

**Abre `starter/main.go`** y descomenta la sección `PASO 1`.

Resultado esperado:
```
10 / 4 = 2.50
Error esperado: dividir: divisor no puede ser cero
```

---

### Paso 2: Retornos nombrados

Los retornos nombrados actúan como variables locales declaradas al inicio. El `return` desnudo retorna sus valores actuales.

```go
// estadisticas retorna la suma y el promedio de un slice.
// Los nombres documentan qué representa cada valor.
func estadisticas(nums []float64) (suma, promedio float64) {
    for _, n := range nums {
        suma += n
    }
    if len(nums) > 0 {
        promedio = suma / float64(len(nums))
    }
    return // naked return
}
```

**Descomenta la sección `PASO 2`.**

Resultado esperado:
```
Suma: 15.00, Promedio: 3.00
```

---

### Paso 3: Función variádica

Una función variádica acepta un número variable de argumentos. Dentro de la función, el parámetro `...T` es un `[]T`.

```go
// unir concatena strings con un separador.
func unir(sep string, partes ...string) string {
    // partes es []string aquí dentro
    ...
}
```

**Descomenta la sección `PASO 3`.**

Resultado esperado:
```
a-b-c
Go es potente y rápido
```

---

### Paso 4: Expansión de slice con `...`

Cuando ya tienes un slice y quieres pasarlo a una función variádica, usa el operador `...` para expandirlo.

```go
palabras := []string{"hola", "mundo"}
resultado := unir(" ", palabras...) // expande el slice
```

**Descomenta la sección `PASO 4`.**

Resultado esperado:
```
hola mundo Go
```

---

### Paso 5: Función como valor

Las funciones son valores de primera clase: se pueden asignar a variables, pasar como argumentos y retornar desde otras funciones.

```go
// tipo de función: toma un int y retorna un int
type transformFn func(int) int

func aplicar(nums []int, fn transformFn) []int { ... }
```

**Descomenta la sección `PASO 5`.**

Resultado esperado:
```
Originales: [1 2 3 4 5]
Dobles:     [2 4 6 8 10]
Cuadrados:  [1 4 9 16 25]
```
