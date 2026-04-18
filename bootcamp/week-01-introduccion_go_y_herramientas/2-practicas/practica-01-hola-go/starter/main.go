// Paquete main: qué → declara que este archivo pertenece al programa ejecutable.
// Para qué → Go requiere exactamente un 'package main' con 'func main()' por binario.
// Impacto   → cambiar a otro nombre (ej. 'package utils') convertiría esto en librería,
//
//	no en ejecutable — 'go run' fallaría.
package main

// import: qué → importa "fmt", el paquete de formateo de la librería estándar.
// Para qué → nos da Println, Printf y Sprintf sin necesidad de dependencias externas.
// Impacto   → Go no compila si hay un import no usado; mantener imports limpios es obligatorio.
import (
	"fmt"
)

// main: qué → función de entrada; el runtime la llama al ejecutar el binario.
// Para qué → aquí secuenciamos los pasos de aprendizaje de la práctica.
// Impacto   → cuando main() retorna normalmente, el proceso sale con código 0 (éxito).
func main() {
	// ============================================
	// PASO 1: El programa mínimo
	// ============================================
	// Qué     → imprimir texto en la salida estándar (stdout).
	// Para qué → verificar que el entorno Go funciona y que entendemos la estructura básica.
	// Impacto  → fmt.Println agrega automáticamente '\n' al final;
	//            fmt.Print no lo hace — la diferencia importa al componer salidas.
	// Descomenta las siguientes líneas:
	// fmt.Println("¡Hola, Go!")
	// fmt.Println()

	// ============================================
	// PASO 2: Declarar variables
	// ============================================
	// Qué     → crear variables de tipo string usando las dos sintaxis de Go.
	// Para qué → 'var' es explícita y útil en scope de paquete;
	//            ':=' infiere el tipo y solo funciona dentro de funciones.
	// Impacto  → Go es de tipado estático: una vez declarada, la variable no puede
	//            cambiar de tipo; intentarlo es un error de compilación.
	// Descomenta las siguientes líneas:
	// var lenguaje string = "Go"
	// version := "1.24"
	// fmt.Printf("Lenguaje: %s | Versión: %s\n", lenguaje, version)
	// fmt.Println()

	// ============================================
	// PASO 3: Múltiples valores de retorno
	// ============================================
	// Qué     → llamar a saludar() que devuelve dos valores a la vez.
	// Para qué → los múltiples retornos son la base del manejo de errores idiomático en Go;
	//            permiten retornar (resultado, error) sin usar excepciones.
	// Impacto  → si solo necesitas uno de los valores, usa '_' para descartar el otro:
	//            saludo, _ := saludar("Gopher")  — Go exige que todas las variables se usen.
	// Descomenta las siguientes líneas:
	// saludo, longitud := saludar("Gopher")
	// fmt.Printf("Saludo: %q (longitud: %d)\n", saludo, longitud)
	// fmt.Println()

	// ============================================
	// PASO 4: Manejo de errores
	// ============================================
	// Qué     → llamar a dividir() y verificar si retornó un error.
	// Para qué → en Go los errores son valores ordinarios (interface error), no excepciones;
	//            el llamador es responsable de decidir qué hacer con ellos.
	// Impacto  → ignorar el error con '_' compila pero es un bug; el programa puede continuar
	//            con datos inválidos. La convención es verificar err != nil inmediatamente.
	// Descomenta las siguientes líneas:
	// resultado, err := dividir(10, 2)
	// if err != nil {
	//     fmt.Println("Error:", err)
	//     return
	// }
	// fmt.Printf("10 / 2 = %.2f\n", resultado)
	//
	// _, err = dividir(5, 0)   // reutilizar 'err' con '=' porque ya fue declarada arriba
	// if err != nil {
	//     fmt.Println("División inválida:", err)
	// }
	// fmt.Println()

	// '_' descarta el valor retornado por fmt.Sprintf.
	// Para qué → mantener "fmt" en el bloque import sin que Go lo marque como no usado
	//            mientras los pasos de arriba siguen comentados.
	// Impacto   → elimina este línea una vez que descomentes cualquier paso anterior.
	_ = fmt.Sprintf("")
}

// saludar: qué → construye un saludo y retorna también la cantidad de bytes del string.
// Para qué → ilustra la firma de múltiples retornos (T1, T2) sin usar error.
// Impacto   → len() cuenta BYTES, no caracteres; para texto UTF-8 con tildes o emojis
//
//	el resultado difiere de la cantidad visible de caracteres.
//	Usar utf8.RuneCountInString() si necesitas caracteres reales.
func saludar(nombre string) (string, int) {
	saludo := "Hola, " + nombre + "!"
	// Retornar dos valores separados por coma — Go los empaqueta automáticamente.
	return saludo, len(saludo)
}

// dividir: qué → divide a entre b y retorna el resultado o un error descriptivo.
// Para qué → modela la firma idiomática (resultado, error) que verás en toda la std lib.
// Impacto   → retornar 0 como primer valor cuando hay error es convención en Go;
//
//	el llamador no debe usar ese 0 si err != nil.
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// fmt.Errorf: qué → crea un error con mensaje formateado (como Printf pero retorna error).
		// Para qué  → incluir los valores en el mensaje facilita el debugging.
		// Impacto   → en Go 1.13+ puedes usar %w en lugar de %v para «envolver» otro error.
		return 0, fmt.Errorf("no se puede dividir %v entre cero", a)
	}
	// nil: qué → el zero value de la interface error — significa «sin error».
	return a / b, nil
}
