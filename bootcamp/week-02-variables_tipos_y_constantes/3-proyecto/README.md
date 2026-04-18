## 🏛️ Proyecto Semana 02 — Catálogo con Tipos y Constantes

> **🎯 ÚNICO ENTREGABLE**: Este proyecto es el **único entregable obligatorio** para aprobar la semana.

### 🎯 Objetivo

Construir un catálogo de inventario en línea de comandos que use el sistema de tipos de Go, constantes con `iota` para categorías y tipos personalizados. El programa debe adaptarse al dominio que tu instructor te asignó.

### 📋 Tu Dominio Asignado

**Dominio**: _[Tu instructor te asignará tu dominio al inicio del bootcamp]_

Ejemplos de adaptación:

| Dominio | Categorías con iota | Constantes fijas |
|---------|---------------------|------------------|
| Biblioteca | `Fiction`, `NonFiction`, `Reference` | Nombre de la biblioteca, capacidad |
| Farmacia | `OTC`, `Prescription`, `Controlled` | Nombre de la farmacia, dirección |
| Gimnasio | `Cardio`, `Strength`, `Flexibility` | Nombre del gym, capacidad máxima |
| Restaurante | `Starter`, `Main`, `Dessert`, `Beverage` | Nombre del restaurante, aforo |

### ✅ Requisitos Funcionales

1. Define al menos **una constante de texto** con el nombre del dominio asignado
2. Define al menos **una constante numérica** (capacidad, año, precio máximo, etc.)
3. Crea un **tipo personalizado** con `iota` para categorizar elementos del dominio (mínimo 3 categorías)
4. Declara variables con `var` y `:=` para representar al menos un elemento del dominio
5. Realiza al menos **una conversión de tipo explícita** necesaria para el dominio
6. Imprime información con `fmt.Printf` usando al menos 3 verbos distintos (`%d`, `%s` o `%q`, `%T` o `%v`)

### 🛠️ Estructura de Entrega

```
3-proyecto/
└── starter/
    └── main.go   ← el único archivo que debes modificar
```

### 🚀 Ejecución

```bash
cd 3-proyecto/starter
go run .
```

### 📋 Criterios de Evaluación

Ver [rubrica-evaluacion.md](../rubrica-evaluacion.md) — sección Producto.

### 💡 Notas para el Aprendiz

- Los TODOs en `starter/main.go` son guías, no restricciones. Puedes agregar más funciones, tipos y constantes si tu dominio lo requiere.
- No uses ninguno de los dominios de la lista restringida como nombre o contexto del programa.
- Ejecuta `go vet .` antes de entregar y corrige cualquier advertencia.
