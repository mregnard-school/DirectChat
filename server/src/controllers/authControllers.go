package controllers

import (
	"encoding/json"
	"log"
	"server/models"
	u "server/utils"
	"net/http"
)

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
	log.Printf("client before: %s", client)
	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	log.Printf("client after decoding: %s", client)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	/**
	client: &{
	{%!s(uint=0)
	0001-01-01 00:00:00 +0000 UTC
	0001-01-01 00:00:00 +0000 UTC
	<nil>}
	batman
	batman
	[]
	[] }

	 */

	resp := models.Login(client.Pseudo, client.Password)
	u.Respond(w, resp)
}

