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
	"github.com/marvel/utils"
	"github.com/marvel/constant"
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
	utils.Init()
	utils.Strategy = flag.String("s", "PREFETCH", "Cache strategy (TTL or PREFETCH)")
	flag.Parse()

	cache.Init()
	if *utils.Strategy == "TTL" {
		fmt.Printf("Caching Strategy is %s. %s\n", *utils.Strategy, constant.TTL)
	} else if *utils.Strategy == "PREFETCH" {
		fmt.Printf("Caching Strategy is %s. %s\n", *utils.Strategy, constant.PREFETCH)
	} else {
		panic(fmt.Sprintf("\n#######################################################################################\n"+
			"Invalid cache strategy '%s'. Supported cache strategies are TTL or PREFETCH. %s %s"+
			"\n#######################################################################################\n", *utils.Strategy, constant.TTL, constant.PREFETCH))
	}

	if *utils.Strategy == "PREFETCH" {
		c := cron.New(cron.WithSeconds())
		fmt.Printf("\nUsing cache prefetch time of %d minute.\n", utils.Cfg.Marvel.PREFETCH)
		c.AddFunc(fmt.Sprintf("@every %dm", utils.Cfg.Marvel.PREFETCH), controller.InvalidateAndRefresh)
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
		"Started Marvel characters API successfully!" +
		"\n########################################")
	routing()
}
