package api

import (
	"encoding/json"
	"net/http"
)

func QueryHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		NotFound(w, r)
		return
	case "POST":
		query := r.URL.Path[len("/query/"):]
		switch query {
		case "checkExistence":
			checkExistence(w, r)
		default:
			NotFound(w, r)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func checkExistence(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}
	json.NewDecoder(r.Body).Decode(&data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response = struct {
		Message string `json:"message"`
		OK      bool   `json:"ok"`
	}{}
	switch data.Field {
	case "username":
		u := User{Username: data.Value}
		if u.UsernameExists() {
			response.Message = "Username already exists"
			response.OK = false
		} else {
			response.Message = "New username"
			response.OK = true
		}
	case "email":
		u := User{Email: data.Value}
		if u.EmailExists() {
			response.Message = "Email already exists"
			response.OK = false
		} else {
			response.Message = "New email"
			response.OK = true
		}
	default:
		http.Error(w, "Field not found", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(response)
}
