---
mode: agent
description: Genera un ejercicio guiado completo (starter) para una semana del bootcamp de Go.
---

Crea un ejercicio guiado paso a paso para el bootcamp bc-go.

## Datos del ejercicio

- **Semana**: [COMPLETAR — ej. 05]
- **Número de ejercicio**: [COMPLETAR — ej. 02]
- **Tema**: [COMPLETAR — ej. Slices con make y append]
- **Ruta destino**: `bootcamp/week-<NN>-<slug>/2-practicas/practica-<NN>-<tema>/`

## Archivos a generar

```
practica-NN-tema/
├── README.md
└── starter/
    └── main.go
```

> Los ejercicios NO tienen carpeta `solution/`. El alumno aprende descomentando
> el código y verificando que funcione con `go run .`.

## Formato obligatorio del README.md

Cada paso debe incluir:

1. Título con el concepto (`### Paso N: Nombre del concepto`)
2. Explicación breve del concepto (2-4 líneas)
3. Ejemplo Go ejecutable en bloque de código (con comentarios en español)
4. Instrucción: **Abre `starter/main.go`** y descomenta la sección correspondiente
5. Comando para ejecutar: `go run .`

## Formato obligatorio del starter/main.go

```go
package main

import (
    "fmt"
)

func main() {
    // ============================================
    // PASO 1: Nombre del Concepto
    // ============================================
    // Explicación educativa del concepto en español.
    // Descomenta las siguientes líneas:
    // resultado := operacion(valor)
    // fmt.Println("Resultado:", resultado)
    fmt.Println()

    // ============================================
    // PASO 2: Siguiente Concepto
    // ============================================
    // Explicación del siguiente paso.
    // Descomenta las siguientes líneas:
    // ...
    fmt.Println()
}

// funcionAuxiliar implementa la lógica del ejercicio.
// Ya está implementada — el alumno solo descomenta las llamadas en main.
func funcionAuxiliar(param string) string {
    return param
}
```

## Reglas críticas

- ❌ NUNCA usar `// TODO:` en ejercicios (eso es exclusivo para proyectos)
- ✅ SIEMPRE usar código comentado que el alumno descomenta
- Las funciones auxiliares ya están implementadas en el starter — el alumno solo
  descomenta las llamadas en `main()`
- El código comentado debe ser compilable al descomentarlo sin modificaciones
- Mínimo 3 pasos, máximo 6 pasos por ejercicio
- Cada paso debe ser independiente (puede ejecutarse aunque los otros estén comentados)
- Datos de prueba realistas, nunca `foo`, `bar`, `test1`
- Añadir `_ = fmt.Sprintf("")` al final de `main()` si hay imports no usados al tener pasos comentados

## Verificación final

Antes de entregar, verifica mentalmente que:
- [ ] El archivo `starter/main.go` compila con `go run .` (sin descomentar nada)
- [ ] Cada sección descomentada individualmente también compila
- [ ] `go vet .` pasa sin errores
- [ ] Los comentarios son claros y educativos en español
