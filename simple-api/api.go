package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	EnterDate string `json:"enterDate"`
	LeaveDate string `json:"leaveDate"`
}

func registerEnter(w http.ResponseWriter, r *http.Request) {
	user := User{ Id: "123", Name: "Jacob Meneses", EnterDate: "2018-01-01", LeaveDate: "2018-02-02" }

	b, _ := json.Marshal(user)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func registerLeave(w http.ResponseWriter, r *http.Request) {
	user := User{ Id: "123", Name: "Jacob Meneses", EnterDate: "2018-01-01", LeaveDate: "2018-02-02" }

	b, _ := json.Marshal(user)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func main() {
	http.HandleFunc("/register-enter/", registerEnter)
	http.HandleFunc("/register-leave/", registerLeave)

	fmt.Println("Running server on: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}