package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

var (
	users = []User{{1, "Ivan"}, {2, "Petr"}, {3, ""}}
)

func main() {
	http.HandleFunc("/users", authMiddleWare(loggerMiddleWare(handleUsers)))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func authMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("x-id")
		if userId == "" {
			log.Printf("[%s] %s - error: userId is not provided\n", r.Method, r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", userId)

		r = r.WithContext(ctx)

		next(w, r)
	}
}

func loggerMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idFromCtx := r.Context().Value("id")
		userID, ok := idFromCtx.(string)
		if !ok {
			log.Printf("[%s] %s - error: userId is invalid", r.Method, r.URL)
			return
		}
		log.Printf("[%s] %s by userID %s\n", r.Method, r.URL, userID)
		next(w, r)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

	// w.WriteHeader(http.StatusBadRequest)
	// w.WriteHeader(http.StatusBadRequest)
	// w.Write([]byte("Hello World\n"))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	w.Write(resp)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users = append(users, user)

}
