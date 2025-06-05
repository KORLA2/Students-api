package student

import (
	"encoding/json"
	"log/slog"
	"myapp/response"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Student struct {
	id    int
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required"`
}

func New() http.HandlerFunc {

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

		slog.Info("Creating Student")

		response.WriteJson(w, http.StatusCreated, student)

	}

}
