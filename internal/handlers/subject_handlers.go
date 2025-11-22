package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"proyecto_api_escolar/internal/database"
	"proyecto_api_escolar/internal/models"

	"github.com/gorilla/mux"
)

func CreateSubject(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var subject models.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, "Cuerpo JSON inválido", http.StatusBadRequest)
		return
	}

	if subject.Name == "" {
		http.Error(w, "El campo 'name' es obligatorio", http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&subject)
	if result.Error != nil {
		http.Error(w, "Error al crear materia: " +result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subject)
}

func GetSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["subject_id"])
	if err != nil {
		http.Error(w, "ID de materia inválido", http.StatusBadRequest)
		return
	}

	var subject models.Subject
	result := database.DB.First(&subject, id)
	if result.Error != nil {
		http.Error(w, "Materia no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(subject)
}

func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["subject_id"])
	if err != nil {
		http.Error(w, "ID de materia inválido", http.StatusBadRequest)
		return
	}

	var updatedSubject models.Subject
	if err := json.NewDecoder(r.Body).Decode(&updatedSubject); err != nil {
		http.Error(w, "Cuerpo JSON inválido", http.StatusBadRequest)
		return
	}

	if updatedSubject.Name == "" {
		http.Error(w, "El campo 'name' es obligatorio", http.StatusBadRequest)
		return
	}

	var subject models.Subject
	result := database.DB.First(&subject, id)
	if result.Error != nil {
		http.Error(w, "Materia no encontrada", http.StatusNotFound)
		return
	}

	subject.Name = updatedSubject.Name
	result = database.DB.Save(&subject)
	if result.Error != nil {
		http.Error(w, "Error al actualizar materia", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(subject)
}

func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["subject_id"])
	if err != nil {
		http.Error(w, "ID de materia inválido", http.StatusBadRequest)
		return
	}

	var subject models.Subject
	result := database.DB.First(&subject, id)
	if result.Error != nil {
		http.Error(w, "Materia no encontrada", http.StatusNotFound)
		return
	}

	result = database.DB.Delete(&subject)
	if result.Error != nil {
		http.Error(w, "Error al eliminar materia", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}		