package main

import (
	"myapp/student"
	"net/http"
)

// "github.com/go-playground/validator/v10"

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/student", student.New())

	// if err := validator.New().Struct(URL); err != nil {

	// 	response.ValidationError(err.(validator.ValidationErrors))

	// }

	http.ListenAndServe(":8080", mux)

}
