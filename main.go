package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

)

type Student struct {
	ID int
	Name string
	Age int8
	Class []int
}

type Students []Student 

func main()  {
	fmt.Println("My Web App")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/about", AboutPage)
	myRouter.HandleFunc("/api/mucsic", MucsicApi)
	myRouter.HandleFunc("/api/student", StudentApi)
	myRouter.HandleFunc("/api/students", ListStudentsApi).Methods("GET")
	myRouter.HandleFunc("/api/students", testPostStudentsApi).Methods("POST")


	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func testPostStudentsApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Test Post Students Api")
}

func StudentApi(w http.ResponseWriter, r *http.Request) {
	var student = Student { ID: 1, Name: "Vinh", Age: 20, Class: []int{1, 2, 3, 4, 5}}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func ListStudentsApi(w http.ResponseWriter, r *http.Request) {
	var ListStudent = Students{
		Student { ID: 1, Name: "Hoa", Age: 20, Class: []int{1, 2, 3, 4, 5}},
		Student { ID: 2, Name: "Tuan", Age: 21, Class: []int{1, 2, 3, 4, 5}},
		Student { ID: 3, Name: "Kiet", Age: 22, Class: []int{1, 2, 3, 4, 5}},
		Student { ID: 4, Name: "Nhu", Age: 23, Class: []int{1, 2, 3, 4, 5}},
		Student { ID: 5, Name: "Thach", Age: 24, Class: []int{1, 2, 3, 4, 5}},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ListStudent)
}

func MucsicApi(w http.ResponseWriter, r *http.Request) {
	var data = map[string] interface{} {
		"nama": "Loi Cam On",
		"casi": "Tuan Tai",
	}	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the HomePage!</h1>")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the AboutPage!</h1>")
}