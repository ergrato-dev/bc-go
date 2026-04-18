# Glosario — Semana 03: Control de Flujo

Términos clave ordenados alfabéticamente.

---

## B

### `break`
Termina inmediatamente el bucle o sentencia `switch` más cercano. En un `for` anidado, solo sale del bucle interno.

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // sale cuando i llega a 5
    }
}
```

---

## C

### `continue`
Salta el resto del cuerpo del bucle actual y pasa directamente a la siguiente iteración.

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue  // salta los pares
    }
    fmt.Println(i)  // solo imprime impares
}
```

---

## D

### `default`
Rama de un `switch` que se ejecuta cuando ningún `case` coincide. Opcional pero recomendado para manejar casos no previstos.

### `defer`
Sentencia que pospone la ejecución de una función hasta que la función que la contiene retorne. Múltiples `defer` se ejecutan en orden LIFO. Los argumentos se evalúan en el momento de la declaración, no al ejecutarse.

```go
defer fmt.Println("se ejecuta al salir")
defer recurso.Close()  // patrón idiomático de limpieza
```

---

## F

### `fallthrough`
Fuerza que la ejecución continúe en el siguiente `case` de un `switch`, sin evaluar su condición. Comportamiento opuesto al default de Go. Usar con moderación.

### `for`
La **única** estructura de bucle de Go. Puede adoptar 4 formas: clásico (init; cond; post), while (solo condición), infinito (sin condición) y range.

---

## I

### `if` con inicialización
Variante de `if` que ejecuta una sentencia de inicialización antes de evaluar la condición. La variable declarada solo existe dentro del bloque `if/else`.

```go
if v, err := os.Open(path); err != nil {
    return err
}
// v y err no existen aquí
```

---

## L

### LIFO (_Last In, First Out_)
Modelo de stack en que el último elemento añadido es el primero en salir. Es el orden en que Go ejecuta múltiples `defer` dentro de una función: el último `defer` declarado es el primero en ejecutarse.

---

## R

### `range`
Palabra clave usada con `for` para iterar sobre slices, arrays, maps, strings y channels. Retorna índice y valor en cada iteración. Para strings, itera por **runes**, no bytes.

```go
for i, v := range slice { }     // índice y valor
for k, v := range mapa  { }     // clave y valor
for _, r := range "café" { }    // runes Unicode
```

---

## S

### Scope de `if` / `switch`
Las variables declaradas en la sentencia de inicialización de `if` o `switch` solo existen dentro de ese bloque. Esto reduce el scope innecesario y mantiene el código más limpio.

### `switch`
Sentencia de selección que evalúa una expresión y ejecuta el primer `case` que coincida. No tiene fall-through implícito (no necesita `break`). Puede operar sin expresión para comportarse como `if/else if`.

---

## Z

### Zero value en control de flujo
Las variables de contador usadas en bucles (`i := 0`) no necesitan inicializarse explícitamente si se usan con `:=`. El compilador las inicializa al valor declarado. No confundir con variables de paquete que toman su zero value de forma automática.

---

> ← [Semana 02 — Glosario](../../week-02-variables_tipos_y_constantes/5-glosario/README.md)
