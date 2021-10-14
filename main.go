// main package that sets up the api routes and initalizes the application.
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"github.com/robfig/cron/v3"
	_ "github.com/hassahma/notification/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/hassahma/notification/controller"
	"github.com/hassahma/notification/constant"
)

// sets up routing for api controllers and swagger ui.
func routing() {
	routes := mux.NewRouter().StrictSlash(true)
	routes.HandleFunc("/notification/subscribe", controller.PostNotify).Methods("POST")
	routes.HandleFunc("/notification/test", controller.PostNotifyTest).Methods("POST")
	routes.HandleFunc("/notification/activate", controller.PostNotifyActivate).Methods("POST")

	// Swagger
	routes.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":" + constant.PORT, routes))
}

// @title Notification API
// @version 1.0
// @description This is an api for notification service.
// @termsOfService http://swagger.io/terms/

// @contact.name Dr Ahmad Hassan
// @contact.url https://www.linkedin.com/in/ahmadhassan
// @contact.email ahmad.hassan@gmail.com
func main() {
	fmt.Printf("\n\n#####################################################################################\n" +
		"Started Notification service successfully on http://localhost:%s/swagger/index.html" +
		"\n#####################################################################################\n\n", constant.PORT)
	routing()
}
