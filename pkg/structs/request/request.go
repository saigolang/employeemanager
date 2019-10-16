package request

type Request struct {
	ID                string `json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	WorkAuthorization string `json:"workAuthorization"`
	EndClient         string `json:"endClient"`
}
