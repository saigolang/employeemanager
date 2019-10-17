package mapper

import (
	"employeemanager/pkg/structs/response"
	"log"
)

func GetEmployees(employees []response.Employee) (finalResponse []response.FinalResponse) {
	for _, r := range employees {
		finalResponse = append(finalResponse, response.FinalResponse{
			ID:                r.ID,
			FirstName:         r.FirstName,
			LastName:          r.LastName,
			DateOfBirth:       r.DateOfBirth,
			WorkAuthorization: r.WorkAuthorization,
			VisaExpiryDate:    r.VisaExpiryDate,
			EndClient:         r.EndClient,
		})
	}
	log.Println("employees is ", employees)
	return
}

func GetEmployee() {

}
