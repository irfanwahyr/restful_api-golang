package main

import (
	"fmt"
	"latihan_golang/config"
	"latihan_golang/routes"
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()
	r := mux.NewRouter() // membuat router
	routes.RouteIndex(r)

	log.Println("server run at", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r) //create server local

}
