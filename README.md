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
    git clone https://github.com/Monlazaro/proyecto_api_escolar.git
    cd proyecto_api_escolar
2.- Instalar dependencias
    bash
    go mod tidy
3.- Ejecutar la API
    bash
    go run cmd/main.go
**La API se iniciará en http://localhost:8081**

## Endpoints disponibles
- Estudiantes (/api/students)
| Método | Endpoint                          | Descripción                          |
|--------|-----------------------------------|--------------------------------------|
| POST   | `/api/students`                   | Crear un nuevo estudiante            |
| GET    | `/api/students`                   | Obtener todos los estudiantes        |
| GET    | `/api/students/{student_id}`      | Obtener un estudiante por ID         |
| PUT    | `/api/students/{student_id}`      | Actualizar un estudiante             |
| DELETE | `/api/students/{student_id}`      | Eliminar un estudiante               |

- Materias (/api/subjects)
| Método | Endpoint                           | Descripción                     |
|--------|------------------------------------|---------------------------------|
| POST   | `/api/subjects`                    | Crear una nueva materia         |
| GET    | `/api/subjects/{subject_id}`       | Obtener una materia por ID      |
| PUT    | `/api/subjects/{subject_id}`       | Actualizar una materia          |
| DELETE | `/api/subjects/{subject_id}`       | Eliminar una materia            |

- Calificaciones (/api/grades)
| Método | Endpoint                                            | Descripción                                      |
|--------|-----------------------------------------------------|--------------------------------------------------|
| POST   | `/api/grades`                                       | Crear una nueva calificación                     |
| PUT    | `/api/grades/{grade_id}`                            | Actualizar una calificación                      |
| DELETE | `/api/grades/{grade_id}`                            | Eliminar una calificación                        |
| GET    | `/api/grades/{grade_id}/student/{student_id}`        | Obtener una calificación específica              |
| GET    | `/api/grades/student/{student_id}`                  | Obtener todas las calificaciones de un estudiante|

## Ejemplos de uso (Postman)

**Crear un estudiante**
* Método: POST
* URL: http://localhost:8081/api/students
* Body >> raw >> JSON:
  {
  "name": Carlos Eduardo Valdez Hernandez
  "group": A
  "email": cvaldez27@alumnos.uaq.mx
  }

**Crear una materia**
   * Método: POST
   * URL: http://localhost:8081/api/subjects
   * Body >> raw >> JSON:
  {
  "name": Geografía
  }

**Asignar una calificación**
    * Método: POST
    * URL: http://localhost:8081/api/grades
    * Body >> raw >> JSON:
  {
      "student_id": 1,
      "subject_id": 1,
      "grade": 9.8
  }

**Obtener todas las calificaciones de un estudiante**
    * Método: GET
    * URL: http://localhost:8081/api/grades/student/1
    * Send
- se puede cambiar el ID por el que se quiere consultar, diferente a 1 ya que éste ya está asignado al alumno Aldo Ugalde Olguín.

**Obtener una calificación específica**
    * Método: GET
    * URL: http://localhost:8081/api/grades/1/student/1
    * Send
- se pueden cambiar los ID por los que se quieren consultar, diferente a 1 ya que éste ya está asignado al alumno Aldo Ugalde Olguín.

**Obtener todas las calificaciones de un estudiante**
    bash
    curl http://localhost:8081/api/grades/student/1

## Se añadieron las características adicionales
- Persistencia: Los datos se almacenan en school.db (SQLite)
- Validación de datos:
      * email (único y obligatorio)
      * name (obligatorio en estudiantes y materias
      * grade (entre 0 y 10)
- Llaves foráneas: Se validan las relaciones entre grades, students y subjects.
- Relaciones: Las respuestas incluyen datos anidados del estudiante y la materia

## Nota importante:
El servidor debe estar en ejecución (go run cmd/main.go) antes de hacer las peticiones.
