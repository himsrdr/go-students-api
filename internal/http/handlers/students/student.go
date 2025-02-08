package students

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/himsrdr/students-api/internal/storage"
	types "github.com/himsrdr/students-api/internal/type"
	"github.com/himsrdr/students-api/internal/utils/response"
)

func Get(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		rowId, err := strconv.ParseInt(id, 10, 10)
		if err != nil {
			fmt.Println(err)
			return
		}
		student, err := storage.GetStudentById(rowId)
		if err != nil {
			fmt.Println(err)
		}
		response.WriteJson(w, http.StatusAccepted, student)
	}

}

func Create(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		student := types.Student{}

		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return

		}

		err = validator.New().Struct(student)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}

		id, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}

		//w.Write([]byte("welcome to project"))
		response.WriteJson(w, http.StatusAccepted, map[string]int64{"id ": id})

	}

}
