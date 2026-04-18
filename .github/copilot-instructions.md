# 🤖 Instrucciones para GitHub Copilot

## 📋 Contexto del Bootcamp

Este es un **Bootcamp de Go Zero to Hero** estructurado para llevar a estudiantes de cero a héroe en desarrollo con el lenguaje de programación Go (Golang).

### 📊 Datos del Bootcamp

- **Duración**: 20 semanas (~5 meses)
- **Dedicación semanal**: 8 horas
- **Total de horas**: ~160 horas
- **Nivel de salida**: Desarrollador Go Junior
- **Enfoque**: Go moderno con el toolchain oficial, sin frameworks innecesarios
- **Stack**: Go 1.24+, testing nativo, net/http estándar, database/sql, SQLite/PostgreSQL, Docker

---

## 🎯 Objetivos de Aprendizaje

Al finalizar el bootcamp, los estudiantes serán capaces de:

- ✅ Dominar la sintaxis y fundamentos de Go moderno
- ✅ Aplicar el sistema de tipos estático y la inferencia de Go
- ✅ Diseñar con interfaces, composición y el enfoque OOP de Go
- ✅ Escribir código concurrente seguro con goroutines, channels y sync
- ✅ Manejar errores de forma idiomática en Go
- ✅ Organizar proyectos con módulos y paquetes Go
- ✅ Construir APIs REST con la librería estándar `net/http`
- ✅ Trabajar con bases de datos usando `database/sql` y GORM
- ✅ Escribir tests con el paquete `testing` estándar
- ✅ Desplegar aplicaciones con Docker

---

## 📚 Estructura del Bootcamp

### Distribución por Etapas

#### **Fundamentos (Semanas 1-6)** - 48 horas

- Introducción a Go: historia, ecosistema, instalación, `go` CLI, primer programa
- Variables, tipos de datos básicos, constantes, zero values e iota
- Control de flujo: `if`, `for` (única forma de bucle), `switch`, `defer`
- Funciones: múltiples retornos, valores nombrados, variadic, first-class, closures
- Arrays, Slices: internals (len/cap), operaciones, `make`, slice of slices
- Maps y Structs: definición, inicialización, campos anónimos, métodos de valor

#### **OOP en Go (Semanas 7-10)** - 32 horas

- Punteros: `&`, `*`, paso por valor vs referencia, `new()`
- Métodos e Interfaces: method sets, interfaces implícitas, duck typing
- Composición y Embedding: type assertions, type switches, `any`
- Manejo de errores: `error` interface, `fmt.Errorf`, `%w`, `errors.Is`, `errors.As`, custom errors

#### **Ecosistema Go (Semanas 11-13)** - 24 horas

- Paquetes y Módulos: `go mod`, visibilidad, `init()`, organización de proyectos
- I/O estándar, archivos y JSON: `os`, `bufio`, `io`, `encoding/json`, `Marshal/Unmarshal`
- Testing: `testing` package, table-driven tests, `testify`, benchmarks, `go test`

#### **Concurrencia (Semanas 14-17)** - 32 horas

- Goroutines y el runtime de Go: `go`, scheduler M:N, `runtime.GOMAXPROCS`
- Channels: buffered vs unbuffered, directional channels, `range` sobre channels, `close`
- `select`, `sync`: `Mutex`, `RWMutex`, `WaitGroup`, `Once`, `sync.Map`
- Patrones de concurrencia: worker pools, pipelines, fan-in/fan-out, `context`

#### **Web y APIs REST (Semanas 18-20)** - 24 horas

- `net/http`: handlers, `ServeMux`, middleware, request/response, JSON APIs
- REST APIs completas: routing con `chi`, base de datos con `database/sql` + SQLite/PostgreSQL
- Docker, CI/CD y proyecto final integrador

---

## 🗂️ Estructura de Carpetas

Cada semana sigue esta estructura estándar:

```
bootcamp/week-XX-tema_principal/
├── README.md                 # Descripción y objetivos de la semana
├── rubrica-evaluacion.md     # Criterios de evaluación detallados
├── 0-assets/                 # Imágenes, diagramas y recursos visuales
├── 1-teoria/                 # Material teórico (archivos .md)
├── 2-practicas/              # Ejercicios guiados paso a paso
├── 3-proyecto/               # Proyecto semanal integrador
│   ├── README.md             # Instrucciones del proyecto (dominio genérico)
│   ├── starter/              # Código inicial para el estudiante
│   └── solution/             # ⚠️ OCULTA - Solo para instructores (.gitignore)
├── 4-recursos/               # Recursos adicionales
│   ├── ebooks-free/
│   ├── videografia/
│   └── webgrafia/
└── 5-glosario/               # Términos clave de la semana (A-Z)
    └── README.md
```

