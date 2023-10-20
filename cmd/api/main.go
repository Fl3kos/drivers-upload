// SimpsonGo.
//
// API que sirve frases aletorias de los Simpsons. ¿Quieres ver cómo va su desarrollo? https://osmandi.now.sh/post/simpsongo
//
//	Schemes: https
//	Host: simpsongo.now.sh
//	BasePath: /
//	Version: 1.0
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: https://osmandi.now.sh
//
//
//	Consumes:
//	- application/json
//	- application/xml
//
//	Produces:
//	- application/json
//	- application/xml
//
// swagger:meta
package main

import (
	"encoding/json"
	https "support-utils/methods/http"
	logs "support-utils/methods/log"
	"support-utils/methods/rest"

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
	router.HandleFunc("/acl/{warehouse}", rest.AclPost).Methods(http.MethodPut)
	router.HandleFunc("/drivers/{warehouse}", rest.DriverPost).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}
