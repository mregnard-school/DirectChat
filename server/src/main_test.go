package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"server/app"
	"server/models"
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

func clearTable(table string) {
	deletion := fmt.Sprintf("DELETE FROM %s", table)
	updateId := fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", table)
	models.GetDB().Exec(deletion)
	models.GetDB().Exec(updateId)
}

func clearTables() {
	clearTable("clients")
	clearTable("client_address")
	clearTable("friendships")
	clearTable("ips")
}

func getSimpleClient() *models.Client{
	pseudo := fmt.Sprintf("test_client_%d", NbClient)
	client := &models.Client{
		Pseudo: pseudo,
		Password: "test_password",
	}
	NbClient ++
	return client
}
