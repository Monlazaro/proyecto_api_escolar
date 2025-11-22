# Sistema de Control Escolar (API REST)

API RESTful desarrollada en **Go** para gestionar estudiantes, materias y calificaciones. Se implementan operaciones CRUD y cumple con los requisitos solicitados en el documento del Proyecto Final.

## Tecnologías utilizadas
- **Lenguaje:** Go
- **Framework de rutas:** gorilla/mux
- **Base de datos:** SQLite
- **ORM:** GORM
- **Formato de respuesta:** JSON

## Estructura del proyecto
proyecto_api_escolar/
├── cmd/
│ └── main.go # Punto de entrada
├── internal/
│ ├── models/ # Modelos de datos (student, subject, grade)
│ ├── handlers/ # Lógica de los     endpoints (student, subject, grade)
│ └── database/ # Conexión y migración de la base de datos (bd)
├── go.mod
├── go.sum
└── README.md

## Instalación y ejecución

1.- Clonar el repositorio
    bash
    git clone 