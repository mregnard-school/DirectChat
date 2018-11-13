package controllers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/services"
	u "server/utils"
)

/**
400 -> route qui existe pas
422-UnprocessableEntity-> données au mauvaais format
500-> problème serveur
*/

var CreateClient = func(w http.ResponseWriter, r *http.Request) {

	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", http.StatusUnprocessableEntity))
		return
	}
	if resp, ok := services.Validate(client); !ok {
		u.Respond(w, resp)
		return
	}

	newClient, err := client.Create()
	if err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	u.RespondWithJSON(w, http.StatusCreated, newClient)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.RespondWithError(w, http.StatusUnauthorized, "Wrong Formatting")
		return
	}
	client, code, message := services.Login(client)
	if code != http.StatusOK {
		u.RespondWithError(w, code, message)
		return
	}

	u.RespondWithJSON(w, code, client)
}
