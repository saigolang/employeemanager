package app

import (
	"employeemanager/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run() {
	router := mux.NewRouter().StrictSlash(true)
	// todo configure logging here
	router.HandleFunc("/employees", routes.GetAllEmployeesInformation).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", router))
}
