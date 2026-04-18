# Glosario — Semana 06: Maps y Structs

Términos técnicos introducidos esta semana, ordenados alfabéticamente.

---

## C

### Campo (field)
Variable con nombre y tipo que forma parte de un struct. Los campos se acceden con el operador punto: `p.Nombre`. Los campos en PascalCase son exportados (visibles fuera del paquete); los en camelCase son privados.

### Comparable
Un tipo es comparable si puede usarse con los operadores `==` y `!=`. Los tipos básicos (int, string, bool) son comparables. Un struct es comparable si todos sus campos lo son. Los maps, slices y funciones **no** son comparables.

## D

### delete
Función builtin que elimina una entrada de un map: `delete(m, clave)`. Si la clave no existe, no hace nada (no produce error ni pánico). Para eliminar todas las entradas usar `clear(m)` (Go 1.21+).

## E

### Embedding (incrustación)
Técnica de Go para componer tipos incluyendo un tipo dentro de otro sin nombre de campo explícito. El tipo incrustado "promueve" sus métodos al tipo contenedor. Se verá en detalle en la Semana 09.

### Exportado
Un identificador es exportado cuando su nombre empieza con mayúscula (PascalCase). Es visible desde otros paquetes. Ejemplo: `Nombre`, `Precio`, `ID`. Ver también: *no exportado*.

## H

### Hash table (tabla hash)
Estructura de datos interna que usa maps en Go. Asocia claves con valores calculando la posición mediante una función hash. Proporciona acceso O(1) en promedio. El orden de las entradas no está garantizado.

## L

### Literal de struct
Expresión para crear e inicializar un struct en una sola línea. Siempre usar la forma con nombres de campo: `Producto{ID: 1, Nombre: "Go"}`. La forma posicional `Producto{1, "Go"}` es frágil.

## M

### map
Tipo de datos en Go que implementa una tabla hash: asocia claves únicas con valores. Se crea con `make(map[K]V)` o literal. No es seguro para uso concurrente. El orden de iteración con `range` es aleatorio.

### Método de valor
Función con receptor de tipo valor: `func (p Producto) Resumen() string`. El receptor es una copia del valor — modificaciones dentro del método no afectan al original. Usar cuando el método solo necesita leer el struct.

### Método de puntero
Función con receptor de tipo puntero: `func (p *Producto) Activar()`. El receptor es el puntero al original — las modificaciones sí afectan al valor original. Se verá en detalle en la Semana 07.

## N

### nil map
Un map declarado con `var m map[K]V` sin inicializar es `nil`. Leer de un nil map es válido (retorna zero value). Escribir en un nil map provoca un *pánico* en tiempo de ejecución. Siempre inicializar con `make` o literal.

### No exportado
Un identificador que empieza con minúscula (camelCase) solo es visible dentro del paquete donde fue definido. Ejemplo: `nombre`, `precio`, `id`. Ver también: *exportado*.

## P

### Patrón v, ok
Idioma de Go para verificar la existencia de una clave en un map: `v, ok := m[k]`. Si la clave existe, `ok` es `true` y `v` tiene el valor. Si no existe, `ok` es `false` y `v` es el zero value. Evita la ambigüedad con el zero value.

## R

### Receptor (receiver)
El primer argumento especial de un método que lo asocia a un tipo. Puede ser de valor `(p Producto)` o de puntero `(p *Producto)`. El nombre del receptor debe ser corto y consistente (convencionalmente la inicial del tipo).

## S

### Semántica de valor
En Go, los structs, arrays y tipos básicos son tipos de valor: al asignar o pasar como argumento, se crea una copia completa. Cada variable tiene su propia copia independiente. Contrastar con maps y slices, que tienen semántica de referencia.

### Struct
Tipo compuesto que agrupa campos con nombre y tipo fijos, conocidos en tiempo de compilación. Es un tipo de valor: se copia completo al asignar. Puede tener métodos. Comparable con `==` si todos sus campos son comparables.

## Z

### Zero value de struct
Cuando se declara `var p Producto`, todos los campos se inicializan con su zero value: `0` para numéricos, `""` para string, `false` para bool. Es siempre un estado válido y seguro en Go.
