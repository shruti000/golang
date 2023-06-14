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

// The name mux stands for "HTTP request multiplexer". Like the standard http.
//
//	ServeMux, mux. Router matches incoming requests against a list of registered routes and calls a
//	 handler for the route that matches the URL or other conditions.
//
// model for courses
type Courses struct {
	courseId    string  `json:"courseid"`
	courseName  string  `json:"coursenname"`
	coursePrice int     `json:"Price"`
	author      *Author `json:"author"`
}

// model fro author
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// we create fake databse ad slice
var courses []Courses

// helper methods just like some condiitons method calld as middleware like loggers,exceptions
func (c *Courses) IsEmpty() bool {
	return c.courseId == "" && c.courseName == ""
}

func main() {

	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	//routing
	//hanflefucntion are baccillt used for the toirng
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//for fort to use listenANdServe
	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// is the home page method
// every method has 2 parametrs complusory (w http.ResponseWriter,r *http.Request)
// the frst emthod is the writer method to write  and 2nd method is the read whcih is a request emthod to request the body
// write write in byte format so we use []byte
// it has reposne writer which is an interfcae which inbuilf has 3 method Header() http.Header which
// cintains the applciation nd json file,Write([]byte) it will wte (int, error)
// WriteHeader(statusCode int)
// the requwst is used o genertae the requs
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to API Building </h1>"))
}

// getall method to get all the coirses
// NewWncoder
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The ALL Courses are :")
	//before using thta weneed to set it to appciion.json
	w.Header().Set("Content-Type", "application/json")
	//JSON encoding of any type and encodes/writes it any writable
	//bacial;y encod eis to write
	json.NewEncoder(w).Encode(courses)
}

// get sungle course
func getSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Single Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	//it can be done using mux Vars which takes the request
	params := mux.Vars(r)
	fmt.Println(params)

	//after tha value we get on params we will check wether the id there and in databse i.e slice matches
	for _, value := range courses {
		if value.courseId == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with the id ")

}

// add the course
// basically to add we need first the json data
// fir add course we are tking the coirse from the baclend
// like we are decodng the json
// here we use request as request is sending the data
func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add the Courses")
	w.Header().Set("Content-Type", "application/json")

	//what if bosy is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about the data is 0 {}
	var course Courses
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.courseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// upadte the course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating the course")
	w.Header().Set("Content-Type", "application/json")
	//mux gives access to vars which helps to grab the conent from the request
	params := mux.Vars(r)

	var course Courses
	//loop throught the values
	//in update we will find the id gievn the we will remobve that json then add out new json with the same id\
	for index, value := range courses {
		if value.courseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.courseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}
