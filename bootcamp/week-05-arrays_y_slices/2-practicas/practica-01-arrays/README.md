# Práctica 01 — Arrays en Go

## Objetivo

Explorar los arrays de tamaño fijo: declaración, iteración con `range`, semántica de valor (copia) y comparación con `==`.

## Instrucciones

Abre `starter/main.go` y descomenta las secciones de cada paso en orden. Verifica que el programa compile y muestre los resultados esperados antes de pasar al siguiente paso.

---

### Paso 1: Declaración e inicialización

Go ofrece tres formas de inicializar un array. El tamaño forma parte del tipo.

```go
// Con zero value
var puntuaciones [5]int

// Con literal
colores := [3]string{"rojo", "verde", "azul"}

// Tamaño inferido con ... (el compilador cuenta los elementos)
planetas := [...]string{"Mercurio", "Venus", "Tierra", "Marte"}
```

**Abre `starter/main.go`** y descomenta la sección **PASO 1**.

Resultado esperado:
```
--- Paso 1: Declaración ---
Puntuaciones iniciales: [0 0 0 0 0]
Puntuaciones: [95 87 92 78 88]
Colores: [rojo verde azul]
Planetas (4): [Mercurio Venus Tierra Marte]
```

---

### Paso 2: range y estadísticas

`range` sobre un array devuelve `(índice, copia del valor)` en cada iteración.

```go
temps := [7]float64{22.5, 24.0, 19.3, 21.8, 25.1, 23.4, 20.7}

var suma float64
for i, temp := range temps {
    fmt.Printf("Día %d: %.1f°C\n", i+1, temp)
    suma += temp
}
fmt.Printf("Promedio: %.2f°C\n", suma/float64(len(temps)))
```

**Descomenta la sección PASO 2** y verifica que el promedio es `22.51°C`.

---

### Paso 3: Semántica de valor — copia al asignar y pasar

Asignar un array crea una **copia independiente**. La misma regla aplica al pasar a funciones.

```go
original := [3]int{10, 20, 30}
copia := original
copia[0] = 99
// original no se modifica
```

Para modificar el original desde una función, se pasa un **puntero**:

```go
func triplicar(arr *[3]int) {
    for i := range arr {
        arr[i] *= 3
    }
}
```

**Descomenta la sección PASO 3** y la función `triplicar` al final del archivo.

---

### Paso 4: Comparación con == y arrays como claves de mapa

Los arrays son **comparables** con `==` si sus elementos lo son. Esto permite usarlos como claves de mapa, algo imposible con slices.

```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
fmt.Println(a == b) // true

type Coordenada [2]float64
ciudades := map[Coordenada]string{
    {40.4168, -3.7038}: "Madrid",
}
```

**Descomenta la sección PASO 4** y verifica la ciudad encontrada por coordenadas.
