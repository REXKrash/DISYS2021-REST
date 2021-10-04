package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
}

func main() {

	student := Student{
		Id:   5,
		Name: "Bob",
	}
	addStudent(student)

	course1 := Course{
		Id:   0,
		Name: "DISYS",
	}
	addCourse(course1)

	course2 := Course{
		Id:   1,
		Name: "BDSA",
	}
	addCourse(course2)
	getCourse(1)
	enrollStudent(1, 5)
	getCourse(1)
}

func getCourse(id int) {
	req, err := http.NewRequest("GET", ("http://localhost:8080/v2/course/" + strconv.Itoa(id)), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		r, _ := ioutil.ReadAll(resp.Body)
		var course Course
		json.Unmarshal(r, &course)
		fmt.Printf("%#v", course)
		fmt.Println()
	}
	defer resp.Body.Close()
}

func enrollStudent(courseID int, studentID int) {
	req, err := http.NewRequest("PUT", ("http://localhost:8080/v2/course/" + strconv.Itoa(courseID) + "/enroll/" + strconv.Itoa(studentID)), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		r, _ := ioutil.ReadAll(resp.Body)
		var response Response
		json.Unmarshal(r, &response)
		fmt.Println(response.Message)
	}
	defer resp.Body.Close()
}

func addCourse(data Course) {
	courseBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := bytes.NewReader(courseBytes)

	req, err := http.NewRequest("POST", "http://localhost:8080/v2/course", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		r, _ := ioutil.ReadAll(resp.Body)
		var response Response
		json.Unmarshal(r, &response)
		fmt.Println(response.Message)
	}
	defer resp.Body.Close()
}

func addStudent(data Student) {
	studentBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := bytes.NewReader(studentBytes)

	req, err := http.NewRequest("POST", "http://localhost:8080/v2/student", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		r, _ := ioutil.ReadAll(resp.Body)
		var response Response
		json.Unmarshal(r, &response)
		fmt.Println(response.Message)
	}
	defer resp.Body.Close()
}
