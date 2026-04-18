# 🏛️ Proyecto Semana 03 — Procesador de Registros con Control de Flujo

> **🎯 ÚNICO ENTREGABLE**: Este proyecto es el **único entregable obligatorio** para aprobar la semana.

## 🎯 Objetivo

Construir un procesador de registros en línea de comandos que clasifique, filtre y reporte elementos del dominio asignado. El programa debe usar `for range`, `switch`, `defer` e `if` con inicialización de forma integrada y coherente.

## 📋 Tu Dominio Asignado

**Dominio**: _[Tu instructor te asignará tu dominio al inicio del bootcamp]_

Ejemplos de adaptación:

| Dominio | Elementos | Categorías con switch | defer |
|---------|-----------|----------------------|-------|
| Biblioteca | Libros | Género: Ficción/NoFicción/Referencia | "Cerrando catálogo..." |
| Farmacia | Medicamentos | Tipo: OTC/Receta/Controlado | "Cerrando sistema de inventario..." |
| Gimnasio | Equipos | Zona: Cardio/Fuerza/Flexibilidad | "Cerrando registro de uso..." |
| Restaurante | Platos | Tipo: Entrada/Principal/Postre/Bebida | "Cerrando caja..." |

## ✅ Requisitos Funcionales

1. Define una **slice** de al menos 5 elementos de tu dominio (structs o valores simples)
2. Usa `for range` para iterar sobre todos los elementos
3. Usa `switch` para clasificar cada elemento (mínimo 3 casos + default)
4. Usa `if` con **inicialización de variable** al menos una vez
5. Usa `defer` de forma idiomática al inicio del `main` para imprimir un mensaje de cierre
6. Imprime un **resumen estadístico** al final (total de elementos, conteo por categoría)

## 🛠️ Estructura de Entrega

```
3-proyecto/
└── starter/
    └── main.go   ← el único archivo que debes modificar
```

## 🚀 Ejecución

```bash
cd 3-proyecto/starter
go run .
```

## 📋 Criterios de Evaluación

Ver [rubrica-evaluacion.md](../rubrica-evaluacion.md) — sección Producto.

## 💡 Notas para el Aprendiz

- El `defer` de cierre debe declararse al inicio del `main`, inmediatamente después de cualquier inicialización de "apertura".
- Usa `switch` sin expresión si necesitas comparar rangos o condiciones complejas.
- No uses dominios de la lista restringida (biblioteca, farmacia, gimnasio, etc.) como ejemplo en tu código.
- Ejecuta `go vet .` antes de entregar.
