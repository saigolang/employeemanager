package mapper

import "employeemanager/pkg/structs/response"

func getEmployees(employees []response.Employee) (finalResponse []response.Employee) {
	for _, r := range employees {
		employees = append(employees, response.Employee{
			ID:                r.ID,
			FirstName:         r.FirstName,
			LastName:          r.LastName,
			DateOfBirth:       r.DateOfBirth,
			WorkAuthorization: r.WorkAuthorization,
			VisaExpiryDate:    r.VisaExpiryDate,
			EndClient:         r.EndClient,
		})
	}
	return
}
