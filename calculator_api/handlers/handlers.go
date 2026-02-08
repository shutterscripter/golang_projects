package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type Request struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}
type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Something went Wrong, Try again!", http.StatusBadRequest)
		return
	}

	result := req.A + req.B
	fmt.Println("Result :", result)
	resp := Response{
		Result: result,
	}
	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Something went Wrong, Try again!", http.StatusInternalServerError)
	}

}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Something went Wrong, Try again!", http.StatusBadRequest)
		return
	}

	result := req.A - req.B
	fmt.Println("Result :", math.Round(result))
	resp := Response{
		Result: result,
	}
	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Something went Wrong, Try again!", http.StatusInternalServerError)
	}

}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Something went Wrong, Try again!", http.StatusBadRequest)
		return
	}

	var result float64 = req.A * req.B
	fmt.Println("Result :", math.Round(result))
	resp := Response{
		Result: result,
	}
	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Something went Wrong, Try again!", http.StatusInternalServerError)
	}

}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Something went Wrong, Try again!", http.StatusBadRequest)
		return
	}

	if req.B == 0 {
		http.Error(w, "Value of B can not be 0", http.StatusBadRequest)
		return
	}

	result := req.A / req.B
	fmt.Println("Result :", result)
	resp := Response{
		Result: result,
	}
	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Something went Wrong, Try again!", http.StatusInternalServerError)
	}

}
