package app

import (
	"github.com/gorilla/mux"
	"server/controllers"
	"log"
	"net/http"
)

type Application struct {
	Router 	*mux.Router
}

func (a *Application) Initialize() {
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

func (a *Application) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *Application) InitializeRoutes() {
	a.Router.HandleFunc("/api/clients/new", controllers.CreateClient).Methods("POST")
	a.Router.HandleFunc("/api/clients/login", controllers.Authenticate).Methods("POST")
	a.Router.HandleFunc("/api/clients/{id:[0-9]}", controllers.GetClient).Methods("GET")
	a.Router.HandleFunc("/api/clients/{id:[0-9]}", controllers.UpdateClient).Methods("PUT")
	a.Router.HandleFunc("/api/clients/{id:[0-9]}", controllers.DeleteClient).Methods("DELETE")
	a.Router.HandleFunc("/api/clients/{id:[0-9]}/friends", controllers.AddFriend).Methods("POST")
	a.Router.Use(JwtAuthentication) //attach JWT auth middleware
}




