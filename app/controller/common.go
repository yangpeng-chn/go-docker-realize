package controller

import (
	"encoding/json"
	"net/http"
)

func SetResponseHeaders(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Language, Content-Type")
}

func WriteResponse(w http.ResponseWriter, code int, body interface{}) error {
	res, err := json.Marshal(body)
	if err != nil {
		return err
	}
	w.WriteHeader(code)
	w.Write(res)
	return nil
}
