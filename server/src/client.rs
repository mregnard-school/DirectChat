use diesel;
use diesel::prelude::*;
use diesel::mysql::MysqlConnection;
use super::schema::{clients};

#[table_name = "clients"]
#[derive(Serialize, Deserialize, Queryable, Insertable)]
pub struct Client {
    pub id: Option<i32>,
    pub pseudo: String,
    pub password: String,
    pub email: String
    // pub friends: Vec<Client>,
    // pub ips: Vec<String>
}

impl Client {
    pub fn create(client: Client, connection: &MysqlConnection) -> Client {
        diesel::insert_into(clients::table)
            .values(&client)
            .execute(connection)
            .expect("Error creating new client");

        clients::table.order(clients::id.desc()).first(connection).unwrap()
    }

    pub fn read(connection: &MysqlConnection) -> Vec<Client> {
        clients::table.order(clients::id.asc()).load::<Client>(connection).unwrap()
    }

    pub fn update(id: i32, client: Client, connection: &MysqlConnection) -> bool {
        // diesel::update(clients::table.find(id)).set(&client).execute(connection).is_ok()
        false
    }

    pub fn delete(id: i32, connection: &MysqlConnection) -> bool {
        diesel::delete(clients::table.find(id)).execute(connection).is_ok()
    }
}
