package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"server/src/models"
	u "server/src/utils"
	"log"
	"net/http"
	"strconv"
)

var UpdateClient = func(w http.ResponseWriter, r *http.Request) {

	client := &models.Client{}


	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	//@TODO need to find a way to iterate through every struct value
	clientFromDb, errs := models.GetClient(uint(id))
	log.Print(clientFromDb, errs)
	//@TODO get the client and replace values by changed value
	//@TODO save friend in pivot table
	err = json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	log.Print(client)
	json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	log.Print(client)
	if err != nil {
		log.Print(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := client.Update()
	u.Respond(w, resp)
}

var DeleteClient = func(w http.ResponseWriter, r *http.Request) {

	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := client.Delete()
	u.Respond(w, resp)
}

var GetClient = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	client, err := models.GetClient(uint(id))
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			u.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	u.RespondWithJSON(w, http.StatusOK, client)
}

