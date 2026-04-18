# Glosario — Semana 04: Funciones

Términos técnicos clave de esta semana, ordenados alfabéticamente.

---

## C

### Closure

Función que captura variables de su entorno léxico (el scope donde fue definida). Las variables capturadas persisten en el heap mientras el closure exista.

```go
func contador() func() int {
    n := 0             // n escapa al heap
    return func() int { n++; return n }
}
```

---

## E

### Entorno léxico (Lexical environment)

El conjunto de variables visibles desde el punto donde se define una función. Un closure "cierra sobre" estas variables, por eso se llama closure.

---

## F

### First-class function (función de primera clase)

Una función que puede ser tratada como un valor: asignada a una variable, pasada como argumento y retornada por otra función.

```go
op := fmt.Println  // asignar a variable
op("hola")         // llamar desde variable
```

### Función anónima (function literal)

Función declarada sin nombre. Se usa cuando la lógica es corta y solo se necesita una vez.

```go
doble := func(n int) int { return n * 2 }
```

### Función de orden superior (HOF — Higher-Order Function)

Función que recibe una función como argumento y/o retorna una función. Permite abstraer el patrón de una operación.

```go
func aplicar(nums []int, fn func(int) int) []int { ... }
```

### Función variádica

Función que acepta un número variable de argumentos del mismo tipo usando `...T`. Dentro de la función, el parámetro es `[]T`.

```go
func sum(nums ...int) int { ... }
sum(1, 2, 3)     // llamada normal
sum(slice...)    // expansión de slice
```

---

## H

### Heap escape (escape al heap)

Cuando el compilador detecta que una variable local será usada fuera de su función (ej: capturada por un closure), la mueve automáticamente al heap para que su vida útil se extienda.

---

## N

### Naked return (retorno desnudo)

`return` sin valores en una función con retornos nombrados. Retorna los valores actuales de las variables nombradas.

```go
func f() (x int) { x = 5; return } // retorna 5
```

---

## O

### Operador de expansión `...`

Usado al llamar una función variádica para expandir un slice como argumentos individuales.

```go
nums := []int{1, 2, 3}
sum(nums...) // equivale a sum(1, 2, 3)
```

---

## R

### Retorno nombrado (named return)

Valor de retorno con nombre en la firma de la función. Actúa como variable local inicializada al zero value.

```go
func divide(a, b float64) (resultado float64, err error) { ... }
```

### Retorno múltiple (multiple return)

Capacidad de Go de retornar más de un valor desde una función. El patrón estándar es `(valor, error)`.

```go
func open(path string) (*File, error) { ... }
f, err := open("datos.txt")
```

---

## T

### Tipo de función (function type)

Tipo definido por la firma de una función: sus parámetros y valores de retorno. Se puede nombrar con `type`.

```go
type predFn func(int) bool  // función que toma int, retorna bool
```
