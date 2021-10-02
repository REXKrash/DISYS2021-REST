package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	runLocalhost()
}

func runLocalhost() {
	type Student struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
	type Response struct {
		Message string `json:"message"`
	}
	data := Student{
		Id:   5,
		Name: "Bob",
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
