package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/13ubbleBoy/REST_APIs_In_Golang/internal/storage"
	"github.com/13ubbleBoy/REST_APIs_In_Golang/internal/types"
	"github.com/13ubbleBoy/REST_APIs_In_Golang/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		// data jo postman se  send kar rhe wo yaha accept kar rhe, data encoded form me aata hai.
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student) // jo bhi json aa raha eske ander decode kar lo
		if errors.Is(err, io.EOF) {
			// returning json response
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"))) // http.StatusBadRequest = 400 http error status code.
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors) // typecasting errors because neeche wale function me wahi type of error accept ho rha hai.
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		slog.Info("user created successfully", slog.String("userID", fmt.Sprint(lastId)))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		// w.Write([]byte("welcome to student api"))
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId}) // http.StatusCreated = 201 http resource created code
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Error("error getting user", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting all students")

		students, err := storage.GetStudents()
		if err != nil { // error == nil means everything OK no error
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusOK, students) // 'w' pe directly write karga "WriteJson" function json me.
	}
}
