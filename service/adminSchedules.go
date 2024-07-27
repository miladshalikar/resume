package service

import (
	"database/sql"
	"net/http"
	"project1/dto"
	"project1/utills"
)

func AdminSchedules(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.InvalidMethod})
			return
		}

		//------------------------------------------

	}
}
