package controllers

import (
	"api-get-redis/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func C_getProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var output models.Output
	fmt.Println("asdasdasd")
	output.Status = "success"
	output.Message = "coba"
	payload, isValid, errMessage := models.GetProfile()
	fmt.Println(payload, isValid, errMessage)
	if isValid != true {
		output.Status = "error"
		output.Message = errMessage.Error.Error()
	} else {
		output.Status = "success"
		output.Message = ""
		output.Data = *payload
	}
	json.NewEncoder(w).Encode(output)
}
