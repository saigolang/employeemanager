package main

import (
	"bytes"
	"employeemanager/pkg/health"
	"employeemanager/pkg/routes"
	"employeemanager/pkg/structs/response"
	"encoding/csv"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	fileName = "employee.csv"
)

func main() {
	// todo configure logging here

	//load the file
	rawCSVData := load("employee.csv")
	// employees response
	employees := mapData(rawCSVData)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/employees", routes.GetAllEmployeesInformation(employees)).Methods(http.MethodGet)
	router.HandleFunc("/employee", routes.GetEmployeeByID(employees)).Methods(http.MethodGet)
	router.HandleFunc("/health", health.Health).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func load(fileName string) (csvData [][]string) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("error is ", err.Error())
	}
	reader := csv.NewReader(bytes.NewReader(fileBytes))
	//setting it to -1 so that if any record's column value is empty it will be ignored
	reader.FieldsPerRecord = -1
	csvData, err = reader.ReadAll()
	if err != nil {
		log.Println("error is ", err.Error())
	}
	if len(csvData) == 0 {
		log.Println("no data found")
	}
	return
}

func mapData(rawData [][]string) (employees []response.Employee) {
	for _, row := range rawData[1:] {
		employees = append(employees, response.Employee{
			ID:                row[0],
			FirstName:         row[1],
			LastName:          row[2],
			DateOfBirth:       row[3],
			WorkAuthorization: row[4],
			VisaExpiryDate:    row[5],
			EndClient:         row[6],
		})
	}
	return
}
