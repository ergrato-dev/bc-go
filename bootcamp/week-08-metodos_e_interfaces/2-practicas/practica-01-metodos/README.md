# Práctica 01 — Métodos: Receptores, Method Set y `fmt.Stringer`

## Descripción

En esta práctica definirás métodos sobre structs con receptores de valor y de puntero, implementarás `fmt.Stringer` para representación textual personalizada, y observarás cómo el method set afecta la satisfacción de interfaces.

---

## Paso 1: Método con receptor de valor

Un método con receptor de valor `(r Tipo)` trabaja sobre una copia. Es adecuado para operaciones de solo lectura.

```go
type Rectangulo struct {
    ancho, alto float64
}

func (r Rectangulo) Area() float64 {
    return r.ancho * r.alto
}
```

**Abre `starter/main.go`** y descomenta el bloque del Paso 1.

---

## Paso 2: Método con receptor de puntero

Un método con receptor de puntero `(r *Tipo)` accede al original y puede modificarlo.

```go
func (r *Rectangulo) Escalar(factor float64) {
    r.ancho *= factor
    r.alto *= factor
}
```

**Descomenta el bloque del Paso 2** en `starter/main.go`.

---

## Paso 3: Implementar `fmt.Stringer`

Si un tipo tiene un método `String() string`, `fmt.Println` lo usa automáticamente para la representación textual.

```go
func (r Rectangulo) String() string {
    return fmt.Sprintf("Rect(%.1f x %.1f)", r.ancho, r.alto)
}

r := Rectangulo{ancho: 3, alto: 4}
fmt.Println(r) // Rect(3.0 x 4.0) — usa String()
```

**Descomenta el bloque del Paso 3** en `starter/main.go`.

---

## Paso 4: Verificar satisfacción de interface en compilación

El patrón `var _ Interface = Tipo{}` fuerza al compilador a verificar la satisfacción:

```go
var _ fmt.Stringer = Rectangulo{} // compila si tiene String() string
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
--- Paso 1: método de valor ---
Área: 12.00

--- Paso 2: método de puntero ---
Después de escalar x2: 6.0 x 8.0

--- Paso 3: fmt.Stringer ---
Rect(6.0 x 8.0)

--- Paso 4: verificación de interface ---
Stringer verificado en compilación
```
