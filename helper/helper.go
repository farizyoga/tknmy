package helper

import (
	"encoding/json"
	"net/http"
)

func CreateResponse(w http.ResponseWriter, code int, data interface{}, message string) {
	r := make(map[string]interface{})

	r["data"] = data
	r["message"] = message
	r["code"] = code

	w.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(r)

	if err != nil {
		panic("error creating response")
	}

	w.WriteHeader(code)
	w.Write(json)
}
