package main

import (
	"encoding/json"
	"errors"
)

type EntryPointExample struct {
}

type EntryPointResponse struct {
	Message string `json:"something"`
}

func (entry EntryPointExample) SetInput(bytes []byte) error {
	return nil
}

func (entry EntryPointExample) MethodsAllowed() []string {
	return []string{"GET"}
}

func (entry EntryPointExample) PostMethod() ([]byte, error) {
	return []byte{0x00}, errors.New("Method POST not implemented")
}
func (entry EntryPointExample) GetMethod() ([]byte, error) {
	var response = EntryPointResponse{Message: "Hello world api"}

	return json.Marshal(response)
}
