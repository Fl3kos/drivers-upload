package main

import (
	"encoding/json"
	https "support-utils/methods/http"
	logs "support-utils/methods/log"

	"github.com/gorilla/mux"

	//logger "drivers-create/methods/log"
	"log"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	// Supongamos que obtienes los datos del usuario por su ID
	user := User{
		ID:       userID,
		Username: "john_doe",
		Email:    "john@example.com",
	}

	// Convertir el objeto User a JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Establecer la cabecera Content-Type a "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON
	w.Write(jsonData)
}

func main() {
	logs.InitLogger()
	router := mux.NewRouter()
	router.HandleFunc("/message", https.HandleMessage).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", GetUserByID).Methods(http.MethodGet)
	router.HandleFunc("/driver", https.DriverGet).Methods(http.MethodPost)
	router.HandleFunc("/drivers/{warehouse}", https.DriverPost).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}
