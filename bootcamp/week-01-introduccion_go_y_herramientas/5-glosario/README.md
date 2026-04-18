# Glosario — Semana 01: Introducción a Go y Herramientas

Términos clave ordenados alfabéticamente. Cada definición responde:
**qué es**, **para qué sirve** e **impacto** si no se conoce.

---

## B

### Binary / Binario
**Qué** → archivo ejecutable producido por `go build`. Contiene código máquina nativo del sistema operativo objetivo.
**Para qué** → distribuir programas Go sin que el destinatario tenga Go instalado.
**Impacto** → un binario generado en Linux no ejecuta en Windows sin recompilar; usar `GOOS=windows go build` para cross-compilation.

---

## C

### Compilación estática
**Qué** → Go incluye todas las dependencias dentro del binario por defecto.
**Para qué** → el binario es autocontenido; no necesita librerías externas en el sistema destino.
**Impacto** → los binarios Go son más grandes que los de C, pero más simples de desplegar (no hay "dependency hell").

### `const`
**Qué** → declaración de un valor inmutable en tiempo de compilación.
**Para qué** → expresar valores que no deben cambiar (PI, tamaños máximos, nombres de configuración).
**Impacto** → a diferencia de `var`, una constante no ocupa memoria en heap; el compilador la incrusta directamente.

---

## E

### Exportado (exported)
**Qué** → identificador (función, tipo, variable, constante) cuyo nombre empieza con mayúscula (PascalCase).
**Para qué** → forma parte de la API pública del paquete; otros paquetes pueden usarlo.
**Impacto** → renombrar un símbolo exportado a minúscula es un breaking change para todos los importadores.

---

## F

### `fmt`
**Qué** → paquete de la librería estándar para formateo e impresión de texto.
**Para qué** → `fmt.Println`, `fmt.Printf`, `fmt.Errorf`, `fmt.Sprintf` — las funciones más usadas en Go.
**Impacto** → `fmt.Printf` no agrega salto de línea; `fmt.Println` sí. Confundirlos es un error frecuente en principiantes.

### `func main()`
**Qué** → función de entrada obligatoria en `package main`.
**Para qué** → el runtime de Go la llama al iniciar el proceso; es donde comienza la ejecución.
**Impacto** → si `main()` retorna, el proceso termina con código 0 (éxito). Usar `os.Exit(1)` para salir con error.

---

## G

### `go build`
**Qué** → comando del toolchain que compila el código fuente y produce un binario.
**Para qué** → crear el ejecutable final para producción o distribución.
**Impacto** → sin `-o nombre`, el binario toma el nombre del directorio del módulo.

### `go fmt`
**Qué** → formateador automático de código Go.
**Para qué** → elimina discusiones de estilo; todo el código Go formateado luce igual.
**Impacto** → el código no formateado no es un error, pero los revisores lo detectarán; `gofmt` se ejecuta automáticamente al guardar en VS Code con la extensión oficial.

### `go mod init`
**Qué** → crea el archivo `go.mod` e inicializa un módulo nuevo.
**Para qué** → es el primer comando a ejecutar en cualquier proyecto Go nuevo.
**Impacto** → sin `go.mod`, el compilador no puede resolver imports locales ni descargar dependencias.

### `go mod tidy`
**Qué** → sincroniza `go.mod` y `go.sum` con los imports reales del código.
**Para qué** → agrega dependencias faltantes y elimina las que ya no se usan.
**Impacto** → ejecutar antes de cada commit garantiza que el módulo es consistente.

### `go run`
**Qué** → compila y ejecuta en un solo paso, sin producir un binario en disco.
**Para qué** → desarrollo iterativo rápido; no hay que hacer build antes de probar un cambio.
**Impacto** → más lento que ejecutar un binario precompilado porque recompila en cada llamada.

### `go vet`
**Qué** → herramienta de análisis estático que detecta errores comunes en el código.
**Para qué** → atrapar bugs antes de ejecutar: variables no usadas, formatos incorrectos en Printf, etc.
**Impacto** → `go vet` no reemplaza al compilador; detecta errores semánticos que el compilador pasa por alto.

