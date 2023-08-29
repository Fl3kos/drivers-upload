package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"support-utils/consts"
	"support-utils/methods/createUsers"
	logs "support-utils/methods/log"
	"support-utils/structs/handlers"
	"support-utils/structs/responses"
)

func AuthEndpointCall(usersJson string) {
	logs.Debugln("URL to POST:", consts.AuthEndpointUrl)

	var jsonStr = []byte(usersJson)
	req, err := http.NewRequest("POST", consts.AuthEndpointUrl, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	logs.Debugln("response Status:", resp.Status)
	if resp.Status != "204 No Content" {
		fmt.Println("Error calling endpoint, check logs")
		logs.Errorln("Error sending the data, check ACL and couchbase, after publish with postman")
		return
	}
}

func AclEndpointCall(usersJson, username, token string) {
	url := fmt.Sprintf(consts.AclEndpointUrl, username)
	logs.Debugln("URL to PUT:", url)
	var jsonStr = []byte(usersJson)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authority", "com.pro.api.dgrp.io")
	req.Header.Set("accept", "application/json, text/plain")
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("client_id", "dKSBbUAiriDZzZoVC9pLqstZHsCD0tJfx6GycX3Ox9FIG4cm")
	req.Header.Set("origin", "https://acl-web.com.pro.webs.dgrp.io")
	req.Header.Set("referer", "https://acl-web.com.pro.webs.dgrp.io")
	req.Header.Set("x-diagroup-application-id", "ACL")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	logs.Debugln("response Status:", resp.Status)
	if resp.Status != "204 No Content" {
		fmt.Println("Error calling endpoint, check logs")
		logs.Errorln("Error sending the data, check ACL and couchbase, after publish with postman")
		return
	}
}

type Message struct {
	Text string `json:"text"`
}

func HandleMessage(w http.ResponseWriter, r *http.Request) {
	// Decodificar los datos JSON recibidos en una estructura Message
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		fmt.Println("Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	// Realizar alguna acción con los datos recibidos
	//logger.Debugf("Mensaje recibido:", message.Text)

	// Responder con un mensaje de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message.Text))
}

func DriverPost (w http.ResponseWriter, r *http.Request) {
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


func DriverGet(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	// Decodificar los datos JSON recibidos en una estructura Message
	var drivers handlers.Drivers
	err := json.NewDecoder(r.Body).Decode(&drivers)

	if err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(drivers)

	response := responses.DriverResponse{
		Name:        drivers.DriverA[0].Name,
		Username: drivers.DriverA[0].PhoneNumber,
		Password:         "drivers.DriverA[0].Dni",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Establecer la cabecera Content-Type a "application/json"
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON
	w.Write(jsonData)
}
