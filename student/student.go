package student

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"myapp/response"
	"myapp/storage"
	"myapp/student/Type"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func New(s storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var student Type.Student
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

func GetStudent(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		slog.Info("Getting Student with ID  give me a moment...", slog.String("ID", id))
		Id, _ := strconv.Atoi(id)
		studentdata, err := s.GetStudent(int64(Id))

		if err != nil {
			slog.Error("Error fecthing student", slog.String("ID", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		slog.Info("Successfully fecthed Student", slog.String("ID", id))
		response.WriteJson(w, http.StatusAccepted, studentdata)

	}

}

func ListAllStudents(s storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("LIsting All th students in the database")
		students, err := s.ListAllStudents()

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		slog.Info("Successfully Listed all the students from Database")
		response.WriteJson(w, http.StatusOK, students)
	}
}

func UpdateStudent(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student Type.Student
		json.NewDecoder(r.Body).Decode(&student)

		id := r.PathValue("id")
		Id, _ := strconv.Atoi(id)

		_, err := s.UpdateStudent(int64(Id), student.Name, student.Age, student.Email)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		student.Id = Id
		slog.Info("Successfully Updated the Student who has", slog.String("Id", id))
		response.WriteJson(w, http.StatusOK, student)

	}

}
func DeleteStudent(s storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		Id, _ := strconv.Atoi(id)
		slog.Info("Deleting the Student with ID", slog.String("ID", id))
		Sid, err := s.DeleteStudent(int64(Id))
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		slog.Info("Successfully Deleted the Student with ID", slog.String("ID", id))
		response.WriteJson(w, http.StatusOK, map[string]int64{"DeleteID": Sid})

	}
}
