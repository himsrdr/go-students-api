package students

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	types "github.com/himsrdr/students-api/internal/type"
	"github.com/himsrdr/students-api/internal/utils/response"
)

func Create() http.HandlerFunc {
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

		//w.Write([]byte("welcome to project"))
		response.WriteJson(w, http.StatusAccepted, map[string]string{"Status": "success"})

	}

}
