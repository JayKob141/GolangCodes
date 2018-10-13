package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type EntryPoint interface {
	SetInput([]byte) error
	GetMethod() ([]byte, error)
	PostMethod() ([]byte, error)
	MethodsAllowed() []string
}

func wrapHandler(entryPoint EntryPoint) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		var overallError error

		AtExitPrintIfAnyError := func() {
			if overallError != nil {
				fmt.Println(overallError.Error())
				http.Error(response, overallError.Error(), 500)
			}
		}
		defer AtExitPrintIfAnyError()

		method := request.Method
		methods := entryPoint.MethodsAllowed()
		allowed := false

		for _, value := range methods {
			if value == method {
				allowed = true
				break
			}
		}

		if !allowed {
			overallError = errors.New("Method not specified by entrypoint")
			return
		}

		bodyInBytes, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if err != nil {
			overallError = err
			return
		}

		err = entryPoint.SetInput(bodyInBytes)
		if err != nil {
			overallError = err
			return
		}

		var someBytes []byte
		switch method {
		case "GET":
			someBytes, err = entryPoint.GetMethod()
		case "POST":
			someBytes, err = entryPoint.PostMethod()
		default:
			overallError = errors.New("Method not defined on entrypoint")
			return
		}

		if err != nil {
			overallError = err
			return
		}

		response.Header().Set("Content-type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(someBytes)
	}
}

func main() {
	http.HandleFunc("/v1/example", wrapHandler(EntryPointExample{}))
	fmt.Println("Running server on: http//localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
