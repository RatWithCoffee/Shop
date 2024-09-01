package user

import (
	"encoding/json"
	"fmt"
	session2 "graphql/internal/session"
	"io/ioutil"
	"net/http"
)

type UserHandler struct {
	Sessions  session2.SessionManagerInterface
	UsersRepo *UserRepo
}

type UserBody struct {
	User User `json:"user"`
}

func (userHandler *UserHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var userBody UserBody
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	err := json.Unmarshal(body, &userBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := userBody.User
	if !isValidUser(user) {
		http.Error(w, fmt.Sprintf("empty fields are not allowed"), http.StatusBadRequest)
		return
	}

	err = userHandler.UsersRepo.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userHandler.Sessions.Create(w, user)
}
