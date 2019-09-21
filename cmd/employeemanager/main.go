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

func main() {
	fmt.Println("This is a main method")
	router := mux.NewRouter().StrictSlash(true)
	// todo configure logging here
	router.HandleFunc("/employees", GetAllEmployeesInformation).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAllEmployeesInformation(w http.ResponseWriter, r *http.Request) {
	csvFile, err := os.Open("employee.csv")
	if err != nil {
		log.Println("error in retrieving the data from csv file is ", err.Error())
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var employees []response.Employee
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

	log.Println("employee data is ", employees)

	//employeeResponse, err := json.Marshal(employees)
}
