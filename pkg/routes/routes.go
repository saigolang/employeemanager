package routes

import (
	"bytes"
	"employeemanager/pkg/mapper"
	"employeemanager/pkg/structs/response"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {

}

func GetAllEmployeesInformation(w http.ResponseWriter, r *http.Request) {
	employeeResponse := load("employee.csv")
	result := mapData(employeeResponse)
	// let's do the mapping
	finalResponse := mapper.GetEmployees(result)

	respBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(respBodyBytes).Encode(&finalResponse)
	if err != nil {
		log.Println("error in marshalling the response")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBodyBytes.Bytes())
	fmt.Println("employees is ", finalResponse)
	return
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