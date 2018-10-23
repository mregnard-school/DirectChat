#![feature(plugin)]
#![plugin(rocket_codegen)]

//For db
#[macro_use]
extern crate diesel;
extern crate r2d2;
extern crate r2d2_diesel;

extern crate serde_derive;
extern crate serde;
extern crate dotenv;

extern crate rocket;
extern crate rocket_contrib;

mod db;
pub mod schema;

use rocket_contrib::{Json, Value};

mod client;
use self::client::{Client};

#[post("/", data = "<client>")]
fn create(client: Json<Client>, connection: db::Connection) -> Json<Client> {
    let insert = Client { id: None, ..client.into_inner() };
    Json(Client::create(insert, &connection))
}

#[get("/")]
fn read(connection: db::Connection) -> Json<Value> {
    Json(json!(Client::read(&connection)))
}

#[put("/<_id>", data = "<client>")]
fn update(_id: i32, client: Json<Client>, connection: db::Connection) -> Json<Value> {
    let update = Client { id: _id, ..client.into_inner() };
    Json(json!({
        "success": Client::update(id, update, &connection)
    }))
}

#[delete("/<id>")]
fn delete(id: i32, connection: db::Connection) -> Json<Value> {
    Json(json!({
        "success": Client::delete(id, &connection)
    }))
}

#[get("/<name>/<age>")]
fn hello(name: String, age: u8) -> String {
    format!("Hello, {} year old named {}!", age, name)
}
fn main() {

    rocket::ignite()
        .manage(db::connect())
        .mount("/client", routes![create, update, delete])
        .mount("/clients", routes![read])
        .mount("/hello", routes![hello])
        .launch();
}
