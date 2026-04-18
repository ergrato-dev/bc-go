# Proyecto Semana 05 — Gestor de Colección Funcional

## Contexto

Un acuario científico necesita un sistema para gestionar su colección de exhibiciones. Cada exhibición tiene un identificador único, un nombre y una categoría (peces, invertebrados, reptiles, etc.).

**IMPORTANTE**: Este es solo un **ejemplo de dominio**. Tu instructor te ha asignado un dominio diferente (biblioteca, farmacia, gimnasio, etc.). Adapta todos los tipos, nombres de campos y lógica de negocio a tu dominio asignado.

## Objetivo

Implementar un gestor de colección usando slices de Go: creación con `make`, inserción con `append`, búsqueda, filtrado, eliminación y clonación con `copy`.

## Estructura del proyecto

```
3-proyecto/starter/
└── main.go     ← único archivo a modificar
```

## Entregable

Un programa en Go que:

1. Inicialice una colección con `make` y capacidad apropiada
2. Agregue elementos con `append`
3. Busque un elemento por ID
4. Filtre por categoría retornando un nuevo slice
5. Elimine un elemento por ID preservando el orden
6. Clone la colección con `copy`
7. Imprima la colección completa

## Requisitos técnicos

- Usar `make([]Elemento, 0, cap)` para inicializar
- Usar `append` para agregar (siempre capturar el retorno)
- Implementar filtrado con nil slice + append
- Usar `copy` para clonar sin aliasing
- `go vet ./...` sin errores

## Criterios de evaluación

Ver [rubrica-evaluacion.md](../rubrica-evaluacion.md)

## Ejemplo de ejecución

```
=== Gestor de Exhibiciones del Acuario ===

[1] Agregar exhibiciones...
Colección (3 elementos):
  [1] Tiburón Toro       | Peces        |
  [2] Pulpo Gigante      | Invertebrados|
  [3] Tortuga de Mar     | Reptiles     |

[2] Buscar por ID...
Encontrado: [2] Pulpo Gigante (Invertebrados)

[3] Filtrar por categoría "Peces"...
Peces (1 resultado): [[1] Tiburón Toro]

[4] Eliminar ID 2...
Colección tras eliminar (2 elementos):
  [1] Tiburón Toro       | Peces        |
  [3] Tortuga de Mar     | Reptiles     |

[5] Clonar colección...
Modificar clon no afecta al original ✓
```

## Pistas

- Para eliminar preservando orden: `append(s[:i], s[i+1:]...)`
- Para clonar: `copy(dst, src)` donde `dst = make([]T, len(src))`
- Para filtrar: `var resultado []T` + `append` en el bucle `range`
