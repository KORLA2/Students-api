package student

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"myapp/response"
	"myapp/storage"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Student struct {
	Id    int
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required"`
}

func New(s storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var student Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))

			return
		}

		lastid, _ := s.CreateStudent(student.Name, student.Age, student.Email)

		slog.Info("Student Created Successfully and ", slog.String("StudentID", fmt.Sprint(lastid)))

		student.Id = int(lastid)
		response.WriteJson(w, http.StatusCreated, student)

	}

}
