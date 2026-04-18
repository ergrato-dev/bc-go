---
applyTo: "**/1-teoria/*.md"
---

# Convenciones de Teoría — bc-go Bootcamp

## Límites de extensión (NON-NEGOTIABLE)

| Elemento                       | Límite                               |
|--------------------------------|--------------------------------------|
| Líneas por archivo             | **Mínimo 150**                       |
| Objetivos                      | 3-5 ítems con verbo de acción        |
| Secciones numeradas            | 4-6 (`## 1.`, `## 2.`…)             |
| Ejemplos de código por sección | Máximo 2                             |
| Checklist final                | 4-6 ítems como preguntas concretas   |
| Referencias                    | 2-4 links a documentación oficial    |

## Estructura obligatoria de cada archivo de teoría

```markdown
# Título del Concepto

![Descripción del diagrama](../0-assets/nombre-del-concepto.svg)

## 🎯 Objetivos

- Verbo de acción + concepto 1
- Verbo de acción + concepto 2
- Verbo de acción + concepto 3

## 1. Primera sección

Explicación conceptual...

```go
// Ejemplo comentado en español
func ejemplo() {
    // código ilustrativo
}
```

## 2. Segunda sección

...

## ✅ Checklist de verificación

- [ ] ¿Pregunta concreta 1?
- [ ] ¿Pregunta concreta 2?
- [ ] ¿Pregunta concreta 3?
- [ ] ¿Pregunta concreta 4?

## 📚 Recursos adicionales

- [Effective Go — Sección relevante](https://go.dev/doc/effective_go#section)
- [pkg.go.dev — paquete relevante](https://pkg.go.dev/package)
```

## SVG obligatorio por archivo

**Cada archivo de teoría DEBE incluir al menos un SVG vinculado** — no es opcional.

```markdown
![Descripción accesible del diagrama](../0-assets/nombre-del-concepto.svg)
```

- El SVG debe estar en `bootcamp/week-XX-slug/0-assets/`
- Usar nombre descriptivo: `goroutine-scheduler.svg`, `slice-internals.svg`
- Un SVG puede ser compartido entre varios archivos de teoría si aplica

## Qué NO incluir

- ❌ Tablas de comparación de más de 5 filas
- ❌ Secciones de "Herramientas recomendadas" (van en `4-recursos/`)
- ❌ Notas de compatibilidad extensas (una línea `>` es suficiente)
- ❌ Más de 2 ejemplos de código por sección
- ❌ ASCII art para diagramas (usar SVGs)
- ❌ Comandos de instalación (van en el README de la semana)

## Idioma

- Texto explicativo: **español**
- Nomenclatura de código Go (variables, funciones, tipos): **inglés**
- Comentarios dentro del código Go: **español**
- Títulos y headings: **español**

## Numeración de archivos

Prefijo numérico de dos dígitos:

```
01-introduccion.md
02-tipos-basicos.md
03-control-flujo.md
```

## Referencias obligatorias

Siempre incluir al menos una referencia a:

- [Effective Go](https://go.dev/doc/effective_go)
- [pkg.go.dev](https://pkg.go.dev) — paquete relevante de la semana
- [A Tour of Go](https://go.dev/tour) — sección correspondiente (si aplica)
- [Go by Example](https://gobyexample.com/) — ejemplo correspondiente (si aplica)
