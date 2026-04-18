# Semana 07 — Proyecto: Registro con Métodos de Puntero

## Objetivo

Extender el struct de tu dominio asignado con métodos de puntero que permitan modificar el estado del receptor. Al finalizar, tu struct tendrá tanto métodos de valor (para consultar) como métodos de puntero (para actualizar).

## Contexto

Usarás como base el struct que desarrollaste en la semana anterior (maps y structs). Ahora agregarás comportamiento: métodos que modifican el estado interno del struct a través de un receptor puntero.

## Dominio de ejemplo (NO usar para tu entrega)

Este README usa el dominio **Museo** como ejemplo genérico para ilustrar los conceptos. Adapta todo al dominio que te fue asignado.

---

## Requerimientos

### 1. Definir el struct con al menos 3 campos

```go
// Ejemplo: Museo (usa tu dominio asignado)
type Obra struct {
    titulo     string
    artista    string
    enExhibicion bool
}
```

Adapta el struct a tu dominio (Biblioteca → Libro, Farmacia → Medicamento, etc.).

### 2. Método de consulta con receptor de valor

Implementa al menos un método que **solo lea** información sin modificar el receptor:

```go
// Receptor de valor — no puede (ni necesita) modificar Obra
func (o Obra) descripcion() string {
    estado := "en bodega"
    if o.enExhibicion {
        estado = "en exhibición"
    }
    return fmt.Sprintf("%s de %s — %s", o.titulo, o.artista, estado)
}
```

### 3. Método de actualización con receptor de puntero

Implementa al menos dos métodos que **modifiquen** el estado del receptor:

```go
// Receptor de puntero — modifica el struct original
func (o *Obra) activar() {
    o.enExhibicion = true
}

func (o *Obra) desactivar() {
    o.enExhibicion = false
}
```

### 4. Función que recibe puntero y modifica

Implementa una función (no método) que reciba un puntero a tu struct y lo modifique:

```go
// TODO: implementar según tu dominio
// Ejemplo: func actualizarArtista(o *Obra, nuevoArtista string)
```

### 5. Demostrar nil check

Tu `main` debe demostrar que verificas nil antes de usar un puntero:

```go
var p *Obra // nil
if p != nil {
    fmt.Println(p.descripcion())
} else {
    fmt.Println("sin obra asignada")
}
```

---

## Código inicial

El archivo `starter/main.go` contiene la estructura base. Debes:

1. Adaptar `Item` al struct de tu dominio
2. Completar todos los `// TODO`
3. Ejecutar `go run main.go` sin errores
4. Verificar con `go vet ./...` sin advertencias

---

## Criterios de evaluación

| Criterio | Puntos |
|----------|--------|
| Struct con ≥ 3 campos coherentes con el dominio | 15 |
| Método de valor que lee sin modificar | 15 |
| Método de puntero que modifica el receptor | 20 |
| Función que recibe puntero y modifica | 20 |
| Nil check explícito en main | 15 |
| Código compila y `go vet` sin errores | 15 |
| **Total** | **100** |

---

## Entrega

- Archivo: `starter/main.go` con todos los TODOs resueltos
- El código debe compilar y ejecutar sin errores
- `go vet ./...` sin advertencias
