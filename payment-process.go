package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := &Person{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/user/register", createUser)
	http.ListenAndServe(":8081", nil)
}
