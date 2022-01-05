package users

import (
	"encoding/json"
	"github.com/devemio/go-rest-api/internal/domain/shared"
	models "github.com/devemio/go-rest-api/internal/domain/users"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

var users []models.User

func Get(w http.ResponseWriter, r *http.Request) {
	if len(users) == 0 {
		users = append(users, models.User{
			Id:           1,
			Username:     "UserA",
			EmailAddress: "user-a@gmai.com",
		})

		users = append(users, models.User{
			Id:           2,
			Username:     "UserB",
			EmailAddress: "user-b@gmai.com",
		})
	}

	json.NewEncoder(w).Encode(users)
}

func Find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if strconv.Itoa(user.Id) == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(shared.NewNotFoundDto())
}

func Create(w http.ResponseWriter, r *http.Request) {
	var dto models.UserInDto
	_ = json.NewDecoder(r.Body).Decode(&dto)
	users = append(users, models.User{
		Id:           rand.Intn(100000000),
		Username:     dto.Username,
		EmailAddress: dto.EmailAddress,
	})
	json.NewEncoder(w).Encode(users)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if strconv.Itoa(user.Id) == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
