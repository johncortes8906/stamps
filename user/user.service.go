package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

func handleUser(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", userPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		user, err := getUser(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPut:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if *user.ID != userID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(&user)
		err = updateUser(user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case http.MethodDelete:
		err := removeUser(userID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//SetupRoutes sets the routes used by user package
func SetupRoutes(apiBasePath string) {
	usersHandler := http.HandlerFunc(handleUsers)
	userHandler := http.HandlerFunc(handleUser)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, userPath), cors.Middleware(usersHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, userPath), cors.Middleware(userHandler))
}