### `go.mod`
**Qué** → archivo de configuración del módulo Go (nombre, versión de Go, dependencias).
**Para qué** → define la identidad del módulo y sus dependencias de forma reproducible.
**Impacto** → siempre versionarlo en git; sin él, otros desarrolladores no pueden reproducir el build.

### `go.sum`
**Qué** → archivo con los hashes criptográficos de cada dependencia descargada.
**Para qué** → garantizar que nadie modificó las dependencias (protección contra supply-chain attacks).
**Impacto** → nunca editar manualmente; siempre versionarlo en git junto con `go.mod`.

### Gopher
**Qué** → la mascota oficial del lenguaje Go, diseñada por Renée French.
**Para qué** → identidad visual de la comunidad; aparece en conferencias, stickers y materiales educativos.
**Impacto** → ninguno técnico, pero reconocerla te hace parte de la comunidad Go 🐹.

---

## M

### Módulo (module)
**Qué** → colección de paquetes Go con un nombre y versión comunes, definida por `go.mod`.
**Para qué** → unidad de distribución y versionado; lo que publicas en pkg.go.dev.
**Impacto** → un repositorio git = típicamente un módulo; un módulo puede tener muchos paquetes.

---

## N

### No exportado (unexported)
**Qué** → identificador cuyo nombre empieza con minúscula (camelCase); privado al paquete.
**Para qué** → encapsular detalles de implementación que no deben ser parte de la API pública.
**Impacto** → cambiar un símbolo no exportado no afecta a otros paquetes — es un "cambio seguro".

---

## O

### `os.Args`
**Qué** → slice `[]string` con los argumentos de línea de comandos. `os.Args[0]` es el binario; `[1]` en adelante son los argumentos del usuario.
**Para qué** → permitir que el usuario personalice el comportamiento del programa al ejecutarlo.
**Impacto** → acceder a `os.Args[1]` sin verificar `len(os.Args) > 1` causa un panic en tiempo de ejecución.

---

## P

### `package`
**Qué** → declaración que agrupa archivos `.go` en una unidad de compilación con nombre.
**Para qué** → organizar el código; controlar la visibilidad de los identificadores.
**Impacto** → todos los archivos `.go` en el mismo directorio deben tener el mismo nombre de paquete (excepto `_test.go`).

---

## R

### `runtime`
**Qué** → paquete de la librería estándar que expone información del entorno de ejecución de Go.
**Para qué** → obtener la versión de Go (`runtime.Version()`), el sistema operativo (`runtime.GOOS`) y la arquitectura (`runtime.GOARCH`).
**Impacto** → también controla el scheduler de goroutines (`runtime.GOMAXPROCS`) — tema de la semana 14.

---

## T

### Toolchain
**Qué** → el conjunto de herramientas CLI que vienen con la instalación de Go: `go build`, `go run`, `go test`, `go fmt`, `go vet`, `go mod`, etc.
**Para qué** → cubren el ciclo completo de desarrollo sin necesitar herramientas externas.
**Impacto** → a diferencia de otros ecosistemas, Go incluye todo lo necesario en una sola instalación.

---

## V

### `var`
**Qué** → declaración explícita de una variable con tipo opcional.
**Para qué** → usar en scope de paquete (fuera de funciones) o cuando se necesita declarar sin inicializar (zero value).
**Impacto** → dentro de funciones se prefiere `:=` (declaración corta) por brevedad; `var` es obligatorio fuera de funciones.

---

## Z

### Zero Value
**Qué** → el valor que Go asigna automáticamente a cualquier variable declarada sin inicializar.
**Para qué** → eliminar bugs de "variable no inicializada" que ocurren en C/C++.
**Impacto** → `int` → `0`, `string` → `""`, `bool` → `false`, punteros/slices/maps → `nil`. Conocer los zero values evita confusiones al imprimir variables recién declaradas.
