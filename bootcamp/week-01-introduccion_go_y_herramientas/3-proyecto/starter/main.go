package main
// Paquete main: qué → punto de entrada del programa ejecutable.
// Para qué → Go requiere 'package main' + 'func main()' para producir un binario.
// Impacto   → sin esta declaración, 'go build' no genera un ejecutable.

// import: qué → lista de paquetes de la librería estándar que usaremos.
// Para qué → "fmt" provee funciones de formato/impresión;
//             "os" expone acceso al sistema operativo (args, salida con código).
// Impacto   → Go no permite imports no usados — el compilador falla si sobran.
import (
	"fmt"
	"os"
)

// NOTA PARA EL APRENDIZ:
// Adapta este programa al dominio que tu instructor te asignó.
// Cambia los nombres, mensajes y lógica para que tengan sentido
// en el contexto de tu dominio (Biblioteca, Farmacia, Gimnasio, etc.)
//
// Ejemplos de adaptación:
// - Biblioteca   → systemName = "Biblioteca Municipal", Item = "libro"
// - Farmacia     → systemName = "Farmacia Central", Item = "medicamento"
// - Gimnasio     → systemName = "GymFit", Item = "membresía"
// - Restaurante  → systemName = "Restaurante El Sabor", Item = "mesa"

// systemName: qué → constante de paquete con el nombre del sistema.
// Para qué  → las constantes son inmutables; ideales para valores fijos de configuración.
// Impacto   → si intentas reasignarla en tiempo de ejecución, el compilador lo rechaza.
// TODO: Reemplaza este valor con el nombre real de tu dominio.
const systemName = "Sistema de Gestión"

// main: qué → función de entrada; el runtime de Go la ejecuta al iniciar el programa.
// Para qué  → aquí orquestamos la lectura de argumentos, la lógica y la salida.
// Impacto   → cuando main() retorna, el proceso termina con código 0 (éxito).
func main() {
	// os.Args: qué → slice de strings con los argumentos de línea de comandos.
	// Para qué → permite que el usuario pase su nombre al ejecutar: go run . Ana
	// Impacto  → os.Args[0] siempre es el nombre del binario; los datos empiezan en [1].
	// TODO: Usa os.Args para leer el nombre del usuario.
	//       Si no se pasa argumento (len(os.Args) < 2), usa "Visitante" como default.
	userName := "Visitante"
	// buildWelcome: qué → llamada a función que construye el saludo.
	// Para qué → separar la lógica de construcción del mensaje de la función main
	//             mejora la legibilidad y facilita el testing unitario.
	// Impacto  → si retorna error, debemos manejarlo antes de usar 'message';
	//             usar 'message' sin verificar 'err' sería un bug silencioso.
	// TODO: Genera el mensaje de bienvenida llamando a buildWelcome().
	//       Maneja el error correctamente con if err != nil.
	message, err := buildWelcome(systemName, userName)
	if err != nil {
		// os.Exit(1): qué → termina el proceso con código de salida 1.
		// Para qué  → un código distinto de 0 señala error al shell/CI que invocó el programa.
		// Impacto   → los defer pendientes NO se ejecutan; usar solo para errores fatales.
		// TODO: Imprime el error de forma descriptiva y termina con os.Exit(1).
		fmt.Println("Error al generar bienvenida:", err)
		os.Exit(1)
	}

	// fmt.Println: qué → imprime el mensaje en stdout con salto de línea.
	// Para qué  → mostrar la bienvenida al usuario final.
	// Impacto   → si message está vacío el usuario ve una línea en blanco; validarlo en buildWelcome.
	fmt.Println(message)

	// TODO: Imprime información adicional relevante de tu dominio.
	// Elige la línea que corresponda y descoméntala:
	// Biblioteca: fmt.Println("Catálogo disponible: 10,000 libros")
	// Farmacia:   fmt.Println("Atención: Lunes a Sábado 8:00-20:00")
	// Gimnasio:   fmt.Println("Clases disponibles hoy: 12")
}

// buildWelcome: qué → construye el saludo de bienvenida personalizado.
// Para qué  → centralizar la lógica de formato evita repetición y facilita cambios.
// Impacto   → retornar (string, error) sigue el patrón idiomático de Go:
//
//	el llamador decide cómo manejar el fallo, no esta función.
//
// Parámetros:
//   - system: nombre del sistema (ej. "Biblioteca Municipal")
//   - user:   nombre del usuario (ej. "Ana" o "Visitante")
//
// TODO: Implementa la lógica de bienvenida apropiada para tu dominio.
//
//	Retorna error si system o user están vacíos.
func buildWelcome(system, user string) (string, error) {
	// Validación defensiva: qué → verificar que los parámetros no estén vacíos.
	// Para qué  → una función que acepta entradas inválidas produce comportamiento inesperado.
	// Impacto   → retornar error temprano (early return) mantiene el código plano y legible.
	// TODO: Validar que system no sea vacío.
	if system == "" {
		// fmt.Errorf: qué → crea un valor de tipo error con mensaje formateado.
		// Para qué  → incluir el nombre de la función en el mensaje facilita el debugging.
		// Impacto   → el llamador recibe contexto suficiente para saber qué falló y dónde.
		return "", fmt.Errorf("buildWelcome: el nombre del sistema no puede estar vacío")
	}

	// TODO: Validar que user no sea vacío.
	if user == "" {
		return "", fmt.Errorf("buildWelcome: el nombre del usuario no puede estar vacío")
	}

	// fmt.Sprintf: qué → formatea un string sin imprimirlo.
	// Para qué  → componer el mensaje con los valores dinámicos de system y user.
	// Impacto   → más eficiente y legible que concatenar strings con '+' en varias partes.
	// TODO: Construir y retornar el mensaje de bienvenida.
	// Ejemplo:   return fmt.Sprintf("Bienvenido a %s, %s!", system, user), nil
	return "", fmt.Errorf("buildWelcome: no implementado — completa el TODO")
}
