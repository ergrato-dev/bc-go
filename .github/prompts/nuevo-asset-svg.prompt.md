---
mode: agent
description: Genera un diagrama SVG educativo para el material de teoría del bootcamp de Go.
---

Genera un archivo SVG con un diagrama educativo para el bootcamp bc-go.

## Parámetros

- **Semana**: [COMPLETAR — ej. 14]
- **Concepto a visualizar**: [COMPLETAR — ej. Goroutine scheduler M:N]
- **Ruta destino**: `bootcamp/week-<NN>-<slug>/0-assets/<nombre>.svg`
- **Propósito**: [COMPLETAR — ej. Ilustrar cómo el runtime distribuye goroutines en OS threads]

## Tipos de diagrama más comunes por tema

| Tema                | Tipo de diagrama recomendado                    |
|---------------------|-------------------------------------------------|
| Memory model        | Stack/heap boxes con flechas de referencia       |
| Goroutines          | M:N scheduler con G/M/P boxes                   |
| Channels            | Pipeline con producer → channel → consumer       |
| Interfaces          | Struct + métodos satisfaciendo interface         |
| Punteros            | Variable → address → value (memory boxes)        |
| Slices              | Header (ptr, len, cap) + backing array           |
| HTTP handlers       | Request → Router → Handler → Response            |
| Error wrapping      | Error chain con `%w`                             |
| Context             | Context tree con cancelación en cascada          |

## Especificaciones visuales obligatorias

### Paleta de colores (dark theme)

```
Fondo principal:    #0d1117  (negro GitHub)
Fondo de elementos: #0f2535  (azul muy oscuro)
Borde/acento Go:    #00ADD8  (azul Go oficial)
Texto principal:    #e6edf3  (blanco suave)
Texto secundario:   #7d8590  (gris)
Flechas/líneas:     #00ADD8  (azul Go) o #30363d (gris oscuro)
Éxito/verde:        #3fb950
Error/rojo:         #f85149
```

### Tipografía

- Fuente: `font-family="Inter, Roboto, 'Open Sans', sans-serif"`
- Títulos: `font-size="14"` bold
- Etiquetas: `font-size="12"` normal
- Código inline: `font-family="'JetBrains Mono', 'Fira Code', monospace"`

### Estructura del archivo

```xml
<svg xmlns="http://www.w3.org/2000/svg"
     width="800" height="450" viewBox="0 0 800 450"
     role="img" aria-label="Descripción accesible del diagrama">

  <!-- Fondo -->
  <rect width="800" height="450" fill="#0d1117" rx="12"/>

  <!-- Título -->
  <text x="400" y="36"
        font-family="Inter, sans-serif" font-size="16" font-weight="bold"
        fill="#e6edf3" text-anchor="middle">
    Título del Diagrama
  </text>

  <!-- Elementos del diagrama -->
  ...

  <!-- Leyenda (si aplica) -->
  ...
</svg>
```

## Reglas obligatorias

- ❌ Sin degradés (`linearGradient`, `radialGradient`)
- ❌ Sin fuentes serif
- ✅ Usar únicamente colores sólidos de la paleta
- ✅ Incluir atributo `role="img"` y `aria-label` descriptivo para accesibilidad
- ✅ El SVG debe ser legible a 800×450px mínimo
- ✅ Todo texto en español dentro del SVG
- ✅ Usar `rx` para esquinas redondeadas en elementos (valor: 4-8)
- ✅ `viewBox` siempre presente para escalado responsivo
- ✅ Nombrar el archivo descriptivamente: `slice-internals.svg`, `goroutine-scheduler.svg`

## Verificación

Después de generar, vincula el SVG en el archivo de teoría correspondiente:

```markdown
![Descripción del diagrama](../0-assets/nombre-del-archivo.svg)
```
