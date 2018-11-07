package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"server/app"
	"server/models"
	"os"
	"testing"
)

var a app.Application
var NbClient int

func TestMain(m *testing.M) {
	log.Print("Running test !")
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	a = app.Application{}
	a.Initialize()

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("test_db_name")
	dbHost := os.Getenv("db_host")

	models.Open(username, password, dbName, dbHost)
	if !ensureTableExists() {
		fmt.Print("La db n'existe pas ")
		panic(m)
	}
	NbClient = 1
	code := m.Run()

	dropTables()
	os.Exit(code)
}

func ensureTableExists() bool {
	client := models.GetDB().HasTable(&models.Client{})
	address := models.GetDB().HasTable(&models.Ip{})
	clientClient := models.GetDB().HasTable(&models.Friendship{})
	clientAddress := models.GetDB().HasTable("client_address")
	return client && address && clientClient && clientAddress
}
func dropTables() {
	models.GetDB().DropTable(&models.Client{}, &models.Ip{}, &models.Friendship{}, "client_client", "client_address")
}

func checkResponseCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(request *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, request)
	return rr
}

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
		bufferClient := clientToBuffer(t, getSimpleClient())
		req, _ := http.NewRequest("POST", "/api/clients/new", bufferClient)
		resp := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, resp.Code)
		client := &models.Client{}
		json.NewDecoder(resp.Body).Decode(client)
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

func clientToBuffer(t *testing.T, client *models.Client) *bytes.Buffer {
	payload, err := json.Marshal(client)
	if err != nil{
		t.Errorf("error occurs when encoding client: %s", err.Error())
	}
	return bytes.NewBuffer(payload)
}

func UseClient(client *models.Client, t *testing.T) *models.Client {
	payload, err := json.Marshal(client)
	if err != nil{
		t.Errorf("error occurs when encoding client: %s", err.Error())
	}
	req, _ := http.NewRequest("POST", "/api/clients/new", bytes.NewBuffer(payload))
	resp := executeRequest(req)
	log.Print(resp)
	c := models.Client{}
	json.NewDecoder(resp.Body).Decode(c)
	log.Print(resp)
	return client
}