### 📁 Carpetas Raíz

- **`assets/`**: Recursos visuales globales (logos, headers, etc.)
- **`docs/`**: Documentación general que aplica a todo el bootcamp
- **`scripts/`**: Scripts de automatización y utilidades
- **`bootcamp/`**: Contenido semanal del bootcamp

---

## 🎓 Componentes de Cada Semana

### 1. **Teoría** (1-teoria/)

- Archivos markdown con explicaciones conceptuales
- **Mínimo 150 líneas por archivo** de teoría (incluyendo código, explicaciones y ejemplos)
- Ejemplos de código Go con comentarios claros en español
- **Al menos un asset SVG vinculado** por cada archivo de teoría
- Referencias a documentación oficial de Go

### 2. **Prácticas** (2-practicas/)

- Ejercicios guiados paso a paso
- Incremento progresivo de dificultad
- Casos de uso del mundo real

#### 📋 Formato de Ejercicios

Los ejercicios son **tutoriales guiados**, NO tareas con TODOs. El estudiante aprende descomentando código:

**README.md del ejercicio:**

```markdown
### Paso 1: Crear una función con múltiples retornos

Explicación del concepto con ejemplo:

\`\`\`go
// Ejemplo explicativo
func divide(a, b float64) (float64, error) {
if b == 0 {
return 0, fmt.Errorf("división por cero")
}
return a / b, nil
}
\`\`\`

**Abre `starter/main.go`** y descomenta la sección correspondiente.
```

**starter/main.go:**

```go
// ============================================
// PASO 1: Función con múltiples retornos
// ============================================
fmt.Println("--- Paso 1: Múltiples retornos ---")

// Go permite retornar múltiples valores de una función
// Descomenta las siguientes líneas:
// result, err := divide(10, 2)
// if err != nil {
//     fmt.Println("Error:", err)
// } else {
//     fmt.Printf("Resultado: %.2f\n", result)
// }

fmt.Println()
```

> ⚠️ **IMPORTANTE**: Los ejercicios NO tienen carpeta `solution/`. El estudiante aprende descomentando el código y verificando que funcione correctamente.

#### ❌ NO usar este formato en ejercicios:

```go
// ❌ INCORRECTO - Este formato es para PROYECTOS, no ejercicios
result := 0 // TODO: Implementar
```

#### ✅ Usar este formato en ejercicios:

```go
// ✅ CORRECTO - Código comentado para descomentar
// Descomenta las siguientes líneas:
// result, err := divide(10, 2)
// fmt.Println(result, err)
```

### 3. **Proyecto** (3-proyecto/)

- Proyecto integrador que consolida lo aprendido
- README.md con instrucciones claras y **dominio genérico adaptable**
- Código inicial en `starter/`
- Carpeta `solution/` oculta (en `.gitignore`) solo para instructores
- Criterios de evaluación específicos
- **Único entregable obligatorio** de cada semana

#### 🏛️ Política de Dominios Únicos (Anticopia)

**Cada aprendiz recibe un dominio único asignado por el instructor al inicio del trimestre:**

- 📖 Biblioteca
- 💊 Farmacia
- 🏋️ Gimnasio
- 🏫 Escuela
- 🏬 Tienda de mascotas
- 🏪 Restaurante
- 🏭 Banco
- 🚕 Agencia de taxis
- 🏥 Hospital
- 🎥 Cine
- 🏞️ Hotel
- ✈️ Agencia de viajes
- 🏎️ Concesionario de autos
- 👗 Tienda de ropa
- 🛠️ Taller mecánico

**⚠️ IMPORTANTE para desarrollo de contenidos:**

- Los ejemplos en los proyectos NO deben usar dominios de la lista anterior
- Usar ejemplos genéricos o dominios diferentes (ej: Museo, Planetario, Acuario)
- Esto evita "regalar" soluciones a aprendices con esos dominios asignados

#### 📋 Formato de Proyecto (con TODOs)

