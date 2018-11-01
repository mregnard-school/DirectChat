package main_test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"server/app"
	"server/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a app.Application

func TestMain(m *testing.M) {
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
		panic(m)
	}
	code := m.Run()

	//dropTables()
	os.Exit(code)
}

func ensureTableExists() bool {
	client := models.GetDB().HasTable(&models.Client{})
	address := models.GetDB().HasTable(&models.Ip{})
	client_client := models.GetDB().HasTable("client_client")
	client_address := models.GetDB().HasTable("client_address")
	return (client && address && client_client && client_address)
}
func dropTables() {
	models.GetDB().DropTable(&models.Client{}, &models.Ip{}, "client_client", "client_address")
}

func clearTable(table string) {
	deletion := fmt.Sprintf("DELETE FROM %s", table)
	updateId := fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", table)
	models.GetDB().Exec(deletion)

	models.GetDB().Exec(updateId)
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

func clearTables() {
	clearTable("clients")
	clearTable("client_address")
	clearTable("client_client")
	clearTable("ips")
}

func TestGetNonExistentClient(t *testing.T) {
	clearTable("clients")
}

func GetToken(id int) string {
	tk := &models.Token{UserId: uint(id)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	return tokenString
}
func TestCreateIp(t *testing.T) {
	clearTable("ips")
	Ip := getSimpleIp("localhost")

	models.GetDB().Create(&Ip)

	db_ip := &models.Ip{}
	err := models.GetDB().Table("ips").Where("id = ?", 1).First(db_ip).Error
	if err != nil {
		t.Errorf("Error when getting ip: '%v'", err)
	}
	if db_ip.Address != Ip.Address {
		t.Errorf("Expected ip: '%s', got '%s'", Ip.Address, db_ip.Address)
	}

}

func getSimpleIp(addr string) models.Ip {
	Ip := models.Ip{
		Address: addr,
	}
	return Ip
}

func getSimpleClient() *models.Client{
	client := &models.Client{
		Pseudo: "test_client",
		Password: "test_password",
	}
	return client
}

func getClientWithIp() *models.Client {
	return set2IpToClient(getSimpleClient(), "127.0.0.1", "localhost")
}

func set2IpToClient(client *models.Client, _ip1 string, _ip2 string) *models.Client{
	ip1 := getSimpleIp(_ip1)
	ip2 := getSimpleIp(_ip2)
	ips := make([]models.Ip, 2)
	ips[0] = ip1
	ips[1] = ip2
	client.Ips = ips
	return client
}

func setFriendInClient(client *models.Client) (*models.Client,*models.Client,*models.Client){
	_friend1 := getSimpleClient()
	_friend1.Create()
	_friend2 := getSimpleClient()
	_friend2.Pseudo = "test_client_2"
	_friend2.Create()
	friend,_ := models.GetClient(uint(1))
	friend2,_ := models.GetClient(uint(2))
	friends := make([]models.Client, 2)
	friends[0] = *friend
	friends[1] = *friend2
	client.Friends = friends
	return client, friend, friend2
}

func generateFriendToNewClient(client *models.Client, nbFriends int) (*models.Client){
	friends := make([]models.Client, nbFriends)
	for i := 0; i < nbFriends; i ++ {
		tmp_f := getSimpleClient()
		tmp_f.Pseudo = fmt.Sprintf("%s_%d", tmp_f.Pseudo, i)
		tmp_f.Create()
		db_friend, _ := models.GetClient(uint(i + 1))
		friends[i] = *db_friend
	}
	client.Friends = friends
	return client
}


func getClientWithFriends() (*models.Client,*models.Client,*models.Client) {
	client := getSimpleClient()
	client.Pseudo = "client_with_Friend"
	return setFriendInClient(client)
}


func TestCreateSimpleClient(t *testing.T) {
	clearTable("clients")
	clearTable("ips")
	clearTable("client_address")
	client :=  getSimpleClient()
	returnClient, err := client.Create()

	if err != nil {
		t.Errorf("Error when creating the client: %s", err)
	}
	if returnClient.Password != "" {
		t.Errorf("The password should be empty, instead got '%s'", returnClient.Password)
	}
	clientFromDB, _ := models.GetClient(1)
	compareClient(client, clientFromDB, t)
}

func compareClient(client *models.Client, clientFromDB *models.Client, t *testing.T) {
	if clientFromDB == nil {
		t.Error("Client empty")
		return
	}
	if clientFromDB.Pseudo != client.Pseudo {
		t.Errorf("The pseudo expected was '%s', got '%s'", client.Pseudo, clientFromDB.Pseudo)
	}
	err := bcrypt.CompareHashAndPassword([]byte(clientFromDB.Password), []byte(client.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		t.Errorf("The password expected was '%s', got '%s'", client.Password, clientFromDB.Password)
	}
	if lenCliDb:=len(clientFromDB.Ips); lenCliDb != len(client.Ips) {
		t.Errorf("Not the same amount of ips. Expected : '%d', got: '%d", len(client.Ips), lenCliDb)
		return
	}
	for i:=0; i < len(clientFromDB.Ips); i++ {
		if tmp := clientFromDB.Ips[i].Address; tmp != client.Ips[i].Address {
			t.Errorf("Not the same address at index %d. Expected: '%s', got '%s'", i, client.Ips[i].Address, tmp)
		}
	}
}

func TestCreateClientWithIp(t *testing.T){
	clearTable("clients")
	clearTable("ips")
	clearTable("client_address")
	client := getClientWithIp()
	_, err := client.Create()
	if err != nil {
		t.Errorf("Error when creating client: '%s'", err)
	}
	clientFromDb, _ := models.GetClient(1)
	compareClient(client, clientFromDb, t)

}

func TestCreateClientWithFriends(t *testing.T) {
	clearTables()
	client, f_friend, s_friend := getClientWithFriends()
	client.Create()

	compareClient2Friends(client, t, f_friend, s_friend)
}

func compareClient2Friends(client *models.Client, t *testing.T, f_friend *models.Client, s_friend *models.Client) {
	dbClient,_ := models.GetClient(3)
	clients := []models.Client{}
	models.GetDB().Find(&clients)
	if dbClient == nil {
		t.Error("Client is empty")
		return
	}
	compareClient(client, dbClient, t)
	friends := dbClient.Friends
	if l := len(friends); l != 2 {
		t.Error("Client is supposed to have 2 friends, instead had '%s'", l)
	}
	compareClient(f_friend, &friends[0], t)
	compareClient(s_friend, &friends[1], t)
}

func TestUpdateSimpleClient(t *testing.T) {
	clearTable("clients")
	client := getSimpleClient()
	_, err := client.Create()
	if err != nil {
		t.Errorf("Error when creating client: '%s'", err)
	}
	client.Pseudo = "updatePseudo"
	client.Update()
	clients := []models.Client{}
	models.GetDB().Find(&clients)
	if l := len(clients); l != 1 {
		t.Errorf("Expected 1 client, got '%d'", l)
		return
	}
	clientFromDB := &clients[0]
	compareClient(client, clientFromDB, t)
}

func TestUpdateClientWithIp(t *testing.T) {
	clearTables()
	client := getClientWithIp()
	_, err := client.Create()
	if err != nil {
		t.Errorf("Error when creating client: '%s'", err)
	}
	client.Pseudo = "updatePseudo"
	ip2 := getSimpleIp("new addresss")
	client.Ips[1] = ip2
	client.Update()
	clients := []models.Client{}
	models.GetDB().Find(&clients)
	if l := len(clients); l != 1 {
		t.Errorf("Expected 1 client, got '%d'", l)
		return
	}
	clientFromDB,_ := models.GetClient(1)
	compareClient(client, clientFromDB, t)
}

func TestUpdateComplexFriend(t *testing.T) {
	clearTables()
	_client, f_friend, s_friend := getClientWithFriends()
	client := set2IpToClient(_client, "localhost", "localhost_2")
	insertedClient := addClient(client)
	compareClient2Friends(insertedClient, t, f_friend, s_friend)
	toUpdate := set2IpToClient(_client, "new address 1", "new address 2")
	toUpdate.Update()
}

func compareClientWithFriends(idClient int, client *models.Client, friends []models.Client, t *testing.T) {
	dbClient,_ := models.GetClient(uint(idClient))
	clients := []models.Client{}
	models.GetDB().Find(&clients)
	if dbClient == nil {
		t.Error("Client is empty")
		return
	}
	compareClient(client, dbClient, t)
	dbFriends := dbClient.Friends
	if l := len(dbFriends); l != len(friends){
		t.Error("Client is supposed to have '%d' friends, instead had '%d'", len(friends), l)
	}
	for i := 0; i < len(friends); i++ {
		compareClient(&friends[i], &dbFriends[i], t)
	}
}

func TestAddFriend(t *testing.T) {
	clearTables()
	_client := getSimpleClient()
	_client = generateFriendToNewClient(_client, 3)
	_client.Create()
	friends :=  _client.Friends
	new_friend :=  getSimpleClient()
	new_friend.Pseudo = "new friend"
	new_friend.Create()
	_client.AddFriend(*new_friend)
	friends = append(friends, *new_friend)
	compareClientWithFriends(4, _client, friends, t)
}

func addClient(client *models.Client) *models.Client{
	r, _ := client.Create()
	result, _ := models.GetClient(r.ID)
	return result
}