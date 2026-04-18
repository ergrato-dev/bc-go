## 🏛️ Proyecto Semana 01 — CLI de Bienvenida

> **🎯 ÚNICO ENTREGABLE**: Este proyecto es el **único entregable obligatorio** para aprobar la semana.

### 🎯 Objetivo

Construir una pequeña herramienta de línea de comandos (CLI) que use los fundamentos vistos en la semana: `package main`, variables, funciones con múltiples retornos y manejo de errores. El programa debe adaptarse al dominio que tu instructor te asignó.

### 📋 Tu Dominio Asignado

**Dominio**: _[Tu instructor te asignará tu dominio al inicio del bootcamp]_

Ejemplos de adaptación:

| Dominio | Qué muestra el CLI |
|---------|-------------------|
| Biblioteca | Bienvenida al sistema + nombre de la biblioteca |
| Farmacia | Mensaje de bienvenida + turno del día |
| Gimnasio | Bienvenida + tipo de membresía disponible |
| Restaurante | Saludo al cliente + nombre del restaurante |

### ✅ Requisitos Funcionales

1. El programa imprime una bienvenida personalizada usando el nombre del dominio asignado
2. Acepta al menos un argumento de línea de comandos (ej: nombre del usuario) con `os.Args`
3. Si no se pasa argumento, usa un valor por defecto (zero value o constante)
4. Incluye al menos una función auxiliar con múltiples retornos
5. Verifica y maneja al menos un posible error

### 🛠️ Estructura de Entrega

```
3-proyecto/
├── README.md           # Este archivo
└── starter/
    ├── go.mod          # go mod init github.com/<usuario>/semana01
    └── main.go         # Tu implementación
```

### 💡 Código Inicial

Abre `starter/main.go`. Encontrarás la estructura base con TODOs que debes completar adaptando el código a tu dominio asignado.

### 📊 Criterios de Evaluación

| Criterio | Peso | Descripción |
|----------|------|-------------|
| Funciona con `go run .` | 30% | Sin errores de compilación ni runtime |
| Adaptación al dominio | 25% | Nombres, mensajes y contexto coherentes |
| Función con múltiples retornos | 20% | Al menos una función retorna (valor, error) |
| Manejo de error explícito | 15% | Error verificado con `if err != nil` |
| `go vet ./...` sin errores | 10% | Sin advertencias del analizador estático |
