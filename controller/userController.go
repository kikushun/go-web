package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UserController aa
func UserController(mux *http.ServeMux) {
	mux.Handle("/user", Handler{getUser})
	mux.Handle("/users", Handler{getUsers})
}

// /user
func getUser(w http.ResponseWriter, req *http.Request) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:9200/", nil)
	if err != nil {
		fmt.Println("err!")
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err!!!!!!")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err....")
	}

	var jsonBody interface{}
	err = json.Unmarshal(body, &jsonBody)

	fmt.Fprintln(w, jsonBody)
}

func getUsers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "users")
}
