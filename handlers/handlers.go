package handlers

import (
	"blogger/constants"
	"blogger/entity"
	"encoding/json"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post entity.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(constants.SOMETHING_WENT_WRONG))
	}
}
