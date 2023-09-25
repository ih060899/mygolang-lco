package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	// Seeding data
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "ReactJS",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Imran Hossain",
			Website:  "lco.dev",
		},
	})

	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "MERN Stack",
		CoursePrice: 199,
		Author: &Author{
			Fullname: "Imran Hossain",
			Website:  "go.dev",
		},
	})

	// routhing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a part
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "applicatioan/json")
	_ = json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			err := json.NewEncoder(w).Encode(course)
			if err != nil {
				panic(err)
			}
			return
		}
	}
	_ = json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// what if: body is empty
	if r.Body == nil {
		_ = json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		_ = json.NewEncoder(w).Encode("No data inside JSON")

		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	for _, c := range courses {
		if c.CourseName == course.CourseName {
			_ = json.NewEncoder(w).Encode("Course is already exists")
			return
		}

	}

	// generate unique id, string
	// append course into courses

	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	_ = json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id ,remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			_ = json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: send a response when id is no found
	_ = json.NewEncoder(w).Encode("Course not found")

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove (index, index + 1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			_ = json.NewEncoder(w).Encode("Course Deleted")
			return
		}
	}
	_ = json.NewEncoder(w).Encode("Course not found")
}
