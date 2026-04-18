# Práctica 01 — Punteros: Declarar, Derreferir y Nil Check

## Descripción

En esta práctica aprenderás a trabajar con los fundamentos de punteros en Go: obtener la dirección de una variable con `&`, acceder a su valor con `*`, modificar el original a través del puntero y protegerte contra nil pointer dereference.

---

## Paso 1: Obtener la dirección de una variable con `&`

El operador `&` retorna la dirección de memoria de una variable. El tipo resultante es un puntero: `&int` produce un `*int`.

```go
n := 42
p := &n

fmt.Printf("valor de n: %d\n", n)
fmt.Printf("dirección de n: %p\n", p)
fmt.Printf("tipo de p: %T\n", p) // *int
```

**Abre `starter/main.go`** y descomenta el bloque del Paso 1.

---

## Paso 2: Derreferir un puntero con `*`

El operador `*` aplicado a un puntero accede al valor en la dirección apuntada. Si modificas `*p`, estás modificando la variable original.

```go
n := 100
p := &n

fmt.Println("antes:", n)   // 100
*p = 200                   // modifica n a través de p
fmt.Println("después:", n) // 200 — n cambió
```

**Descomenta el bloque del Paso 2** en `starter/main.go`.

---

## Paso 3: Modificar el original a través de un puntero en función

Una función que recibe un puntero puede modificar la variable original del llamador. Esto no es posible si recibe el valor directamente.

```go
func triplicar(n *int) {
    *n = *n * 3
}

x := 5
triplicar(&x)
fmt.Println(x) // 15 — x fue modificado desde la función
```

**Descomenta el bloque del Paso 3** en `starter/main.go`.

---

## Paso 4: Verificar nil antes de derreferir

El zero value de cualquier puntero es `nil`. Derreferir un puntero nil produce un pánico en tiempo de ejecución. Siempre verifica antes:

```go
func imprimirValor(p *int) {
    if p == nil {
        fmt.Println("puntero nil — sin valor")
        return
    }
    fmt.Println("valor:", *p)
}

var nulo *int  // nil por defecto
v := 77
imprimirValor(nulo)  // puntero nil — sin valor
imprimirValor(&v)    // valor: 77
```

**Descomenta el bloque del Paso 4** en `starter/main.go`.

---

## Verificación

Al terminar todos los pasos, ejecuta:

```bash
cd starter
go run main.go
```

Deberías ver:

```
--- Paso 1: & dirección ---
valor de n: 42
p apunta a la misma dirección que n: true

--- Paso 2: * derreferencia ---
antes: 100
después: 200

--- Paso 3: modificar con función ---
x antes: 5
x después: 15

--- Paso 4: nil check ---
puntero nil — sin valor
valor: 77
```
