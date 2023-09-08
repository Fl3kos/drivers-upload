package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"support-utils/methods/createUsers"
	"support-utils/structs/handlers"
)

func DriverPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	warehouseCode := vars["warehouse"]
	var drivers handlers.Drivers

	err := json.NewDecoder(r.Body).Decode(&drivers)
	if err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)

		return
	}

	response, incorrectDnis := createUsers.CreateDrivers(drivers.DriverA, warehouseCode)
	if incorrectDnis != nil {
		errorLog := fmt.Sprintf("Error con los sientes Dnis: %v", incorrectDnis)
		http.Error(w, errorLog, http.StatusBadRequest)
		return

	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Establecer la cabecera Content-Type a "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON
	w.Write(jsonData)
}

func AclPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	warehouseCode := vars["warehouse"]
	var warehouseUsers handlers.WarehouseUsers

	err := json.NewDecoder(r.Body).Decode(&warehouseUsers)
	if err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusInternalServerError)
		return
	}

	auth := w.Header().Get("Authorization")

	err = createUsers.CreateWarehouseUsers(warehouseUsers, warehouseCode, auth)
	if err != nil {
		http.Error(w, "Error al dar de alta los usuarios", http.StatusNotFound)
		return
	}

	// Establecer la cabecera Content-Type a "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON
	w.Write([]byte("Usuarios dados de alta"))
}
