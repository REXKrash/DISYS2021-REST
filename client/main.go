package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
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
	student := Student{
		Id:   5,
		Name: "Bob",
	}
	addStudent(student)
}

func getCourse(id int) {
	type Response struct {
		Message string `json:"message"`
	}

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
		fmt.Println(course.Name)
	}
	defer resp.Body.Close()
}

func addCourse(data Course) {
	type Response struct {
		Message string `json:"message"`
	}

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
	type Response struct {
		Message string `json:"message"`
	}

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
