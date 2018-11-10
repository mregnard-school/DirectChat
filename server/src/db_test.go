package main

import (
	"fmt"
	"log"
	"server/models"
	"testing"
)




func addClient(client *models.Client) *models.Client{
	r, _ := client.Create()
	result, _ := models.GetClient(r.ID)
	return result
}

func getSimpleIp(addr string) *models.Ip {
	Ip := models.Ip{
		Address: addr,
	}
	return &Ip
}

func getClientWithIp() *models.Client {
	return set2IpToClient(getSimpleClient(), "127.0.0.1", "localhost")
}

func set2IpToClient(client *models.Client, _ip1 string, _ip2 string) *models.Client{
	ip1 := getSimpleIp(_ip1)
	ip2 := getSimpleIp(_ip2)
	ips := make([]*models.Ip, 2)
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
	friends := make([]*models.Client, 2)
	friends[0] = friend
	friends[1] = friend2
	client.Friends = friends
	return client, friend, friend2
}
//
//func getFriendship(client *models.Client, friend *models.Client) *models.Friendship {
//	return &models.Friendship{
//		Friend:friend,
//	}
//}

func generateFriendToNewClient(client *models.Client, nbFriends int) (*models.Client){
	friends := make([]*models.Client, nbFriends)
	for i := 0; i < nbFriends; i ++ {
		tmp_f := getSimpleClient()
		tmp_f.Pseudo = fmt.Sprintf("%s_%d", tmp_f.Pseudo, i)
		tmp_f.Create()
		db_friend, _ := models.GetClient(uint(i + 1))
		friends[i] = db_friend
	}
	client.Friends = friends
	return client
}


func getClientWithFriends() (*models.Client,*models.Client,*models.Client) {
	client := getSimpleClient()
	client.Pseudo = "client_with_Friend"
	return setFriendInClient(client)
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
		t.Errorf("Client is supposed to have 2 friends, instead had '%d'", l)
	}
	compareClient(f_friend, friends[0], t)
	compareClient(s_friend, friends[1], t)
}



func TestGetNonExistentClient(t *testing.T) {
	clearTable("clients")
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

func TestCreateSimpleClient(t *testing.T) {
	clearTable("clients")
	clearTable("ips")
	clearTable("client_address")
	client :=  getSimpleClient()
	log.Print("Avant create")
	returnClient, err := client.Create()
	log.Print("Après create")

	if err != nil {
		t.Errorf("Error when creating the client: %s", err)
	}
	if returnClient.Password != "" {
		t.Errorf("The password should be empty, instead got '%s'", returnClient.Password)
	}
	clientFromDB, _ := models.GetClient(1)
	log.Print("Avant compare client")
	compareClient(client, clientFromDB, t)
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

func TestAddFriend(t *testing.T) {
	clearTables()
	_client := getSimpleClient()
	_client = generateFriendToNewClient(_client, 3)
	_client.Create()
	friends :=  _client.Friends
	newFriend :=  getSimpleClient()
	newFriend.Pseudo = "new friend"
	newFriend.Create()
	client, err := _client.AddFriend(newFriend)
	if err != nil {
		t.Errorf("Eror when adding a friend %v", err)
		return
	}
	friends = append(friends, newFriend)
	compareClientWithFriends(4, client, friends, t)
	for i:= 0; i < len(client.Friendships); i++ {
		if client.Friendships[i].Accepted {
			t.Errorf("This friendship %v with the friend %v is not supposed to be accepted for this client : %v", client.Friendships[i], client.Friends[i], client)
		}
	}
}

func TestMutualFriendShip(t *testing.T) {
	clearTables()
	client := addSimpleClient(t, "localhost")
	friend := addSimpleClient(t, "friend_ip")
	friend, err  := friend.AddFriend(client)
	if	friend.Friendships[0].Accepted {
		t.Errorf("Friend was not add but is accepted %v", friend.Friends[0])
	}
	if err != nil {
		t.Errorf("Error when addding friend %v", err)
		return
	}
	log.Print("avant le add friend")
	client, err = client.AddFriend(friend)
	log.Print("Après le add friend")
	friend, err = models.GetClient(2)
	log.Print("Après le get client")

	if err != nil {
		t.Errorf("Error when addding friend %v", err)
		return
	}
	if	!client.Friendships[0].Accepted {
		t.Errorf("Friend was add but isn't accepted %v", client.Friends[0])
	}
	if	!friend.Friendships[0].Accepted {
		t.Errorf("Friend was add but isn't accepted %v", friend.Friends[0])
	}
}

func TestGetClientByPseudo(t *testing.T) {

}