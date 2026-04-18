<!-- assets/bootcamp-header.svg -->

[![CC BY-NC-SA 4.0](https://img.shields.io/badge/License-CC%20BY--NC--SA%204.0-lightgrey.svg)](LICENSE)
![Semanas](https://img.shields.io/badge/20%20Semanas-5%20meses-00ADD8)
![Horas](https://img.shields.io/badge/160%20Horas-8h%2Fsemana-00ADD8)
![Go](https://img.shields.io/badge/Go-1.24%2B-00ADD8?logo=go)

[English Version](README_EN.md)

---

# 🐹 Bootcamp Go Zero to Hero

## 📋 Descripción

Bootcamp intensivo de 20 semanas (~5 meses) enfocado en el dominio del lenguaje Go moderno. Diseñado para llevar a estudiantes de cero a Desarrollador Go Junior, con énfasis en código idiomático, herramientas oficiales del ecosistema y proyectos del mundo real.

> **Go idiomático desde el día 1** — Sin código legacy, solo las mejores prácticas del lenguaje.

### 🎯 Objetivos

Al finalizar el bootcamp, los estudiantes serán capaces de:

- ✅ Dominar la sintaxis y fundamentos de Go moderno (1.24+)
- ✅ Aplicar el sistema de tipos estático y la inferencia de tipos de Go
- ✅ Diseñar con interfaces, composición y el enfoque OOP de Go
- ✅ Escribir código concurrente seguro con goroutines, channels y sync
- ✅ Manejar errores de forma idiomática en Go
- ✅ Organizar proyectos con módulos y paquetes Go
- ✅ Construir APIs REST con la librería estándar `net/http`
- ✅ Trabajar con bases de datos usando `database/sql` y GORM
- ✅ Escribir tests con el paquete `testing` estándar
- ✅ Desplegar aplicaciones con Docker

### 🚀 ¿Por qué Go?

> Go moderno desde el día 1 — Sin historia pre-1.18, solo las mejores prácticas actuales.

Go es un lenguaje de programación diseñado por Google para construir software confiable y eficiente. Simple, rápido, con concurrencia nativa y un toolchain excelente. Este bootcamp se enfoca en Go moderno con generics, structured logging y las herramientas del ecosistema oficial.

---

## 🗓️ Estructura del Bootcamp

| Etapa | Semanas | Horas | Temas |
|-------|---------|-------|-------|
| Fundamentos | 1-6 | 48h | Sintaxis, tipos, control de flujo, funciones, slices, maps, structs |
| OOP en Go | 7-10 | 32h | Punteros, métodos, interfaces, manejo de errores |
| Ecosistema | 11-13 | 24h | Paquetes, módulos, I/O, JSON, testing |
| Concurrencia | 14-17 | 32h | Goroutines, channels, select, sync, patrones |
| Web & APIs | 18-20 | 24h | net/http, REST, base de datos, Docker |

**Total: 20 semanas | 160 horas de formación intensiva**

---

## 📚 Contenido por Semana

Cada semana incluye:

```
bootcamp/week-XX-tema_principal/
├── README.md                 # Descripción y objetivos
├── rubrica-evaluacion.md     # Criterios de evaluación
├── 0-assets/                 # Imágenes y diagramas
├── 1-teoria/                 # Material teórico
├── 2-practicas/              # Ejercicios guiados
├── 3-proyecto/               # Proyecto semanal
├── 4-recursos/               # Recursos adicionales
│   ├── ebooks-free/
│   ├── videografia/
│   └── webgrafia/
└── 5-glosario/               # Términos clave
```

### 🔑 Componentes Clave

- 📖 **Teoría**: Conceptos fundamentales con ejemplos del mundo real
- 💻 **Práctica**: Ejercicios progresivos hands-on (descomenta y aprende)
- 📝 **Evaluación**: Evidencias de conocimiento, desempeño y producto
- 🎓 **Recursos**: Glosarios, referencias y material complementario

### 📅 Semanas del Bootcamp

| Semana | Tema | Etapa |
|--------|------|-------|
| 01 | [Introducción a Go y herramientas](bootcamp/week-01-introduccion_go_y_herramientas) | Fundamentos |
| 02 | [Variables, tipos y constantes](bootcamp/week-02-variables_tipos_constantes) | Fundamentos |
| 03 | [Control de flujo](bootcamp/week-03-control_de_flujo) | Fundamentos |
| 04 | [Funciones](bootcamp/week-04-funciones) | Fundamentos |
| 05 | [Arrays y Slices](bootcamp/week-05-arrays_y_slices) | Fundamentos |
| 06 | [Maps y Structs](bootcamp/week-06-maps_y_structs) | Fundamentos |
| 07 | [Punteros](bootcamp/week-07-punteros) | OOP en Go |
| 08 | [Métodos e Interfaces](bootcamp/week-08-metodos_e_interfaces) | OOP en Go |
| 09 | [Composición y Embedding](bootcamp/week-09-composicion_y_embedding) | OOP en Go |
| 10 | [Manejo de errores](bootcamp/week-10-manejo_de_errores) | OOP en Go |
| 11 | [Paquetes y Módulos](bootcamp/week-11-paquetes_y_modulos) | Ecosistema |
| 12 | [I/O estándar, archivos y JSON](bootcamp/week-12-io_archivos_y_json) | Ecosistema |
| 13 | [Testing en Go](bootcamp/week-13-testing) | Ecosistema |
| 14 | [Goroutines y el runtime](bootcamp/week-14-goroutines_y_runtime) | Concurrencia |
| 15 | [Channels](bootcamp/week-15-channels) | Concurrencia |
| 16 | [Select y sync](bootcamp/week-16-select_y_sync) | Concurrencia |
| 17 | [Patrones de concurrencia](bootcamp/week-17-patrones_de_concurrencia) | Concurrencia |
| 18 | [net/http y servidores web](bootcamp/week-18-net_http_y_servidores_web) | Web & APIs |
| 19 | [REST APIs y base de datos](bootcamp/week-19-rest_apis_y_base_de_datos) | Web & APIs |
| 20 | [Docker, CI/CD y proyecto final](bootcamp/week-20-docker_cicd_y_proyecto_final) | Web & APIs |

---

## 🛠️ Stack Tecnológico

| Herramienta | Versión | Uso |
|-------------|---------|-----|
| Go | 1.24+ | Lenguaje principal |
| golangci-lint | latest | Linting y análisis estático |
| testify | 1.10+ | Assertions en tests |
| chi | 5.2+ | Router HTTP (semanas web) |
| database/sql | stdlib | Acceso a bases de datos |
| mattn/go-sqlite3 | 1.14+ | Driver SQLite |
| jackc/pgx | 5.7+ | Driver PostgreSQL |
| Docker | 27+ | Containerización (semanas 18-20) |
| Air | 1.61+ | Hot-reload en desarrollo |

---

## 🚀 Inicio Rápido

### Prerrequisitos

- Go 1.24+ ([Instalar Go](https://go.dev/dl/))
- Git para control de versiones
- VS Code (recomendado) con extensiones incluidas

### 1. Clonar el Repositorio

```bash
git clone https://github.com/ergrato-dev/bc-go.git
cd bc-go
```

### 2. Instalar Extensiones de VS Code

```bash
# Abrir en VS Code
code .
# Las extensiones recomendadas aparecerán automáticamente
# O ejecutar: Ctrl+Shift+P → "Extensions: Show Recommended Extensions"
```

### 3. Navegar a la Semana Actual

```bash
cd bootcamp/week-01-introduccion_go_y_herramientas
```

### 4. Seguir las Instrucciones

Cada semana contiene un `README.md` con instrucciones detalladas.

---

## 📊 Metodología de Aprendizaje

### Estrategias Didácticas

- 🎯 Aprendizaje Basado en Proyectos (ABP)
- 🧩 Práctica Deliberada
- 🔄 Go Challenges (problemas idiomáticos)
- 👥 Code Review entre pares
- 🎮 Live Coding

### Distribución del Tiempo (8h/semana)

- Teoría: 2-2.5 horas
- Prácticas: 3-3.5 horas
- Proyecto: 2-2.5 horas

### Evaluación

Cada semana incluye tres tipos de evidencias:

1. Conocimiento 🧠 (30%): Cuestionarios y evaluaciones teóricas
2. Desempeño 💪 (40%): Ejercicios prácticos en clase
3. Producto 📦 (30%): Entregables evaluables (proyectos funcionales)

Criterio de aprobación: Mínimo **70%** en cada tipo de evidencia

---

## 📞 Soporte

- 💬 Discussions: [GitHub Discussions](https://github.com/ergrato-dev/bc-go/discussions)
- 🐛 Issues: [GitHub Issues](https://github.com/ergrato-dev/bc-go/issues)

---

## ⚠️ Exención de Responsabilidad

Este repositorio es un recurso educativo creado con fines de aprendizaje. Al utilizarlo, aceptas los siguientes términos:

- **Solo fines educativos**: El contenido está diseñado exclusivamente para la enseñanza y el aprendizaje.
- **Sin garantías**: El material se proporciona "tal cual", sin garantías de ningún tipo.
- **Código en producción**: Los ejemplos son ilustrativos. Antes de usarlos en entornos productivos, realiza revisiones de seguridad y adaptación.
- **Versiones de software**: Las versiones mencionadas pueden quedar desactualizadas. Consulta la documentación oficial más reciente.
- **Responsabilidad del estudiante**: Cada estudiante es responsable de sus propias implementaciones y decisiones técnicas.

---

## 📄 Licencia

Este proyecto está bajo la licencia [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/) (Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International).

- **Puedes**: compartir y adaptar el material, incluso crear forks educativos.
- **No puedes**: usar este material con fines comerciales.
- **Debes**: dar crédito apropiado y distribuir las adaptaciones bajo la misma licencia.

Ver el archivo [LICENSE](LICENSE) para el texto completo.

---

## 🏆 Agradecimientos

- [The Go Team](https://go.dev/team) — Por crear un lenguaje tan elegante y pragmático
- [Go Blog](https://go.dev/blog/) — Por los artículos de referencia
- [Go by Example](https://gobyexample.com/) — Por los ejemplos prácticos
- Comunidad Go — Por los recursos y contribuciones

---

## 📚 Documentación Adicional

- [🤖 Instrucciones de Copilot](.github/copilot-instructions.md)
- [📜 Código de Conducta](CODE_OF_CONDUCT.md)
- [🔒 Política de Seguridad](SECURITY.md)

---

<div align="center">

🎓 **Bootcamp Go — Zero to Hero**
_De cero a desarrollador Go Junior en 5 meses_

[Comenzar Semana 1](bootcamp/week-01-introduccion_go_y_herramientas) • [Ver Documentación](docs) • [Reportar Issue](https://github.com/ergrato-dev/bc-go/issues)

Hecho con ❤️ para la comunidad de desarrolladores

</div>
