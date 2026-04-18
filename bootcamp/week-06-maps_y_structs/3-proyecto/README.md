# Proyecto — Semana 06: Registro con Maps y Structs

## 🎯 Objetivo

Construir un sistema de registro que combine un **struct** como tipo de dato, un **slice** para mantener el orden de inserción y un **map** como índice de búsqueda O(1). Este patrón es fundamental en el desarrollo real con Go.

## 📋 Contexto

El Museo de Historia Natural necesita un sistema para gestionar su colección de exhibiciones. Cada pieza tiene un identificador único, un nombre, una categoría y un indicador de si está en exposición actualmente.

> **⚠️ Nota para aprendices**: Este dominio es solo un ejemplo de referencia. Tú debes implementar el proyecto **usando tu dominio asignado** (farmacia, gimnasio, biblioteca, etc.), adaptando los nombres de campos, métodos y funciones a tu contexto.

## 🏗️ Requisitos

### Struct principal

Define un struct para representar un elemento de tu dominio con al menos 4 campos:

```go
type Pieza struct {
    ID         int
    Nombre     string
    Categoria  string
    EnExhibicion bool
    // Puedes agregar campos adicionales según tu dominio
}
```

### Funciones a implementar

| Función | Descripción |
|---------|-------------|
| `agregar(p Pieza, reg *[]Pieza, idx map[int]Pieza)` | Agrega al slice y al map |
| `buscar(id int, idx map[int]Pieza) (Pieza, bool)` | Busca por ID usando el map |
| `listar(reg []Pieza)` | Lista todas las piezas en orden |
| `Resumen() string` | Método de valor sobre el struct |

### Restricciones

- `buscar` debe usar el map, **no** `range` sobre el slice
- `agregar` debe actualizar **ambas** estructuras: slice y map
- Usar el patrón `v, ok := m[k]` en `buscar`
- Mínimo un método de valor sobre el struct
- Usar literales con nombre de campo al crear instancias del struct

## ✅ Criterios de aceptación

1. El programa compila y ejecuta sin errores: `go run starter/main.go`
2. `buscar` retorna la pieza correcta cuando existe y `false` cuando no
3. `listar` imprime todas las piezas en el orden en que fueron agregadas
4. El método `Resumen()` formatea la información del struct legiblemente
5. `go vet ./...` sin errores

## 📁 Estructura

```
3-proyecto/
├── README.md      ← este archivo
└── starter/
    └── main.go    ← implementa aquí
```

## 🚀 Cómo ejecutar

```bash
cd starter
go run main.go
```

## 💡 Pistas

- El slice y el map son estructuras separadas: ambas deben actualizarse al agregar
- `append` retorna una nueva slice — usa un puntero (`*[]Pieza`) para modificar el original
- El patrón `v, ok := idx[id]` en `buscar` es exactamente lo que aprendiste en la práctica 01
- Si implementas `String() string` en vez de `Resumen() string`, `fmt.Println` lo llamará automáticamente
