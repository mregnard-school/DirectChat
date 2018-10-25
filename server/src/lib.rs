#[macro_use]
extern crate diesel;
extern crate dotenv;
extern crate gotham;
#[macro_use]
extern crate gotham_derive;
extern crate hyper;
#[macro_use]
extern crate serde_derive;
extern crate serde_json;
extern crate core;

use diesel::prelude::*;
use dotenv::dotenv;
use std::env;

pub mod client;
pub mod schema;

#[derive(Deserialize, StateData, StaticResponseExtender)]
pub struct IdPath {
    pub id: i32,
}

pub fn get_database<'a>() -> MysqlConnection {
    establish_connection()
}

static mut db: Option<&MysqlConnection> = None;

pub fn establish_connection() -> MysqlConnection{
    dotenv().ok();

    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let connection = MysqlConnection::establish(&database_url)
        .unwrap_or_else(|_| panic!("Error connecting to {}", database_url));

   connection
}
