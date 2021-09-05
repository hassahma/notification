package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/robfig/cron/v3"
	_ "github.com/marvel/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/marvel/controller"
	"github.com/marvel/cache"
)

func routing() {
	routes := mux.NewRouter().StrictSlash(true)
	routes.HandleFunc("/characters", controller.GetAllCharacters)
	routes.HandleFunc("/characters/{id}", controller.GetCharacterById)
	// Swagger
	routes.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":9091", routes))
}

func init() {
	strategy := flag.String("s", "TTL", "Cache strategy (TTL or PREFETCH)")
	flag.Parse()

	fmt.Println("Caching Strategy is ", *strategy)
	cache.Init()
	if *strategy != "TTL" && *strategy != "PREFETCH" {
		panic(fmt.Sprintf("\n#######################################################################################\n" +
			"Invalid cache strategy '%s'. Supported cache strategies are TTL or PREFETCH" +
			"\n#######################################################################################\n", *strategy))
	}

	if *strategy == "PREFETCH" {
		c := cron.New(cron.WithSeconds())
		c.AddFunc("@every 15m", controller.InvalidateAndRefresh)
		c.Start()
	}
}

// @title Marvel Characters API
// @version 1.0
// @description This is an api for querying the Marvel characters.
// @termsOfService http://swagger.io/terms/

// @contact.name Dr Ahmad Hassan
// @contact.url https://www.linkedin.com/in/ahmadhassan
// @contact.email ahmad.hassan@gmail.com
func main() {
	fmt.Println("\n\n########################################\n" +
		"Welcome to Marvel characters!" +
		"\n########################################")
	routing()
}
