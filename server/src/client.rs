use diesel;
use diesel::prelude::*;
use diesel::mysql::MysqlConnection;
use super::schema::clients;

#[derive(Serialize, Deserialize, Queryable)]
pub struct Client {
    pub id: i32,
    pub pseudo: String,
    pub password: String,
    pub email: String
    // pub friends: Vec<Client>,
    // pub ips: Vec<String>
}

 impl Client {
     pub fn create(client: Client, connection: &MysqlConnection) -> Client {
         diesel::insert_into(clients::table)
             .values(&InsertableClient::from_client(client))
             .execute(connection);
         Post::belonging_to(client).load(&connection).unwrap()
     }

     pub fn read(connection: &MysqlConnection) -> Vec<Client> {
         clients::table
             .order(clients::id.asc())
             .load::<Client>(connection)
             .unwrap()
     }

     pub fn update(id: i32, client: Client, connection: &MysqlConnection) -> Client {
         diesel::update(clients::table.find(id))
             .set(&client)
             .execute(connection);
         Post::belonging_to(client)
             .load(&connection)
             .unwrap()

     }

     pub fn delete(id: i32, connection: &MysqlConnection) -> usize {
         diesel::delete(clients::table.find(id))
             .execute(connection)
             .unwrap()
     }
 }

// #[derive(Serialize, Deserialize, Insertable)]
// #[table_name = "clients"]
// pub struct InsertableClient {
//     pub pseudo: String,
//     pub password: String,
//     pub email: String
// }
//
// impl InsertableClient {
//     fn from_client(client: Client) -> InsertableClient {
//         InsertableClient {
//             pseudo: client.pseudo,
//             password: client.password,
//             email: client.email
//         }
//     }
// }
