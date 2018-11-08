package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"server/models"
	"testing"
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
	client := useClient(t, clientRegisterd, "156.0.1.2")
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

func TestLoginWrongCredential(t *testing.T) {
	clearTables()
	client := addSimpleClient(t, "localhost")
	c := &models.Client{
		Pseudo: client.Pseudo,
		Password: "wrong_password",
	}
	req, _ := http.NewRequest("POST", "/api/clients/login", clientToBuffer(t, c))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusUnauthorized, resp.Code)
	json.NewDecoder(resp.Body).Decode(client)
	if client == nil {
		t.Error("Client is null")
		return
	}
}

func TestAddFriendApi(t *testing.T) {

}

func TestMutualFriendShip(t *testing.T) {

}

func addSimpleClient(t *testing.T, ip string) *models.Client{
	req, _ := http.NewRequest("POST", "/api/clients/new", clientToBuffer(t, getSimpleClient()))
	req.RemoteAddr = ip
	resp := executeRequest(req)
	c := &models.Client{}
	json.NewDecoder(resp.Body).Decode(c)
	return c
}

func useClient(t *testing.T, client *models.Client, ip string) *models.Client {
	c := &models.Client{
		Pseudo: client.Pseudo,
		Password: "test_password",
	}
	req, _ := http.NewRequest("POST", "/api/clients/login", clientToBuffer(t, c))
	req.RemoteAddr= ip
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
	json.NewDecoder(resp.Body).Decode(client)
	if client == nil {
		t.Error("Client is null")
		return nil
	}
	client.Ips = []*models.Ip{
		{
			Address:ip,
		},
	}
	return client
}
