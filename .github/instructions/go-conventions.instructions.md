---
applyTo: "**/*.go"
---

# Convenciones Go — bc-go Bootcamp

## Estilo obligatorio

- Go 1.24+: usar siempre las últimas features del lenguaje
- Declaración corta `:=` dentro de funciones; `var` solo para scope de paquete o tipo explícito
- `any` en lugar de `interface{}` (Go 1.18+)
- Indentación con **tabs** (no espacios) — `gofmt` estándar
- Longitud máxima de línea: **100 caracteres**
- Imports agrupados: stdlib primero, luego externos, separados por línea en blanco

## Nomenclatura

| Elemento               | Patrón                        | Ejemplo                    |
|------------------------|-------------------------------|----------------------------|
| Variables/funciones    | camelCase                     | `getUserByID`, `maxRetries`|
| Tipos exportados       | PascalCase                    | `UserService`, `HTTPClient`|
| Constantes exportadas  | PascalCase                    | `MaxRetries`, `DefaultPort`|
| Constantes internas    | camelCase                     | `maxRetries`               |
| Archivos               | snake_case                    | `user_handler.go`          |
| Paquetes               | lowercase una palabra         | `handlers`, `models`       |
| Acrónimos              | mayúsculas completas          | `userID`, `httpURL`        |
| Interfaces             | método + `-er` si aplica      | `Reader`, `Stringer`       |

## Manejo de errores

```go
// ✅ CORRECTO — early return, error wrapping con %w
func openFile(path string) (*os.File, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("openFile %s: %w", path, err)
    }
    return f, nil
}

// ❌ INCORRECTO — ignorar error
result, _ := strconv.Atoi("abc")

// ❌ INCORRECTO — panic en lugar de error
func mustParse(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(err) // solo válido en init() o casos verdaderamente irrecuperables
    }
    return n
}
```

## Formato de comentarios educativos

- Comentarios en **español**
- Explicar el "por qué", no el "qué"
- Godoc en inglés si el paquete es público; en español si es material educativo

```go
// getUserByID obtiene un usuario de la base de datos por su ID.
// Retorna error si el usuario no existe o si hay fallo de conexión.
func getUserByID(db *sql.DB, id int) (*User, error) {
    // Preparar query parametrizada para prevenir SQL injection
    row := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id)
    ...
}
```

## Proyectos (3-proyecto/) vs Ejercicios (2-practicas/)

- **Ejercicios**: código Go **comentado** que el alumno descomenta paso a paso
  ```go
  // Descomenta las siguientes líneas:
  // result, err := divide(10, 2)
  // fmt.Println(result, err)
  ```

- **Proyectos**: `// TODO:` para que el alumno implemente desde cero
  ```go
  // TODO: Implementar la función processItems usando lo aprendido en la semana
  func processItems(items []Item) []Item {
      return nil
  }
  ```

## Seguridad

- Nunca construir queries SQL con `fmt.Sprintf` — usar queries parametrizadas
- Validar todos los inputs del usuario en los límites del sistema
- No exponer información sensible en mensajes de error al cliente
- Sanitizar paths de archivos (path traversal)
- Usar `crypto/rand` para valores aleatorios de seguridad, nunca `math/rand`

## Concurrencia (semanas 14-17)

```go
// ✅ CORRECTO — goroutine con WaitGroup y cierre explícito
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // trabajo concurrente
}()
wg.Wait()

// ✅ CORRECTO — channel con close para señalizar fin
ch := make(chan int)
go func() {
    defer close(ch)
    for _, v := range data {
        ch <- v
    }
}()
for v := range ch {
    fmt.Println(v)
}
```

## Rendimiento

- Dimensionar slices con `make([]T, 0, cap)` cuando se conoce el tamaño aprox.
- Usar `strings.Builder` sobre concatenación en bucles
- Usar `bufio` para I/O de archivos grandes
- Pasar structs grandes por puntero para evitar copias innecesarias

## Verificación obligatoria

Todo código generado debe pasar:

```bash
go vet ./...
```

En semanas con tests (semanas 11-20):

```bash
go test ./...
```
