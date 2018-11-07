package controllers

import (
	"encoding/json"
	"server/models"
	"net/http"
	u "server/utils"
	"server/services"
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
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	newClient, err := client.Create()
	if  err != nil {
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
	client, code, message := services.Login(client.Pseudo, client.Password)
	if code != http.StatusOK {
		u.RespondWithError(w, code, message)
	}

	u.RespondWithJSON(w, code, client)
}

