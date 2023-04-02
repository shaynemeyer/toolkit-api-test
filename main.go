package main

import (
	"log"
	"net/http"

	"github.com/shaynemeyer/toolkit/v2"
)

func main() {
	mux := routes()

	log.Println("Starting application on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	
	mux.HandleFunc("/api/login", login)
	mux.HandleFunc("/api/logout", logout)

	return mux
}

func login(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	var payload struct {
		Email string `json:"username"`
		Password string `json:"password"`
	}

	err := tools.ReadJSON(w, r, &payload)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}
	
	var respPayload toolkit.JSONResponse

	if payload.Email == "me@here.com" && payload.Password == "verysecret" {
		respPayload.Error = false
		respPayload.Message = "Logged in"
		_ = tools.WriteJSON(w, http.StatusAccepted, respPayload)
		return
	}

	respPayload.Error = true
	respPayload.Message = "invalid credentials"
	_ = tools.WriteJSON(w, http.StatusUnauthorized, respPayload)
}

func logout(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	payload := toolkit.JSONResponse{
		Message: "Logged out",
	}

	_ = tools.WriteJSON(w, http.StatusAccepted, payload)
}