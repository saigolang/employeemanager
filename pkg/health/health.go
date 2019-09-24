package health

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, req *http.Request) {
	respBodyBytes := new(bytes.Buffer)
	health := "OK"
	err := json.NewEncoder(respBodyBytes).Encode(&health)
	if err != nil {
		log.Println("error in marshalling the response")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBodyBytes.Bytes())
	return
}
