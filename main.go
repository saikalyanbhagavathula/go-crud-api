package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	Gender string `json:"gender"`
}

var students []Student

func getDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")

	json.NewEncoder(w).Encode(students)
	return
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student Student

	json.NewDecoder(r.Body).Decode(&student)

	for _, item := range students {
		if student.Id == item.Id {
			fmt.Fprintf(w, "Id matches with another record...!")
			return
		}
	}
	students = append(students, student)
	json.NewEncoder(w).Encode(students)
	return
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	var params = mux.Vars(r)

	for index, item := range students {
		if params["id"] == item.Id {
			students = append(students[:index], students[index+1:]...)
			break
		}
		fmt.Println("Student Not Found")
	}
	json.NewEncoder(w).Encode(students)
	return
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	for index, item := range students {
		if params["id"] == item.Id {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}

	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	students = append(students, student)
	json.NewEncoder(w).Encode(students)
	return

}
func main() {

	students = append(students, Student{Name: "Name1", Id: "123", Gender: "Male"})
	students = append(students, Student{Name: "Name2", Id: "456", Gender: "Female"})

	router := mux.NewRouter()

	router.HandleFunc("/students", getDetails).Methods("GET")
	router.HandleFunc("/create", addStudent).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteStudent).Methods("PUT")
	router.HandleFunc("/update/{id}", updateStudent).Methods("PUT")

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}
