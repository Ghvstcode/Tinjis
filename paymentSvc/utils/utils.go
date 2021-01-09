package utils

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	StatusCode int
	Data    interface{}
}

func (data Data) Send(w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(data.StatusCode)
	if err := json.NewEncoder(w).Encode(data);  err != nil  {
		return err
	}
	return nil
}

func Response(payload interface{}, statusCode int) *Data {
	return &Data{
		StatusCode: statusCode,
		Data:    payload,
	}
}
