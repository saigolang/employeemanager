package main

import (
	"bufio"
	"employeemanager/pkg/structs/response"
	"encoding/csv"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
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
	fmt.Println("employees is ", employeeResponse)
}

func load(fileName string) (employees []response.Employee) {

	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Println("error in retrieving the data from csv file is ", err.Error())
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		index, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		employees = append(employees, response.Employee{
			ID:                index[0],
			FirstName:         index[1],
			LastName:          index[2],
			DateOfBirth:       index[3],
			WorkAuthorization: index[4],
			VisaExpiryDate:    index[5],
			EndClient:         index[6],
		})
	}
	return
}
