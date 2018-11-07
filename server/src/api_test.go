package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"server/models"
)

func TestLoginNonExistentUser(t *testing.T) {
	clearTable("clients")
	client := getSimpleClient()
	payload, err := json.Marshal(client)
	if err != nil{
		t.Errorf("error occurs when encoding client: %s", err.Error())
	}
	req, _ := http.NewRequest("POST", "/api/clients/login", bytes.NewBuffer(payload))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusUnauthorized, resp.Code)
}

func TestRegisterClients(t *testing.T) {
	clearTables()
	for i := 0 ; i < 3 ; i ++{
		client := addSimpleClient(t, "localhost")
		if client == nil {
			t.Error("Client is null")
			return
		}
		clientFromDb, err := models.GetClient(client.ID)
		if err != nil {
			t.Errorf("Error getting client; %s", err.Error())
			return
		}
		compareClient(client, clientFromDb, t)
	}
}

func TestLoginExistingClient(t *testing.T) {
	clearTables()
	clientRegisterd := addSimpleClient(t, "localhost")
	log.Printf("client registered : %v", clientRegisterd)
	client := useClient(t, clientRegisterd)
	clientFromDb, err := models.GetClient(client.ID)
	if err != nil {
		t.Errorf("Error getting client; %s", err.Error())
		return
	}
	compareClient(client, clientFromDb, t)
}

func clientToBuffer(t *testing.T, client *models.Client) *bytes.Buffer {
	payload, err := json.Marshal(client)
	if err != nil{
		t.Errorf("error occurs when encoding client: %s", err.Error())
	}
	return bytes.NewBuffer(payload)
}

func addSimpleClient(t *testing.T, ipCLient string) *models.Client{
	req, _ := http.NewRequest("POST", "/api/clients/new", clientToBuffer(t, getSimpleClient()))
	req.RemoteAddr = ipCLient
	//log.Printf("Ip: %s", req.RemoteAddr)
	resp := executeRequest(req)
	c := &models.Client{}
	json.NewDecoder(resp.Body).Decode(c)
	return c
}

func useClient(t *testing.T, client *models.Client) *models.Client {
	c := &models.Client{
		Pseudo: client.Pseudo,
		Password: "test_password",
	}
	log.Printf("client trying to log : %v", c)

	req, _ := http.NewRequest("POST", "/api/clients/login", clientToBuffer(t, c))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
	json.NewDecoder(resp.Body).Decode(client)
	if client == nil {
		t.Error("Client is null")
		return nil
	}
	return client
}
