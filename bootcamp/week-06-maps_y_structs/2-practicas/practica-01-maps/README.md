# Práctica 01 — Maps: CRUD y Patrones Idiomáticos

Aprenderás a declarar maps, leer con el patrón `v, ok`, iterar, eliminar entradas y usar los patrones idiomáticos más comunes en Go.

## Prerequisitos

- Semana 05 completada (slices y range)
- Teoría `1-teoria/01-maps.md` leída

## Instrucciones generales

Abre `starter/main.go`. El archivo tiene cuatro pasos. Descomenta las líneas de cada paso uno a uno, ejecuta con `go run starter/main.go` y verifica la salida esperada antes de pasar al siguiente.

---

### Paso 1 — Declarar y escribir en un map

Conceptos: `make`, asignación de entradas, lectura básica.

```go
// Un map asocia claves únicas con valores.
// make(map[K]V) crea el map listo para usar.
// Leer una clave inexistente retorna el zero value del tipo del valor.

stock := make(map[string]int)
stock["lapiz"]   = 50
stock["cuaderno"] = 20
stock["goma"]     = 0  // asignado explícitamente como 0

fmt.Println(stock["lapiz"])          // 50
fmt.Println(stock["goma"])           // 0
fmt.Println(stock["regla"])          // 0 — clave inexistente, mismo resultado que "goma"
```

**Abre `starter/main.go`** y descomenta la sección `PASO 1`.

Salida esperada:
```
--- Paso 1: Declarar y escribir ---
Stock lapiz:    50
Stock goma:     0
Stock regla:    0
```

---

### Paso 2 — Patrón `v, ok` para verificar existencia

Conceptos: segunda variable de retorno en acceso a map, diferencia entre "no existe" y "zero value".

```go
// El problema: stock["goma"] y stock["regla"] retornan ambos 0.
// ¿Cómo saber si "regla" existe o no?
// Solución: patrón v, ok

cantidad, ok := stock["regla"]
if !ok {
    fmt.Println("regla no está registrada")
} else {
    fmt.Printf("regla tiene %d unidades\n", cantidad)
}

// Verificación en el if — idiomático
if cant, ok := stock["goma"]; ok {
    fmt.Printf("goma encontrada: %d unidades\n", cant)
}
```

**Abre `starter/main.go`** y descomenta la sección `PASO 2`.

Salida esperada:
```
--- Paso 2: Patrón v, ok ---
regla no está registrada
goma encontrada: 0 unidades
```

---

### Paso 3 — Iterar y contar frecuencias

Conceptos: `range` sobre map, orden no garantizado, patrón contador.

```go
// range map retorna (clave, valor) en orden aleatorio.
// El patrón freq[k]++ funciona porque el zero value de int es 0.

palabras := []string{"go", "rust", "go", "python", "go", "rust"}
freq := make(map[string]int)
for _, p := range palabras {
    freq[p]++ // si no existe, parte de 0
}

for lenguaje, veces := range freq {
    fmt.Printf("%s: %d veces\n", lenguaje, veces)
}
// go: 3, rust: 2, python: 1 — pero el orden varía cada ejecución
```

**Abre `starter/main.go`** y descomenta la sección `PASO 3`.

Salida esperada (el orden puede variar):
```
--- Paso 3: Frecuencias ---
go: 3 veces
rust: 2 veces
python: 1 veces
```

---

### Paso 4 — delete y map como conjunto (set)

Conceptos: `delete`, `map[T]bool` como set, verificación de pertenencia.

```go
// delete elimina una entrada del map.
// delete sobre clave inexistente no hace nada (no produce error).

delete(stock, "goma")
if _, ok := stock["goma"]; !ok {
    fmt.Println("goma eliminada correctamente")
}

// Map como conjunto — simula un set con map[string]bool
procesados := make(map[string]bool)
procesados["orden-001"] = true
procesados["orden-002"] = true

if procesados["orden-001"] {
    fmt.Println("orden-001 ya fue procesada")
}
if !procesados["orden-003"] {
    fmt.Println("orden-003 pendiente de procesar")
}
```

**Abre `starter/main.go`** y descomenta la sección `PASO 4`.

Salida esperada:
```
--- Paso 4: delete y set ---
goma eliminada correctamente
orden-001 ya fue procesada
orden-003 pendiente de procesar
```

---

## Verificación final

```bash
go run starter/main.go
```

Todos los pasos deben imprimir sus resultados en orden sin errores.
