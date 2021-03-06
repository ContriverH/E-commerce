package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codernishchay/productapi/app/models"
	// "github.com/katoozi/golang-mongodb-rest-api/app/model"
)

// ResponseWriter will write result in http.ResponseWriter
func ResponseWriter(res http.ResponseWriter, statusCode int, message string, data interface{}) error {
	res.WriteHeader(statusCode)
	httpResponse := models.NewResponse(statusCode, message, data)
	err := json.NewEncoder(res).Encode(httpResponse)
	return err
}
