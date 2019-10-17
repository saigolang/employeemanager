package routes

import (
	"bytes"
	"employeemanager/pkg/mapper"
	"employeemanager/pkg/structs/request"
	"employeemanager/pkg/structs/response"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetEmployeeByID(employees []response.Employee) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := mapper.GetEmployees(employees)

		// getting the user request
		request := GetUserRequest(r)

		// getting the filtered records
		filteredResponse := filterResponseByUserRequest(request, result)
		log.Println("filteredResponse is ", filteredResponse)

		respBodyBytes := new(bytes.Buffer)
		err := json.NewEncoder(respBodyBytes).Encode(&filteredResponse)
		if err != nil {
			log.Println("error in marshalling the response")
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respBodyBytes.Bytes())
		return
	}
}

func filterResponseByUserRequest(req request.Request, response []response.FinalResponse) (employees []response.FinalResponse) {
	log.Println("request inside filter finction is", req)
	for _, r := range response {
		if req.ID != "" && r.ID != req.ID {
			continue
		}
		if req.FirstName != "" && r.FirstName != req.FirstName {
			continue
		}
		if req.LastName != "" && r.LastName != req.LastName {
			continue
		}
		if req.WorkAuthorization != "" && r.WorkAuthorization != req.WorkAuthorization {
			continue
		}
		if req.EndClient != "" && r.EndClient != req.EndClient {
			continue
		}
		employees = append(employees, r)
	}
	return
}
func GetUserRequest(r *http.Request) (req request.Request) {
	keys := r.URL.Query()
	req.ID = keys.Get("id")
	req.EndClient = keys.Get("endClient")
	req.WorkAuthorization = keys.Get("workAuthorization")
	req.FirstName = keys.Get("firstName")
	req.LastName = keys.Get("lastName")
	return
}

func GetAllEmployeesInformation(employees []response.Employee) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		/*	employeeResponse := load("employee.csv")
			result := mapData(employeeResponse)*/
		// let's do the mapping

		request := GetUserRequest(r)

		log.Println("request is ", request)
		finalResponse := mapper.GetEmployees(employees)

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
