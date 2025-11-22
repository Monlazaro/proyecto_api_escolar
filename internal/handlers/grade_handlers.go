package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"proyecto_api_escolar/internal/database"
	"proyecto_api_escolar/internal/models"

	"github.com/gorilla/mux"
)

func CreateGrade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var grade models.Grade
	if err := json.NewDecoder(r.Body).Decode(&grade); err != nil {
		http.Error(w, "Cuerpo JSON inválido", http.StatusBadRequest)
		return
	}

	//validación: calificación entre 0 y 10
	if grade.Grade <0 || grade.Grade > 10 {
		http.Error(w, "La calificación debe estar entre 0 y 10", http.StatusBadRequest)
		return
	}

	//validar que el estudiante exista
	var student models.Student
	if err := database.DB.First(&student, grade.StudentID).Error; err != nil {
		http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
		return
	}

	//validar que la materia exista
	var subject models.Subject
	if err := database.DB.First(&subject, grade.SubjectID).Error; err != nil {
		http.Error(w, "Materia no encontrada", http.StatusNotFound)
		return
	}

	//guardar calificación
	result := database.DB.Create(&grade)
	if result.Error != nil {
		http.Error(w, "Error al crear calificación: " +result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(grade)
}

func GetGradeByStudentAndSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	gradeID, _ := strconv.Atoi(vars["grade_id"])
	studentID, _ := strconv.Atoi(vars["student_id"])

	var grade models.Grade
	result := database.DB.Where("grade_id = ? AND student_id = ?", gradeID, studentID).First(&grade)
	if result.Error != nil {
		http.Error(w, "Calificación no encontrada", http.StatusNotFound)
		return
	}

	// Cargar relaciones para mostrar datos completos (punto extra)
	database.DB.Preload("Student").Preload("Subject").First(&grade, gradeID)

	json.NewEncoder(w).Encode(grade)
}

func GetGradesByStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	studentID, err := strconv.Atoi(vars["student_id"])
	if err != nil {
		http.Error(w, "ID de estudiante inválido", http.StatusBadRequest)
		return
	}

	var grades []models.Grade
	result := database.DB.Where("student_id = ?", studentID).Find(&grades)
	if result.Error != nil {
		http.Error(w, "Error al obtener calificaciones", http.StatusInternalServerError)
		return
	}

	// Cargar relaciones (punto extra)
	for i := range grades {
		database.DB.Preload("Student").Preload("Subject").First(&grades[i], grades[i].GradeID)
	}

	json.NewEncoder(w).Encode(grades)
}

// UpdateGrade maneja PUT /api/grades/:grade_id
func UpdateGrade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	gradeID, err := strconv.Atoi(vars["grade_id"])
	if err != nil {
		http.Error(w, "ID de calificación inválido", http.StatusBadRequest)
		return
	}

	var updatedGrade models.Grade
	if err := json.NewDecoder(r.Body).Decode(&updatedGrade); err != nil {
		http.Error(w, "Cuerpo JSON inválido", http.StatusBadRequest)
		return
	}

	if updatedGrade.Grade < 0 || updatedGrade.Grade > 10 {
		http.Error(w, "La calificación debe estar entre 0 y 10", http.StatusBadRequest)
		return
	}

	var grade models.Grade
	result := database.DB.First(&grade, gradeID)
	if result.Error != nil {
		http.Error(w, "Calificación no encontrada", http.StatusNotFound)
		return
	}

	// Validar que el estudiante y materia (si cambian) existan
	if updatedGrade.StudentID != 0 {
		var s models.Student
		if err := database.DB.First(&s, updatedGrade.StudentID).Error; err != nil {
			http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
			return
		}
		grade.StudentID = updatedGrade.StudentID
	}
	if updatedGrade.SubjectID != 0 {
		var s models.Subject
		if err := database.DB.First(&s, updatedGrade.SubjectID).Error; err != nil {
			http.Error(w, "Materia no encontrada", http.StatusNotFound)
			return
		}
		grade.SubjectID = updatedGrade.SubjectID
	}

	grade.Grade = updatedGrade.Grade

	result = database.DB.Save(&grade)
	if result.Error != nil {
		http.Error(w, "Error al actualizar calificación", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(grade)
}

// DeleteGrade maneja DELETE /api/grades/:grade_id
func DeleteGrade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	gradeID, err := strconv.Atoi(vars["grade_id"])
	if err != nil {
		http.Error(w, "ID de calificación inválido", http.StatusBadRequest)
		return
	}

	var grade models.Grade
	result := database.DB.First(&grade, gradeID)
	if result.Error != nil {
		http.Error(w, "Calificación no encontrada", http.StatusNotFound)
		return
	}

	result = database.DB.Delete(&grade)
	if result.Error != nil {
		http.Error(w, "Error al eliminar calificación", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}