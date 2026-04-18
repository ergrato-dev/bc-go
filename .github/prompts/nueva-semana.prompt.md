---
mode: agent
description: Genera el contenido completo de una semana nueva del bootcamp de Go.
---

Crea el contenido completo de la siguiente semana del bootcamp bc-go siguiendo
estrictamente la estructura y convenciones del repositorio.

## Datos de la semana a crear

- **Número de semana**: [COMPLETAR — ej. 05]
- **Slug**: [COMPLETAR — ej. week-05-arrays_y_slices]
- **Tema principal**: [COMPLETAR — ej. Arrays y Slices]
- **Etapa**: [COMPLETAR — Fundamentos (1-6) | OOP en Go (7-10) | Ecosistema (11-13) | Concurrencia (14-17) | Web & APIs (18-20)]
- **Semana anterior**: [COMPLETAR — ej. week-04-funciones]
- **Semana siguiente**: [COMPLETAR — ej. week-06-maps_y_structs]

## Orden de creación obligatorio

Crea los archivos **en este orden exacto**. Espera confirmación entre secciones si el contenido es extenso.

### 1. README de la semana
`bootcamp/<slug>/README.md`

### 2. Rúbrica de evaluación
`bootcamp/<slug>/rubrica-evaluacion.md`
Consulta el formato de los repos de ejemplo (bc-sql, bc-fastapi): conocimiento 30 % / desempeño 40 % / producto 30 %.

### 3. Archivos de teoría (numerados 01-, 02-, …)
```
bootcamp/<slug>/1-teoria/01-<concepto-principal>.md   ← mínimo 150 líneas
bootcamp/<slug>/1-teoria/02-<segundo-concepto>.md     ← si aplica
```
Escribe primero el contenido completo de cada archivo de teoría.
Deja marcadores `<!-- SVG: nombre-descriptivo.svg -->` donde irá cada diagrama.

### 4. Assets SVG (numerados en orden lógico de lectura 01-, 02-, …)
```
bootcamp/<slug>/0-assets/01-<concepto-principal>.svg
bootcamp/<slug>/0-assets/02-<segundo-concepto>.svg    ← si aplica
```
Crea cada SVG y reemplaza los marcadores de la teoría por el enlace real:
`![Descripción](../0-assets/01-nombre.svg)`

### 5. Prácticas
```
bootcamp/<slug>/2-practicas/practica-01-<tema>/README.md
bootcamp/<slug>/2-practicas/practica-01-<tema>/starter/main.go
bootcamp/<slug>/2-practicas/practica-02-<tema>/README.md   ← si aplica
bootcamp/<slug>/2-practicas/practica-02-<tema>/starter/main.go
```

### 6. Proyecto
```
bootcamp/<slug>/3-proyecto/README.md
bootcamp/<slug>/3-proyecto/starter/main.go
```

### 7. Recursos (los tres completos)
```
bootcamp/<slug>/4-recursos/ebooks-free/README.md
bootcamp/<slug>/4-recursos/videografia/README.md
bootcamp/<slug>/4-recursos/webgrafia/README.md
```

### 8. Glosario
`bootcamp/<slug>/5-glosario/README.md`
Términos clave ordenados alfabéticamente con formato qué / para qué / impacto.

## Reglas críticas

### Código Go

- Go 1.24+ idiomático: `:=` dentro de funciones, `any` en lugar de `interface{}`
- Nomenclatura en inglés, comentarios educativos en español
- `go vet ./...` debe pasar sin errores en todo el código generado
- Ejercicios: código **comentado** para que el alumno descomente (`// código`)
- Proyectos: usar `// TODO:` para implementación libre del alumno
- Nunca ignorar errores con `_` salvo en `defer` justificado
- Siempre incluir `package main` e `import` necesarios

### Teoría

- **Mínimo 150 líneas** por archivo de teoría
- **Al menos 1 SVG vinculado** en cada archivo: `![Desc](../0-assets/nombre.svg)`
- 4-6 secciones numeradas con ejemplos de código Go ejecutables
- Checklist de verificación al final (4-6 items como preguntas concretas)
- Texto en español, código en inglés con comentarios en español

### Assets SVG

- Tema dark: fondo `#0d1117`, elementos `#0f2535`, bordes/acentos `#00ADD8`
- Sin degradés (gradients), fuentes sans-serif
- Diagramas educativos: memory model, goroutine flow, channel diagram, etc.
- Nombrar descriptivamente: `goroutine-scheduler.svg`, `slice-internals.svg`

### Proyectos

- El dominio de ejemplo en el proyecto NO debe ser: biblioteca, farmacia, gimnasio,
  escuela, tienda de mascotas, restaurante, banco, agencia de taxis, hospital,
  cine, hotel, agencia de viajes, concesionario, tienda de ropa, taller mecánico
- Usar dominios alternativos: museo, planetario, acuario, zoológico, etc.
- El starter/main.go debe tener `// TODO:` claros y notas para adaptar al dominio

## Distribución del tiempo (tabla en README)

| Actividad  | Tiempo |
|------------|--------|
| Teoría     | 2.5 h  |
| Prácticas  | 3 h    |
| Proyecto   | 2.5 h  |

## Navegación en README

Incluir siempre al final:

```
| Anterior | Siguiente |
|----------|-----------|
| [← Semana N-1 — Tema](../week-NN-slug/README.md) | [Semana N+1 — Tema →](../week-NN-slug/README.md) |
```

Divide la creación por secciones (según el orden de creación obligatorio) y espera
confirmación entre secciones si el contenido es extenso.