A diferencia de los ejercicios, el proyecto SÍ usa TODOs para que el estudiante implemente desde cero:

**starter/main.go:**

```go
// ============================================
// FUNCIÓN: processItems
// Procesa una lista de elementos del dominio
// ============================================

// NOTA PARA EL APRENDIZ:
// Adapta este programa a tu dominio asignado.
// Ejemplos:
// - Biblioteca: procesar libros
// - Farmacia: procesar medicamentos
// - Gimnasio: procesar miembros

// Item representa un elemento del dominio asignado
type Item struct {
	ID   int
	Name string
	// TODO: Agregar campos específicos de tu dominio
	// Ejemplo (Biblioteca): Author string, Available bool
	// Ejemplo (Farmacia): Price float64, Stock int
}

// processItems aplica una operación a cada elemento
// TODO: Implementar usando lo aprendido en la semana
func processItems(items []Item) []Item {
	// 1. Iterar sobre la slice
	// 2. Aplicar transformación según tu dominio
	// 3. Retornar el resultado
	return nil
}
```

### 4. **Recursos** (4-recursos/)

- **ebooks-free/**: Libros gratuitos relevantes
- **videografia/**: Videos tutoriales complementarios
- **webgrafia/**: Enlaces a documentación y artículos

### 5. **Glosario** (5-glosario/)

- Términos técnicos ordenados alfabéticamente
- Definiciones claras y concisas
- Ejemplos de código cuando aplique

---

## 📝 Convenciones de Código

### Estilo Go Moderno (1.24+)

```go
// ✅ BIEN - declaración corta dentro de funciones
name := "Go"
count := 0

// ✅ BIEN - var para variables de paquete o tipo explícito necesario
var globalConfig Config
var timeout time.Duration = 30 * time.Second

// ✅ BIEN - constantes con iota
type Direction int
const (
    North Direction = iota
    South
    East
    West
)

// ✅ BIEN - funciones con múltiples retornos y error
func openFile(path string) (*os.File, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("openFile %s: %w", path, err)
    }
    return f, nil
}

// ✅ BIEN - manejo de error inmediato (early return)
func process(path string) error {
    f, err := openFile(path)
    if err != nil {
        return err
    }
    defer f.Close()
    // ...
    return nil
}

// ✅ BIEN - interfaces pequeñas y enfocadas
type Reader interface {
    Read(p []byte) (n int, err error)
}

// ✅ BIEN - composición en lugar de herencia
type Logger struct {
    Writer io.Writer
    prefix string
}

// ✅ BIEN - goroutine con WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // trabajo concurrente
}()
wg.Wait()

// ❌ MAL - ignorar errores
result, _ := strconv.Atoi("abc")

// ❌ MAL - panic en lugar de error
func mustParse(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(err) // ❌ solo usar panic en init o casos verdaderamente irrecuperables
    }
    return n
}
```

### Nomenclatura

- **Variables y funciones**: camelCase (`getUserByID`, `maxRetries`)
- **Tipos exportados**: PascalCase (`UserService`, `HTTPClient`)
- **Constantes**: PascalCase si exportadas, camelCase si internas
- **Archivos**: snake_case.go (`user_handler.go`, `db_config.go`)
- **Paquetes**: lowercase, una sola palabra (`handlers`, `models`, `utils`)
- **Acrónimos**: mantener en mayúsculas (`userID`, `httpURL`, `APIKey`)
- **Interfaces**: nombre del método + `-er` cuando aplique (`Reader`, `Writer`, `Stringer`)

---

## 🧪 Testing

El bootcamp enseña testing con el paquete **`testing`** estándar de Go y **testify**.

### Estructura de Tests

```go
// user_test.go
package users

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
    // Arrange
    input := UserInput{Name: "Alice", Email: "alice@example.com"}

    // Act
    user, err := CreateUser(input)

    // Assert
    require.NoError(t, err)
    assert.Equal(t, "Alice", user.Name)
    assert.NotZero(t, user.ID)
}

// Table-driven tests - patrón idiomático en Go
func TestDivide(t *testing.T) {
    tests := []struct {
        name    string
        a, b    float64
        want    float64
        wantErr bool
    }{
        {"números positivos", 10, 2, 5, false},
        {"división por cero", 10, 0, 0, true},
        {"números negativos", -10, 2, -5, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := divide(tt.a, tt.b)
            if tt.wantErr {
                assert.Error(t, err)
                return
            }
            require.NoError(t, err)
            assert.Equal(t, tt.want, got)
        })
    }
}
```

---

## 📖 Documentación

### README.md de Semana

Debe incluir:

1. **Título y descripción**
2. **🎯 Objetivos de aprendizaje**
3. **📚 Requisitos previos**
4. **🗂️ Estructura de la semana**
5. **📝 Contenidos** (con enlaces a teoría/prácticas)
6. **⏱️ Distribución del tiempo** (8 horas)
7. **📌 Entregables**
8. **🔗 Navegación** (anterior/siguiente semana)

### Archivos de Teoría

> **Requisitos mínimos**: ≥ 150 líneas por archivo · Al menos 1 SVG vinculado

```markdown
# Título del Tema

![Diagrama del concepto](../0-assets/nombre-del-concepto.svg)

## 🎯 Objetivos

- Objetivo 1
- Objetivo 2

## 📋 Contenido

### 1. Introducción

### 2. Conceptos Clave

### 3. Ejemplos Prácticos

```go
// Ejemplo de código con comentarios educativos en español
```

### 4. Ejercicios

## 📚 Recursos Adicionales

## ✅ Checklist de Verificación
```

---

## 🎨 Recursos Visuales y Estándares de Diseño

### Formato de Assets

- ✅ **Preferir SVG** para todos los diagramas, iconos y gráficos
- ❌ **NO usar ASCII art** para diagramas o visualizaciones
- ✅ Usar PNG/JPG solo para screenshots o fotografías
- ✅ Optimizar imágenes antes de incluirlas

### Criterio para Assets SVG por Semana

Los assets SVG en `0-assets/` de cada semana tienen un propósito educativo específico:

- ✅ **Apoyo visual para comprensión de conceptos teóricos**
- ✅ **Diagramas de arquitectura** (goroutines, channels, HTTP flow, etc.)
- ✅ **Visualización de procesos** (memory model, stack/heap, concurrency patterns)
- ✅ **Headers de semana** para identificación visual

**Reglas obligatorias:**

1. ❗ **Cada archivo de teoría DEBE incluir al menos un SVG vinculado** — no es opcional
2. Todo SVG debe estar **vinculado directamente en el archivo** de teoría que lo usa
3. Usar sintaxis markdown: `![Descripción](../0-assets/nombre.svg)`
4. Incluir texto alternativo descriptivo para accesibilidad
5. Nombrar archivos descriptivamente: `goroutine-scheduler.svg`, `channel-flow.svg`
6. Un SVG puede ser compartido por múltiples archivos de teoría si es relevante para ambos

### Tema Visual

- 🌙 **Tema dark** para todos los assets visuales
- ❌ **Sin degradés** (gradients) en diseños
- ✅ Colores sólidos y contrastes claros
- ✅ Paleta consistente basada en azul Go (#00ADD8) y gopher azul

### Tipografía

- ✅ **Fuentes sans-serif** exclusivamente
- ✅ Recomendadas: Inter, Roboto, Open Sans, System UI
- ❌ **NO usar fuentes serif** (Times, Georgia, etc.)
- ✅ Mantener jerarquía visual clara

---

## 🌐 Idioma y Nomenclatura

### ⚠️ REGLA CRÍTICA: Inglés Técnico + Español Educativo

**NOMENCLATURA TÉCNICA: SIEMPRE EN INGLÉS**

- ✅ Variables, constantes, funciones, tipos
- ✅ Nombres de archivos (.go)
- ✅ Campos de structs
- ✅ Nombres de paquetes

**COMENTARIOS Y DOCUMENTACIÓN: SIEMPRE EN ESPAÑOL**

- ✅ Comentarios de código (`// comentario`)
- ✅ Comentarios de godoc (`// FunctionName ...`) → excepción: godoc en inglés si se va a publicar
- ✅ READMEs y documentación educativa
- ✅ Mensajes de error para el estudiante
- ✅ Explicaciones educativas

```go
// ✅ CORRECTO - nomenclatura en inglés, comentarios en español
// getUserByID obtiene un usuario de la base de datos por su ID.
// Retorna error si el usuario no existe.
func getUserByID(db *sql.DB, id int) (*User, error) {
    // Preparar la consulta parametrizada para evitar SQL injection
    row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

    var u User
    if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
        return nil, fmt.Errorf("getUserByID %d: %w", id, err)
    }
    return &u, nil
}

// ❌ INCORRECTO - nomenclatura en español
func obtenerUsuarioPorID(db *sql.DB, id int) (*Usuario, error) {
    // ...
}
```

---

## 🔐 Mejores Prácticas

### Código Limpio

- Nombres descriptivos y significativos
- Funciones pequeñas con una sola responsabilidad
- Comentarios solo cuando sea necesario explicar el "por qué"
- Evitar anidamiento profundo con early returns
- Aceptar interfaces, retornar tipos concretos

### Seguridad

- Validar TODOS los inputs del usuario
- Usar queries parametrizadas (NUNCA formatear SQL con `fmt.Sprintf`)
- No exponer información sensible en errores al cliente
- Usar HTTPS en producción
- Sanitizar paths de archivos (path traversal)

### Rendimiento Go

- Usar `sync.Pool` para objetos de alta frecuencia de creación
- Evitar copias innecesarias de structs grandes (usar punteros)
- Preferir `strings.Builder` sobre concatenación en bucles
- Usar `bufio` para I/O de archivos
- Dimensionar slices con `make([]T, 0, cap)` cuando se conoce el tamaño

---

## 📊 Evaluación

Cada semana incluye **tres tipos de evidencias**:

1. **Conocimiento 🧠** (30%): Evaluaciones teóricas, cuestionarios
2. **Desempeño 💪** (40%): Ejercicios prácticos en clase
3. **Producto 📦** (30%): Proyecto entregable funcional

### Criterios de Aprobación

- Mínimo **70%** en cada tipo de evidencia
- Entrega puntual de proyectos
- Código funcional y bien documentado
- `go vet` y `golangci-lint` sin errores
- Tests pasando con `go test ./...` (cuando aplique)
- **Implementación coherente con el dominio asignado**
- **Originalidad**: Sin copia de implementaciones de otros aprendices

---

## 🚀 Metodología de Aprendizaje

### Estrategias Didácticas

- **Aprendizaje Basado en Proyectos (ABP)**: Proyectos semanales integradores
- **Dominios Únicos**: Cada aprendiz aplica conceptos a su dominio asignado
- **Práctica Deliberada**: Ejercicios incrementales
- **Go Challenges**: Problemas idiomáticos de Go
- **Code Review**: Revisión de código entre estudiantes
- **Live Coding**: Sesiones en vivo de programación

### Distribución del Tiempo (8h/semana)

- **Teoría**: 2-2.5 horas
- **Prácticas**: 3-3.5 horas
- **Proyecto**: 2-2.5 horas

---

## 🤖 Instrucciones para Copilot

Cuando trabajes en este proyecto:

### Límites de Respuesta

1. **Divide respuestas largas**
   - ❌ **NUNCA generar respuestas que superen los límites de tokens**
   - ✅ **SIEMPRE dividir contenido extenso en múltiples entregas**
   - ✅ Crear contenido por secciones, esperar confirmación del usuario
   - ✅ Priorizar calidad sobre cantidad en cada entrega
   - Razón: Evitar pérdida de contenido y garantizar completitud

2. **Estrategia de División**
   - Para semanas completas: dividir por carpetas (teoria → practicas → proyecto)
   - Para archivos grandes: dividir por secciones lógicas
   - Siempre indicar claramente qué parte se entrega y qué falta
   - Esperar confirmación del usuario antes de continuar

### Generación de Código

1. **Usa siempre Go moderno (1.24+)**
   - Declaración corta `:=` dentro de funciones
   - Generic types cuando simplifiquen el código (`[T any]`)
   - `any` en lugar de `interface{}`
   - `errors.Is` / `errors.As` para inspección de errores
   - `fmt.Errorf` con `%w` para wrapping
   - `context.Context` como primer parámetro en funciones I/O
   - Structured logging con `log/slog` (Go 1.21+)

2. **Entorno de Desarrollo**
   - ✅ Go toolchain oficial (sin Docker obligatorio para aprendizaje inicial)
   - ✅ Docker para semanas de APIs y deployment (semanas 18-20)
   - ✅ `go mod` para gestión de dependencias
   - ✅ `golangci-lint` para linting
   - ✅ `.air` para hot-reload en desarrollo (semanas web)

3. **Gestión de Dependencias**
   - ❌ **NUNCA agregar dependencias innecesarias** — Go std lib es suficiente para la mayoría
   - ✅ Preferir librería estándar sobre frameworks
   - ✅ Para routing: `chi` (minimalista, compatible con `net/http`)
   - ✅ Para testing: `testify` (assert/require)
   - ✅ Para DB: `database/sql` + driver `mattn/go-sqlite3` o `jackc/pgx`
   - ✅ SIEMPRE fijar versiones exactas en `go.mod` (sin rangos)

4. **Comenta el código de manera educativa**
   - Explica conceptos para principiantes
   - Incluye referencias a documentación cuando sea útil
   - Usa comentarios que enseñen, no solo describan
   - Menciona el "por qué" de las decisiones idiomáticas de Go

5. **Proporciona ejemplos completos y funcionales**
   - Código que se pueda copiar y ejecutar con `go run`
   - Incluir `package main` y `import` cuando sea necesario
   - Mostrar tanto lo que se debe hacer como lo que se debe evitar

### Creación de Contenido

1. **Estructura clara y progresiva**
   - De lo simple a lo complejo
   - Conceptos construidos sobre conocimientos previos
   - Repetición espaciada de conceptos clave de Go

2. **Ejemplos del mundo real**
   - Casos de uso prácticos y relevantes
   - Proyectos que los estudiantes puedan mostrar en portfolios
   - Problemas que encontrarán en el desarrollo real con Go

3. **Enfoque idiomático**
   - Siempre preferir el estilo idiomático de Go ("The Go Way")
   - No importar patrones de otros lenguajes directamente
   - Referenciar [Effective Go](https://go.dev/doc/effective_go) y [Go Proverbs](https://go-proverbs.github.io/)

### Respuestas y Ayuda

1. **Explicaciones claras**
   - Lenguaje simple y directo
   - Evitar jerga innecesaria
   - Proporcionar analogías cuando sea útil (especialmente para concurrencia)

2. **Código comentado**
   - Explicar cada paso importante
   - Destacar conceptos clave de Go
   - Señalar errores comunes de principiantes Go (null vs zero value, goroutine leaks, etc.)

3. **Recursos adicionales**
   - Referencias a [pkg.go.dev](https://pkg.go.dev)
   - Referencias a [go.dev/tour](https://go.dev/tour)
   - Referencias a [gobyexample.com](https://gobyexample.com)

---

## 📚 Referencias Oficiales

- **Go Documentation**: https://go.dev/doc/
- **Go Standard Library**: https://pkg.go.dev/std
- **A Tour of Go**: https://go.dev/tour
- **Effective Go**: https://go.dev/doc/effective_go
- **Go Proverbs**: https://go-proverbs.github.io/
- **Go by Example**: https://gobyexample.com/
- **Go Blog**: https://go.dev/blog/

---

## 🔗 Enlaces Importantes

- **Repositorio**: https://github.com/ergrato-dev/bc-go
- **Documentación general**: [docs/README.md](docs/README.md)
- **Primera semana**: [bootcamp/week-01-introduccion_go_y_herramientas/README.md](bootcamp/week-01-introduccion_go_y_herramientas/README.md)

---

## ✅ Checklist para Nuevas Semanas

Cuando crees contenido para una nueva semana:

- [ ] Crear estructura de carpetas completa
- [ ] README.md con objetivos y estructura
- [ ] Material teórico en 1-teoria/
- [ ] Ejercicios guiados (código comentado) en 2-practicas/
- [ ] Proyecto integrador (TODOs) en 3-proyecto/
- [ ] Recursos adicionales en 4-recursos/
- [ ] Glosario de términos en 5-glosario/
- [ ] Rúbrica de evaluación (rubrica-evaluacion.md)
- [ ] Verificar coherencia con semanas anteriores
- [ ] Revisar progresión de dificultad
- [ ] Probar código de ejemplos con `go run` / `go test`
- [ ] `go vet ./...` sin errores en el código de ejemplo

---

## 💡 Notas Finales

- **Prioridad**: Claridad sobre brevedad
- **Enfoque**: Aprendizaje práctico sobre teoría abstracta
- **Objetivo**: Preparar desarrolladores Go listos para trabajar
- **Filosofía**: Go idiomático desde el día 1 — "Clear is better than clever"

---

_Última actualización: Abril 2026_
_Versión: 1.0_
