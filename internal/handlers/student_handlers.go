package handlers

import (
	"encoding/json"
	"net/http"
	"proyecto_api_escolar/internal/database"
	"proyecto_api_escolar/internal/models"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Cuerpo JSON inválido", http.StatusBadRequest)
		return
	}

	// Validación básica
	if student.Name == "" {
		http.Error(w, "El campo 'name' es obligatorio", http.StatusBadRequest)
		return
	}
	if student.Email == "" {
		http.Error(w, "El campo 'email' es obligatorio", http.StatusBadRequest)
		return
	}

	// Guardar en la base de datos
	result := database.DB.Create(&student)
	if result.Error != nil {
		http.Error(w, "Error al crear estudiante: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var students []models.Student
	database.DB.Find(&students)
	json.NewEncoder(w).Encode(students)
}