# Práctica 02 — Structs: Definición, Métodos de Valor y Semántica

Aprenderás a definir structs, inicializarlos con literales con nombre de campo, verificar la semántica de valor e implementar métodos de valor.

## Prerequisitos

- Práctica 01 completada
- Teoría `1-teoria/02-structs-y-metodos-de-valor.md` leída

## Instrucciones generales

Abre `starter/main.go`. Descomenta las líneas de cada paso uno a uno, ejecuta con `go run starter/main.go` y verifica la salida antes de avanzar.

---

### Paso 1 — Definir un struct y usar zero value

Conceptos: `type Nombre struct`, campos tipados, zero value automático, acceso con punto.

```go
// Un struct agrupa campos con nombre y tipo fijos.
// Todos conocidos en compilación — el compilador verifica typos en nombres.

type Producto struct {
    ID       int
    Nombre   string
    Precio   float64
    EnStock  bool
}

// Zero value: todos los campos en su estado base
var p Producto
fmt.Println(p.Nombre == "")  // true
fmt.Println(p.Precio == 0)   // true
fmt.Println(p.EnStock)       // false
```

**Abre `starter/main.go`** y descomenta la sección `PASO 1`.

Salida esperada:
```
--- Paso 1: Zero value ---
Nombre vacío: true
Precio cero: true
En stock: false
```

---

### Paso 2 — Literal con nombre de campo

Conceptos: inicialización nombrada, campos omitidos usan zero value.

```go
// Siempre nombrar los campos — robusto ante cambios al struct
p1 := Producto{
    ID:      1,
    Nombre:  "Cuaderno A4",
    Precio:  2.50,
    EnStock: true,
}

// Inicialización parcial — Precio y EnStock quedan en zero value
p2 := Producto{
    ID:     2,
    Nombre: "Lápiz",
}

fmt.Println(p1.Nombre, p1.Precio) // "Cuaderno A4" 2.5
fmt.Println(p2.EnStock)           // false
```

**Abre `starter/main.go`** y descomenta la sección `PASO 2`.

Salida esperada:
```
--- Paso 2: Literal ---
p1: Cuaderno A4 $2.50
p2 en stock: false
```

---

### Paso 3 — Método de valor

Conceptos: receptor `(p Producto)`, método que retorna string, `fmt.Stringer`.

```go
// El método de valor recibe una copia — no modifica el original.
// Si el método se llama String(), fmt.Println lo usa automáticamente.

func (p Producto) String() string {
    disponible := "sin stock"
    if p.EnStock {
        disponible = "disponible"
    }
    return fmt.Sprintf("[%d] %-20s $%.2f (%s)",
        p.ID, p.Nombre, p.Precio, disponible)
}
```

**Abre `starter/main.go`** y descomenta la sección `PASO 3`.

Salida esperada:
```
--- Paso 3: Método de valor ---
[1] Cuaderno A4          $2.50 (disponible)
[2] Lápiz                $0.00 (sin stock)
```

---

### Paso 4 — Semántica de valor y comparación con `==`

Conceptos: copia al asignar, independencia de copias, comparación campo a campo.

```go
// Asignar un struct crea una copia completa — a y b son independientes
a := Producto{ID: 1, Nombre: "Cuaderno A4", Precio: 2.50, EnStock: true}
b := a
b.Nombre = "Cuaderno A5"
b.Precio = 3.00

fmt.Println(a.Nombre) // "Cuaderno A4" — a no cambió
fmt.Println(b.Nombre) // "Cuaderno A5" — b es independiente

// Comparación con == — compara campo a campo
c := Producto{ID: 1, Nombre: "Cuaderno A4", Precio: 2.50, EnStock: true}
fmt.Println(a == c) // true — todos los campos iguales
fmt.Println(a == b) // false — Nombre y Precio difieren
```

**Abre `starter/main.go`** y descomenta la sección `PASO 4`.

Salida esperada:
```
--- Paso 4: Semántica de valor ---
a.Nombre: Cuaderno A4
b.Nombre: Cuaderno A5
a == c: true
a == b: false
```

---

## Verificación final

```bash
go run starter/main.go
```

Todos los pasos deben imprimirse sin errores.
