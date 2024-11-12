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

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func (c *Course) IsDuplicate() bool {
	for _, course := range courses {
		if course.CourseName == c.CourseName && course.Author.Fullname == c.Author.Fullname {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Hello World")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "reactjs.dev"}})
	courses = append(courses, Course{CourseId: "5", CourseName: "Golang", CoursePrice: 400, Author: &Author{Fullname: "Abimbola Ronald", Website: "go.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "mernstack.dev"}})
	courses = append(courses, Course{CourseId: "6", CourseName: "Django", CoursePrice: 300, Author: &Author{Fullname: "Abimbola Ronald", Website: "djago.dev"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/api/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/api/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/api/course", createOneCourse).Methods("POST")
	r.HandleFunc("/api/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/api/course/{id}", deleteOneCouurse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers - file

// Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Ronald Abimbola A.K.A StarBoy</h1>"))
}

// Get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// get id from request
	params := mux.Vars(r)

	// loop through courses and find id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with CourseId: " + params["id"])
	return
}

//`{ success: true, message: "Successfully logged in." }`

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// What if : Body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// What if: Body is {}

	var course Course

	// Decode the body into course struct
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send valid JSON data")
		return
	}

	if course.IsDuplicate() {
		json.NewEncoder(w).Encode("Course already exists")
		return
	}

	// generate unique id, string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	// get id from request
	params := mux.Vars(r)

	// loop through courses and find id, remove element, and add updated element back

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse)
			course.CourseId = params["id"]
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("{ success: false, message: 'ID NOT FOUND' }")
}

func deleteOneCouurse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("{ success: true, message: 'Record deleted successfully' }")
			return
		}
	}
	json.NewEncoder(w).Encode("{ success: false, message: 'ID NOT FOUND' }")
}
