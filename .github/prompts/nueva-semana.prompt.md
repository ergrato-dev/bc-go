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

Crea los archivos en este orden exacto y espera confirmación entre secciones si el contenido es extenso:

1. `bootcamp/<slug>/README.md`
2. `bootcamp/<slug>/rubrica-evaluacion.md`
3. `bootcamp/<slug>/0-assets/<concepto>.svg` — al menos 1 diagrama SVG (dark theme, color #00ADD8)
4. `bootcamp/<slug>/1-teoria/01-<concepto>.md` — mínimo 150 líneas, con SVG vinculado
5. `bootcamp/<slug>/1-teoria/02-<concepto>.md` (si aplica)
6. `bootcamp/<slug>/2-practicas/practica-01-<tema>/README.md`
7. `bootcamp/<slug>/2-practicas/practica-01-<tema>/starter/main.go`
8. `bootcamp/<slug>/2-practicas/practica-02-<tema>/README.md` (si aplica)
9. `bootcamp/<slug>/2-practicas/practica-02-<tema>/starter/main.go` (si aplica)
10. `bootcamp/<slug>/3-proyecto/README.md`
11. `bootcamp/<slug>/3-proyecto/starter/main.go`
12. `bootcamp/<slug>/4-recursos/ebooks-free/README.md`
13. `bootcamp/<slug>/4-recursos/videografia/README.md`
14. `bootcamp/<slug>/4-recursos/webgrafia/README.md`
15. `bootcamp/<slug>/5-glosario/README.md`

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

Divide la creación por secciones (teoría → prácticas → proyecto) y espera
confirmación entre secciones si el contenido es extenso.
