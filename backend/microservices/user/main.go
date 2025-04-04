package main

import (
	"db"
	"log"
	"middlewares"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"

	"user/controllers"
)

func main() {
	godotenv.Load()

	con := db.Connect(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_HOST"))
	rabbitmq, err := amqp.Dial(os.Getenv("RABBITMQ"))
	if err != nil {
		log.Fatalf("Error while connecting to rabbitmq: ", err)
	}
	var ctxMiddleware = middlewares.GetMiddleware(con, rabbitmq)

	http.Handle("/signup", ctxMiddleware(http.HandlerFunc(controllers.SignUp)))
	http.Handle("/signin", ctxMiddleware(http.HandlerFunc(controllers.SignIn)))
    http.Handle("/emailConfirmation", ctxMiddleware(http.HandlerFunc(controllers.EmailConfirmation)))

    if err := http.ListenAndServe(":" + os.Getenv("PORT"), nil); err != nil {
        log.Fatal(err)
    }
}
