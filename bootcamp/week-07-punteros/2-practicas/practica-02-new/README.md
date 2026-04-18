# Práctica 02 — `new()`, Método de Valor vs Puntero y Auto-deref

## Descripción

En esta práctica explorarás `new(T)` para crear punteros a zero value, la diferencia observable entre métodos de valor (que trabajan sobre copias) y métodos de puntero (que modifican el receptor original), y el mecanismo de auto-deref de Go.

---

## Paso 1: `new(T)` — puntero al zero value

`new(T)` asigna memoria para un valor de tipo `T` en zero value y retorna un `*T`. Es equivalente a `&T{}` para structs.

```go
p := new(int)        // *int, *p == 0
*p = 42
fmt.Println(*p)      // 42

tipo := new(string)  // *string, *tipo == ""
*tipo = "Go"
fmt.Println(*tipo)   // Go
```

**Abre `starter/main.go`** y descomenta el bloque del Paso 1.

---

## Paso 2: Método de valor — no modifica el receptor

Un método con receptor de valor `(c Caja)` recibe una copia del struct. Las modificaciones no afectan al original.

```go
type Caja struct {
    contenido string
}

// receptor de valor — trabaja sobre copia
func (c Caja) vaciar() {
    c.contenido = "" // solo modifica la copia local
}
```

Observa que después de llamar `vaciar()`, el original no cambia.

**Descomenta el bloque del Paso 2** en `starter/main.go`.

---

## Paso 3: Método de puntero — sí modifica el receptor

Un método con receptor de puntero `(c *Caja)` accede al struct original. Las modificaciones sí persisten.

```go
// receptor de puntero — trabaja sobre el original
func (c *Caja) llenar(v string) {
    c.contenido = v // modifica el struct original
}
```

Go aplica auto-deref: cuando llamas `caja.llenar("arroz")`, Go convierte a `(&caja).llenar("arroz")` automáticamente si `caja` es una variable.

**Descomenta el bloque del Paso 3** en `starter/main.go`.

---

## Paso 4: Auto-deref — acceso a campos sin `(*p).campo`

Cuando tienes un puntero a struct, puedes acceder a sus campos directamente con `.` sin derreferir manualmente. Go lo hace por ti.

```go
type Punto struct{ X, Y int }

p := &Punto{X: 1, Y: 2}

// Estas dos líneas son equivalentes:
p.X = 10       // Go convierte a: (*p).X = 10
(*p).X = 10    // equivalente explícito
fmt.Println(p.X) // 10
```

**Descomenta el bloque del Paso 4** en `starter/main.go`.

---

## Verificación

```bash
cd starter
go run main.go
```

Salida esperada:

```
--- Paso 1: new(T) ---
*p == 0 al inicio
*p == 42 después de asignar

--- Paso 2: método de valor (no modifica) ---
antes de vaciar: lleno
después de vaciar: lleno

--- Paso 3: método de puntero (sí modifica) ---
antes de llenar: 
después de llenar: arroz

--- Paso 4: auto-deref ---
p.X == 10
```
