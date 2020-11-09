package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/johncortes8906/stamps/cors"
)

const userPath = "users"

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userList, err := getUserList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response, err := json.Marshal(userList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(response)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		userID, err := insertUser(user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"userId":%d}`, userID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

//SetupRoutes sets the routes used by user package
func SetupRoutes(apiBasePath string) {
	usersHandler := http.HandlerFunc(handleUsers)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, userPath), cors.Middleware(usersHandler))
}
