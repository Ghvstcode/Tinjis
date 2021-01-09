package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/GhvstCode/tinjis/paymentsvc/utils"
	"math/rand"
	"net/http"
)
type PaymentModel struct {
	ID       int     `json:"customer_id"`
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

func Payments(w http.ResponseWriter, r *http.Request) {
	p := &PaymentModel{}
	prob := rand.Intn(2)
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		fmt.Println(err)
		utils.Response(map[string]string{"error": "Bad Request"}, http.StatusBadRequest).Send(w)
		return
	}

	if p.Currency == "" {
		utils.Response(map[string]string{"error": "Currency Field is required"}, http.StatusBadRequest).Send(w)
		return
	}
	if p.ID < 1 {
		utils.Response(map[string]string{"error": "ID must be an Int greater than 1"}, http.StatusBadRequest).Send(w)
		return
	}
	if p.Value < 0 {
		utils.Response(map[string]string{"error": "Value Must be a positive Int"}, http.StatusBadRequest).Send(w)
		return
	}
	if prob == 0 {
		utils.Response(map[string]string{"result": "true"}, http.StatusBadRequest).Send(w)
		return
	}
	utils.Response(map[string]string{"result": "false"}, http.StatusBadRequest).Send(w)
	return
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := "OK"
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res);  err != nil  {
		fmt.Println(err)
	}
}
