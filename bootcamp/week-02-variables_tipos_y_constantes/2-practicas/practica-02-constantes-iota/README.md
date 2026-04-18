# Práctica 02 — Constantes e iota en Go

## 🎯 Objetivo

Declarar constantes simples, agruparlas en bloques `const`, usar `iota` para crear enumeraciones idiomáticas y aplicar el patrón de flags de bits con desplazamiento.

---

## 📋 Pasos

### Paso 1: Constantes simples y bloques `const`

Las constantes se fijan en tiempo de compilación y nunca pueden modificarse.
Un bloque `const` agrupa constantes relacionadas para mayor legibilidad.

```go
// Ejemplo: bloque const para configuración de una aplicación
const (
    appVersion   = "2.0.0"
    maxRetries   = 3
    defaultPort  = 8080
    timeoutSecs  = 30
)

fmt.Printf("App: %s | Puerto: %d | Timeout: %ds\n",
    appVersion, defaultPort, timeoutSecs)
```

**Abre `starter/main.go`** y descomenta la sección **PASO 1**.
Ejecuta con: `go run .`

---

### Paso 2: Constantes tipadas vs no tipadas

Una constante tipada solo es compatible con su tipo exacto.
Una constante no tipada es flexible y se adapta al tipo donde se usa.

```go
// Ejemplo: constante no tipada — compatible con int32 y int64
const maxItems = 100 // no tipada (tipo "ideal" int)

var a int32 = maxItems // ✅ compatible
var b int64 = maxItems // ✅ compatible

// Constante tipada — requiere conversión explícita para otro tipo
const typedMax int32 = 100

var c int64 = int64(typedMax) // ✅ requiere conversión

fmt.Println(a, b, c)
```

**Abre `starter/main.go`** y descomenta la sección **PASO 2**.
Ejecuta con: `go run .`

---

### Paso 3: `iota` — enumeraciones con tipo personalizado

`iota` empieza en `0` y se incrementa en `1` por cada constante del bloque.
Crear un tipo personalizado (`type X int`) agrega seguridad de tipos al enum.

```go
// Ejemplo: estaciones del año con iota
type Season int

const (
    Spring Season = iota // 0
    Summer               // 1
    Autumn               // 2
    Winter               // 3
)

current := Autumn
fmt.Printf("Estación actual: %d\n", current) // 2
fmt.Printf("Tipo: %T\n", current)            // main.Season
```

**Abre `starter/main.go`** y descomenta la sección **PASO 3**.
Ejecuta con: `go run .`

---

### Paso 4: `iota` con desplazamiento de bits — flags de permisos

`1 << iota` genera potencias de 2: cada constante tiene un único bit activo.
Esto permite combinar permisos con `|` y verificarlos con `&`.

```go
// Ejemplo: flags de permisos de acceso
type AccessFlag uint8

const (
    CanRead    AccessFlag = 1 << iota // 0b0001 = 1
    CanWrite                          // 0b0010 = 2
    CanDelete                         // 0b0100 = 4
    CanPublish                        // 0b1000 = 8
)

// Editor: puede leer, escribir y publicar (pero no borrar)
editor := CanRead | CanWrite | CanPublish

fmt.Printf("Editor flags: %04b\n", editor) // 1011
fmt.Printf("¿Puede leer?:    %v\n", editor&CanRead != 0)    // true
fmt.Printf("¿Puede borrar?:  %v\n", editor&CanDelete != 0)  // false
```

**Abre `starter/main.go`** y descomenta la sección **PASO 4**.
Ejecuta con: `go run .`

---

## ✅ Verificación

Cuando termines, la salida debe mostrar:
- Los valores de las constantes de configuración con sus tipos
- La diferencia entre constante tipada y no tipada al asignar a `int32` e `int64`
- El valor numérico de cada estación (0-3) con su tipo `main.Season`
- Los flags de permisos en binario y la verificación con `&`
