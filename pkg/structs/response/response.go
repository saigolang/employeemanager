package response

type Employee struct {
	ID                string `json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	DateOfBirth       string `json:"dateOfBirth"`
	WorkAuthorization string `json:"workAuthorization"`
	VisaExpiryDate    string `json:"visaExpiryDate"`
	EndClient         string `json:"endClient"`
}

type FinalResponse struct {
	ID                string `json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	DateOfBirth       string `json:"dateOfBirth"`
	WorkAuthorization string `json:"workAuthorization"`
	VisaExpiryDate    string `json:"visaExpiryDate"`
	EndClient         string `json:"endClient"`
}
