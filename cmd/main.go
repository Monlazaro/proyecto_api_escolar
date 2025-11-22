package main

import (
	"log"
	"net/http"
	"proyecto_api_escolar/internal/database"
	"proyecto_api_escolar/internal/handlers"
	"github.com/gorilla/mux"
)

func main(){
	//iniciar la base de datos
	database.InitDatabase()

	//crear el enrutador
	r := mux.NewRouter()

	//configurar rutas para estudiantes
	r.HandleFunc("/api/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/api/students", handlers.GetStudents).Methods("GET")

	// rutas para materias
	r.HandleFunc("/api/subjects", handlers.CreateSubject).Methods("POST")
	r.HandleFunc("/api/subjects/{subject_id}", handlers.GetSubject).Methods("GET")
	r.HandleFunc("/api/subjects/{subject_id}", handlers.UpdateSubject).Methods("PUT")
	r.HandleFunc("/api/subjects/{subject_id}", handlers.DeleteSubject).Methods("DELETE")

	//rutas para calificaciones
	r.HandleFunc("/api/grades", handlers.CreateGrade).Methods("POST")
r.HandleFunc("/api/grades/{grade_id}", handlers.UpdateGrade).Methods("PUT")
r.HandleFunc("/api/grades/{grade_id}", handlers.DeleteGrade).Methods("DELETE")
r.HandleFunc("/api/grades/{grade_id}/student/{student_id}", handlers.GetGradeByStudentAndSubject).Methods("GET")
r.HandleFunc("/api/grades/student/{student_id}", handlers.GetGradesByStudent).Methods("GET")

	//iniciar el servidor en el puerto 8081
	log.Println("Servidor iniciado en http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}