package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/REXKrash/DISYS2021-REST/client/v2/swagger"
	sw "github.com/REXKrash/DISYS2021-REST/client/v2/swagger"
)

func main() {

	runSwaggerLocalhost()
}

func runSwaggerLocalhost() {
	var client = swagger.NewAPIClient(swagger.NewConfiguration())
	var student = swagger.Student{Id: 5, Name: "bob"}
	resp, err := client.StudentApi.AddStudent(context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING"), student)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		r, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(r))
	}
}

func runLocalhost() {
	resp, err := http.Get("http://localhost:8080/v2/")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		r, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(r))
	}
	defer resp.Body.Close()
}
