package util_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "util"

	"github.com/gorilla/mux"
)

// User Story - Users should be able to view list of User entity
func TestGetUsers(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

// User Story - Users should be able to create a User entity
func TestCreateUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST")
	userJson := `{"firstname": "shiju", "lastname": "Varghese", "email":
"shiju@xyz.com"}`
	req, err := http.NewRequest(
		"POST",
		"/users",
		strings.NewReader(userJson),
	)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", w.Code)
	}
}

//User Story - The Email Id of a User entity should be unique
func TestUniqueEmail(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST")
	userJson := `{"firstname": "shiju", "lastname": "Varghese", "email":
"shiju@xyz.com"}`
	req, err := http.NewRequest(
		"POST",
		"/users",
		strings.NewReader(userJson),
	)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Error("Bad Request expected, got: %d", w.Code)
	}
}

func TestGetUsersClient(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	usersUrl := fmt.Sprintf("%s/users", server.URL)
	fmt.Println("TestGetUsersClient:" + usersUrl)
	request, err := http.NewRequest("GET", usersUrl, nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
	fmt.Printf("%s----response:%s\n", res.Status)
}

func TestCreateUserClient(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()
	usersUrl := fmt.Sprintf("%s/users", server.URL)
	fmt.Println("TestCreateUserClient" + usersUrl)
	userJson := `{"firstname": "Rosmi", "lastname": "Shiju", "email": "rose@xyz.com"}`
	request, err := http.NewRequest("POST", usersUrl, strings.NewReader(userJson))
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}
}
