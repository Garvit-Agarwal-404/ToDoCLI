package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

// constructor function-> to initialize a struct using a function
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{
		Filename: fileName,
	}

}

// save function ->pointer reciever function
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Filename, fileData, 0644) // what is 0644
}

// load func-> pointer reiever function

func (s *Storage[T]) load(data *T) error {
	fileData, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
