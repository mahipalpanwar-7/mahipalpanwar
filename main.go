package main

import (
	"encoding/json"
	"net/http"
)

// model for course - file

type Course struct{
 CourseId string `json:"courseid"`
 CourseName string `json:"coursename"`
 CoursePrise string `json:"price"`
 Author *Author `json:"author"`
}


type Author struct{
Fullname string `json:"fullname"`
Website string `json:"website"`
}

// fake Database 

var courses []Course

// middleware

func (c *Course) IsEmpty() bool{
	return c.CourseId=="" && c.CourseName==""
}

// controllers 

// home route 

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>welcome to the Course"))
}

func getAllCourses(w  http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func main(){

}