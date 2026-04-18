package main

import (
	"fmt"
	"os"
)

// NOTA PARA EL APRENDIZ:
// Adapta este programa al dominio que tu instructor te asignó.
// Cambia los nombres, mensajes y lógica para que tengan sentido
// en el contexto de tu dominio (Biblioteca, Farmacia, Gimnasio, etc.)
//
// Ejemplos:
// - Biblioteca   → systemName = "Biblioteca Municipal", Item = "libro"
// - Farmacia     → systemName = "Farmacia Central", Item = "medicamento"
// - Gimnasio     → systemName = "GymFit", Item = "membresía"
// - Restaurante  → systemName = "Restaurante El Sabor", Item = "mesa"

// systemName es el nombre del sistema de tu dominio asignado.
// TODO: Reemplaza este valor con el nombre real de tu dominio.
const systemName = "Sistema de Gestión"

func main() {
	// Obtener el nombre del usuario desde los argumentos de línea de comandos.
	// os.Args[0] es el nombre del programa, os.Args[1] sería el primer argumento.
	// TODO: Usa os.Args para leer el nombre del usuario.
	//       Si no se pasa argumento, usa "Visitante" como valor por defecto.
	userName := "Visitante"

	// TODO: Genera el mensaje de bienvenida llamando a buildWelcome().
	//       Maneja el error correctamente con if err != nil.
	message, err := buildWelcome(systemName, userName)
	if err != nil {
		// TODO: Imprime el error de forma descriptiva y termina con os.Exit(1)
		fmt.Println("Error al generar bienvenida:", err)
		os.Exit(1)
	}

	fmt.Println(message)

	// TODO: Imprime información adicional relevante de tu dominio.
	// Ejemplo (Biblioteca): fmt.Println("Catálogo disponible: 10,000 libros")
	// Ejemplo (Farmacia):   fmt.Println("Atención: Lunes a Sábado 8:00-20:00")
	// Ejemplo (Gimnasio):   fmt.Println("Clases disponibles hoy: 12")
}

// buildWelcome construye el mensaje de bienvenida para el sistema.
// Retorna (mensaje, error) — patrón idiomático de Go.
//
// TODO: Implementa la lógica de bienvenida apropiada para tu dominio.
//       Retorna error si systemName o userName están vacíos.
func buildWelcome(system, user string) (string, error) {
	// TODO: Validar que system no sea vacío
	if system == "" {
		return "", fmt.Errorf("buildWelcome: el nombre del sistema no puede estar vacío")
	}

	// TODO: Validar que user no sea vacío
	if user == "" {
		return "", fmt.Errorf("buildWelcome: el nombre del usuario no puede estar vacío")
	}

	// TODO: Construir y retornar el mensaje de bienvenida.
	// Usa fmt.Sprintf para formatear el mensaje con los datos del dominio.
	// Ejemplo: return fmt.Sprintf("Bienvenido a %s, %s!", system, user), nil
	return "", fmt.Errorf("buildWelcome: no implementado — completa el TODO")
}
