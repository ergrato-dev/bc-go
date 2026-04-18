// Paquete main: qué  → punto de entrada del programa ejecutable.
// Para qué → junto con func main(), indica al compilador que este directorio
//             produce un binario, no una librería.
// Impacto   → sin 'package main', go build no genera ningún ejecutable.
package main

// import: qué  → importa tres paquetes de la librería estándar.
// Para qué → "fmt" para imprimir; "os" para leer argumentos del S.O.;
//             "runtime" para metadatos del entorno de ejecución de Go.
// Impacto   → Go falla en compilación si hay imports no usados — mantener
//             solo los necesarios es una regla, no una convención opcional.
import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// ============================================
	// PASO 1: Verificar que el módulo funciona
	// ============================================
	// Qué     → imprimir un mensaje de inicio para confirmar que el entorno
	//           está configurado correctamente.
	// Para qué → antes de agregar lógica, validamos que 'go mod init' se
	//            ejecutó y que 'go run .' funciona sin errores.
	// Impacto  → si falta go.mod, el compilador no puede resolver imports;
	//            siempre inicializar el módulo antes de ejecutar.
	// Descomenta las siguientes líneas:
	// fmt.Println("=== Práctica 02: CLI con os.Args ===")
	// fmt.Println()

	// ============================================
	// PASO 2: Leer argumentos con os.Args
	// ============================================
	// Qué     → os.Args es un []string con los argumentos pasados al programa.
	// Para qué → permite personalizar el comportamiento del CLI sin recompilar;
	//            los argumentos son la forma más simple de input en herramientas CLI.
	// Impacto  → os.Args[0] SIEMPRE es el nombre del binario; el primer dato
	//            del usuario está en os.Args[1]. Acceder a [1] sin verificar
	//            len(os.Args) causa un panic en tiempo de ejecución (index out of range).
	// Descomenta las siguientes líneas:
	// name := "Visitante" // valor por defecto si no se pasa argumento
	// if len(os.Args) > 1 {
	//     // os.Args[1] es el primer argumento real del usuario
	//     name = os.Args[1]
	// }
	// fmt.Printf("Argumento recibido: %q\n", name)
	// fmt.Println()

	// ============================================
	// PASO 3: Extraer lógica a una función
	// ============================================
	// Qué     → delegar la construcción del saludo a formatGreeting().
	// Para qué → separar la lógica de presentación de la orquestación en main()
	//            hace el código más legible, testeable y reutilizable.
	// Impacto  → una función con una sola responsabilidad es más fácil de
	//            probar con go test que un bloque largo dentro de main().
	// Descomenta las siguientes líneas:
	// name := "Visitante"
	// if len(os.Args) > 1 {
	//     name = os.Args[1]
	// }
	// saludo := formatGreeting(name)
	// fmt.Println(saludo)
	// fmt.Println()

	// ============================================
	// PASO 4: Información del sistema con runtime
	// ============================================
	// Qué     → el paquete 'runtime' expone metadatos del entorno de ejecución.
	// Para qué → útil para mostrar versión de Go en herramientas de diagnóstico,
	//            logs de inicio, o comandos 'version' de aplicaciones CLI.
	// Impacto  → runtime.GOOS y runtime.GOARCH son constantes en tiempo de
	//            compilación; cambiarán el valor si compilas para otro sistema
	//            con GOOS=windows go build.
	// Descomenta las siguientes líneas:
	// name := "Visitante"
	// if len(os.Args) > 1 {
	//     name = os.Args[1]
	// }
	// fmt.Println(formatGreeting(name))
	// fmt.Printf("Go versión : %s\n", runtime.Version())
	// fmt.Printf("Sistema    : %s/%s\n", runtime.GOOS, runtime.GOARCH)
	// fmt.Println()

	// ============================================
	// PASO 5: Compilar un binario con go build
	// ============================================
	// No hay código que descomentar aquí.
	// Ejecuta estos comandos en la terminal:
	//
	//   go build -o saluda .
	//   ./saluda TuNombre
	//   ./saluda               ← debe usar "Visitante"
	//
	// Qué     → go build compila el código y produce un binario nativo.
	// Para qué → el binario no requiere que Go esté instalado para ejecutarse;
	//            es la forma de distribuir programas Go en producción.
	// Impacto  → go run recompila en cada ejecución (útil en desarrollo);
	//            go build produce un artefacto final reutilizable.

	// Evitar "declared and not used" mientras los pasos están comentados.
	// Qué     → el identificador '_' descarta el valor; '= fmt.Sprintf("")' crea
	//           un string vacío para mantener el import de fmt activo.
	// Para qué → Go no compila si hay un import sin usar; este truco mantiene
	//           el código compilable mientras los pasos están comentados.
	// Impacto  → eliminar esta línea una vez que descomentes cualquier paso.
	_, _ = fmt.Sprintf(""), os.Args[0]
	_ = runtime.Version()
}

// formatGreeting: qué  → construye el mensaje de saludo con el nombre recibido.
// Para qué → centraliza el formato del saludo en un único lugar;
//             si el formato cambia, solo se modifica aquí.
// Impacto   → al ser una función pura (sin side effects), es trivialmente
//             testeable con go test sin necesidad de mocks.
func formatGreeting(name string) string {
	// fmt.Sprintf: qué  → formatea un string sin imprimirlo (a diferencia de Printf).
	// Para qué → construir el mensaje antes de decidir cómo entregarlo
	//            (imprimir, escribir a archivo, retornar, etc.).
	// Impacto  → más eficiente y legible que concatenar con '+' cuando
	//            hay más de un valor dinámico en el string.
	return fmt.Sprintf("¡Hola, %s! Bienvenido al bootcamp de Go.", name)
}
