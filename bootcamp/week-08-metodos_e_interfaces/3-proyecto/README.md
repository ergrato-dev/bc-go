# Semana 08 — Proyecto: Interface Común para el Dominio

## Objetivo

Aplicar interfaces implícitas de Go a tu dominio asignado. Definirás una interface que represente un comportamiento común, implementarás esa interface en al menos dos structs distintos, y escribirás código polimórfico que opere sobre cualquier tipo que la satisfaga.

## Dominio de ejemplo (NO usar para tu entrega)

Este README usa el dominio **Acuario** como ejemplo. Adapta todo a tu dominio asignado.

---

## Requerimientos

### 1. Definir una interface con al menos un método

La interface debe representar un comportamiento genuino de tu dominio:

```go
// Ejemplo: Acuario
// Exhibible puede mostrarse al público con descripción y precio de entrada
type Exhibible interface {
    Describir() string
    PrecioEntrada() float64
}
```

Adapta a tu dominio:
- Biblioteca → `type Prestable interface { Prestar() string; Disponible() bool }`
- Farmacia → `type Vendible interface { Descripcion() string; Precio() float64 }`
- Gimnasio → `type Registrable interface { FichaResumen() string; Activo() bool }`

### 2. Dos structs distintos que implementen la interface

```go
// Exhibicion permanente
type Exhibicion struct {
    nombre    string
    sala      string
    precio    float64
}

// Evento temporal
type EventoEspecial struct {
    titulo    string
    duracion  string
    precio    float64
}

func (e Exhibicion) Describir() string      { return e.nombre + " — Sala " + e.sala }
func (e Exhibicion) PrecioEntrada() float64 { return e.precio }

func (ev EventoEspecial) Describir() string      { return ev.titulo + " (" + ev.duracion + ")" }
func (ev EventoEspecial) PrecioEntrada() float64 { return ev.precio }
```

### 3. `fmt.Stringer` en al menos un tipo

```go
func (e Exhibicion) String() string {
    return fmt.Sprintf("[Exhibición] %s — $%.2f", e.nombre, e.precio)
}
```

### 4. Función polimórfica

```go
func mostrarInfo(ex Exhibible) {
    fmt.Printf("  %s | Precio: $%.2f\n", ex.Describir(), ex.PrecioEntrada())
}
```

### 5. Slice de interface en `main`

```go
items := []Exhibible{
    Exhibicion{nombre: "Corales del Pacífico", sala: "B", precio: 15.0},
    EventoEspecial{titulo: "Noche de Medusas", duracion: "2h", precio: 25.0},
}

for _, item := range items {
    mostrarInfo(item)
}
```

### 6. Verificación de compilación

```go
var _ Exhibible = Exhibicion{}
var _ Exhibible = EventoEspecial{}
```

---

## Criterios de evaluación

| Criterio | Puntos |
|----------|--------|
| Interface declarada con ≥ 1 método coherente con el dominio | 10 pts |
| Dos structs distintos implementan la interface implícitamente | 15 pts |
| `fmt.Stringer` implementado en al menos un struct | 10 pts |
| Función polimórfica que acepta la interface | 15 pts |
| Slice de interface con valores de distintos tipos en main | 15 pts |
| Verificación `var _ Interface = Tipo{}` | 10 pts |
| Código compila y `go vet ./...` sin errores | 10 pts |
| Coherencia con el dominio asignado | 15 pts |
| **Total** | **100 pts** |

---

## Entrega

- `starter/main.go` con todos los TODOs resueltos
- `go run main.go` ejecuta sin errores
- `go vet ./...` sin advertencias
