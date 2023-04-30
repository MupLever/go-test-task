package main

import (
	"log"

	"github.com/MupLever/go-test-task/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
