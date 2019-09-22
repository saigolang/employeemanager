package main

import (
	"bytes"
	"employeemanager/pkg/mapper"
	"employeemanager/pkg/structs/response"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	fileName = "employee.csv"
)

func main() {
	fmt.Println("This is a main method")
	router := mux.NewRouter().StrictSlash(true)
	// todo configure logging here
	router.HandleFunc("/employees", GetAllEmployeesInformation).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAllEmployeesInformation(w http.ResponseWriter, r *http.Request) {
	employeeResponse := load(fileName)
	// get csv raw data
	//rawData := load(employeeResponse)

	result := mapData(employeeResponse)
	// let's do the mapping
	finalResponse := mapper.GetEmployees(result)

	respBobyBytes := new(bytes.Buffer)
	err := json.NewEncoder(respBobyBytes).Encode(&finalResponse)
	if err != nil {
		log.Println("error in marshalling the response")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBobyBytes.Bytes())
	fmt.Println("employees is ", finalResponse)
	return
}

func load(fileName string) (csvData [][]string) {

	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("error is ", err.Error())
	}
	reader := csv.NewReader(bytes.NewReader(fileBytes))
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
