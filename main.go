package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kooparse/ameyokocho/controllers"
	"github.com/kooparse/ameyokocho/models"
	"log"
	"net/http"
)

func main() {
	models.ConnectDB()
	models.SetupDB()

	router := httprouter.New()

	router.GET("/", controllers.Index)
	router.GET("/isbn/:isbn", controllers.CreateBookFromIsbn)

	log.Println("Server started")
	http.ListenAndServe(":3333", router)
}
