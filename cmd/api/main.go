// Support-utils.
//
// API Con las funcionalidades del equipo de soporte
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"fmt"
	logs "support-utils/methods/log"
	"support-utils/methods/rest"

	"github.com/gorilla/mux"

	//logger "drivers-create/methods/log"
	"log"
	"net/http"
)

func main() {
	logs.InitLogger()
	fmt.Println("init api")
	router := mux.NewRouter()
	router.HandleFunc("/acl/{warehouse}", rest.AclPost).Methods(http.MethodPut)
	router.HandleFunc("/drivers/{warehouse}", rest.DriverPost).Methods(http.MethodPost)
	router.HandleFunc("/pickingLayout", rest.ExpeditionPickingPost).Methods(http.MethodPost)
	router.HandleFunc("/expeditionLayout", rest.ExpeditionExpeditionPost).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}
