package main

import (
	coreController "dates/Modules/Core/Controller"
	userEntity "dates/Modules/User/Entity"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {

	user := userEntity.User{}
	fmt.Println(user)

	mux := http.NewServeMux()
	mux.HandleFunc("/", coreController.Hello)
	mux.HandleFunc("/string", coreController.StringExamples)
	mux.HandleFunc("/register", coreController.Register)

	date := time.Now()
	date = date.Add(-3 * 24 * time.Hour)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("Hello world", date.Format("2006-01-02 15:04:05"))
	http.ListenAndServe(":8080", mux)
}
