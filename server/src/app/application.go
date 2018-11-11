package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/controllers"
)

type Application struct {
	Router *mux.Router
}

func (a *Application) Initialize() {
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

func (a *Application) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *Application) InitializeRoutes() {
	a.Router.MethodNotAllowedHandler = FuckOption()
	a.Router.HandleFunc("/api/clients/new", controllers.CreateClient).Methods("POST")
	a.Router.HandleFunc("/api/clients/login", controllers.Authenticate).Methods("POST")
	a.Router.HandleFunc("/api/clients/{id:[0-9]}", controllers.GetClient).Methods("GET")
	a.Router.HandleFunc("/api/clients/{id:[0-9]}", controllers.UpdateClient).Methods("PUT")
	a.Router.HandleFunc("/api/clients/{id:[0-9]}", controllers.DeleteClient).Methods("DELETE")
	a.Router.HandleFunc("/api/clients/{id:[0-9]}/friends", controllers.AddFriend).Methods("POST")
	a.Router.HandleFunc("/api/clients/{id:[0-9]}/logout", controllers.Logout).Methods("PUT")
	a.Router.Use(JwtAuthentication) //attach JWT auth middleware
}

var FuckOption = func() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			log.Print("C'est de la merde")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			return
		}
	})
}
