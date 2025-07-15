// @title Account management API
// @version 1.0
// @description This API can create and search account of users
// @host localhost:8080
// @BasePath /api/v1

package main

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"myproject.local/authorize/api/account"
	_ "myproject.local/authorize/docs"
	"net/http"
	"os"
)

func main() {

	file, err := os.OpenFile("meu_diario.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error when opening the file:", err)
	}

	log.SetOutput(file)

	account.ConnectDB("postgres://postgres:123456@localhost:5432/account?sslmode=disable")

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/authentication", account.AuthenticationAccount).Methods("POST")
	router.HandleFunc("/api/v1/signup", account.SignupAccount).Methods("POST")
	router.HandleFunc("/api/v1/account/{id}", account.SearchAccountByID).Methods("GET")
	router.HandleFunc("/api/v1/account", account.CreateAccount).Methods("POST")
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	log.Println("Server climbing on the door 8080!")
	log.Fatal(http.ListenAndServe(":8080", router))

}
