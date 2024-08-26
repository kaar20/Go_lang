package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Temporory Db
var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to the Building api Sections")
	r := mux.NewRouter()
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "Flutter",
		CoursePrice: 188,
		Author: &Author{
			FullName: "Mohamed Abdiaziz",
			Website:  "Kaar.dev",
		},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Golang",
		CoursePrice: 188,
		Author: &Author{
			FullName: "Mohamed Abdiaziz",
			Website:  "go.dev",
		},
	})
	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course", newCourse).Methods("POST")

	// Listen
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controller

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> Welcome to the First controller of Go User from GetX User </h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You will Get all Courses in Your Mind")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You will Get Wanted Course if its Availible")
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(courses)
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return

		}
	}
	json.NewEncoder(w).Encode("No Found for this Id course")
	return
}

func newCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You will Create Wanted Course ")
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(courses)

	if r.Body == nil {
		json.NewEncoder(w).Encode("Not Allowed Null Value")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Passed Empty values for the Json")
		return
	}

	for _, courseName := range courses {
		if courseName.CourseName == courseName.CourseName {
			json.NewEncoder(w).Encode("This Course Already Registered")
			return

		}
	}

	// generate random id for Now\
	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}
