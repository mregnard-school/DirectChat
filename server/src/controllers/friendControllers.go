package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/models"
	u "server/utils"
	"strconv"
)

var AddFriend = func(w http.ResponseWriter, r *http.Request) {

	friend := &models.Client{}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	client, _ := models.GetClient(uint(id))


	err = json.NewDecoder(r.Body).Decode(friend) //decode the request body into struct and failed if any error occur
	json.NewDecoder(r.Body).Decode(friend)       //decode the request body into struct and failed if any error occur
	if err != nil {
		log.Print(err)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	friend, err = models.GetClientFromPseudo(friend)
	if err != nil {
		log.Print(err)
		u.Respond(w, u.Message(false, "Invalid pseudo"))
	}

	resp := client.AddFriend(friend)
	u.Respond(w, resp)
}