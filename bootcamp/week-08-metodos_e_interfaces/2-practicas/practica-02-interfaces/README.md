# Práctica 02 — Interfaces: Duck Typing y Polimorfismo

## Descripción

En esta práctica declararás una interface propia, harás que dos tipos distintos la satisfagan implícitamente, escribirás una función polimórfica que trabaje con cualquier valor de la interface, y explorarás `any` para entender sus limitaciones.

---

## Paso 1: Declarar una interface

Una interface define un conjunto de firmas de métodos. Cualquier tipo que los tenga la satisface sin declararlo:

```go
type Medible interface {
    Medir() string
}
```

**Abre `starter/main.go`** y descomenta la sección del Paso 1.

---

## Paso 2: Implementar la interface con dos tipos distintos

Dos tipos completamente distintos pueden satisfacer la misma interface solo teniendo los métodos requeridos:

```go
type Distancia struct{ metros float64 }
type Masa struct{ kg float64 }

func (d Distancia) Medir() string { return fmt.Sprintf("%.2f m", d.metros) }
func (m Masa) Medir() string      { return fmt.Sprintf("%.2f kg", m.kg) }
```

**Descomenta el bloque del Paso 2** en `starter/main.go`.

---

## Paso 3: Función polimórfica

Una función que acepta la interface opera sobre cualquier tipo que la satisfaga:

```go
func mostrar(m Medible) {
    fmt.Println("Medición:", m.Medir())
}
```

Puedes también almacenar distintos tipos en un slice de interface:

```go
elementos := []Medible{Distancia{100}, Masa{75}}
for _, e := range elementos {
    mostrar(e)
}
```

**Descomenta el bloque del Paso 3** en `starter/main.go`.

---

## Paso 4: `any` — limitaciones

`any` acepta cualquier valor pero pierde el tipo concreto. Para operar necesitas una type assertion:

```go
var v any = Distancia{50}

// No puedes llamar Medir() directamente sobre any:
// v.Medir() // ❌ no compila

// Con type assertion sí:
if d, ok := v.(Distancia); ok {
    fmt.Println(d.Medir())
}
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
--- Paso 1 + 2: tipos implementan Medible ---
Medición: 100.00 m
Medición: 75.00 kg

--- Paso 3: slice de interface ---
Medición: 100.00 m
Medición: 75.00 kg

--- Paso 4: any y type assertion ---
Type assertion exitosa: 50.00 m
```
