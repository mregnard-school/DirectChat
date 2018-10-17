#![feature(plugin)]
#![plugin(rocket_codegen)]

extern crate rocket;
#[macro_use] extern crate rocket_contrib;
#[macro_use] extern crate serde_derive;

use rocket_contrib::{Json, Value};

// mod client;
// use client::{Client};

// #[post("/", data = "<client>")]
// fn create(client: Json<Client>) -> Json<Client> {
//     client
// }
//
// #[get("/")]
// fn read() -> Json<Value> {
//     Json(json!([
//         "client 1",
//         "client 2"
//     ]))
// }
//
// #[put("/<id>", data = "<client>")]
// fn update(id: i32, client: Json<Client>) -> Json<Client> {
//     client
// }
//
// #[delete("/<id>")]
// fn delete(id: i32) -> Json<Value> {
//     Json(json!({"status": "ok"}))
// }
#[get("/<name>/<age>")]
fn hello(name: String, age: u8) -> String {
    format!("Hello, {} year old named {}!", age, name)
}
fn main() {
    rocket::ignite()
        // .mount("/client", routes![create, update, delete])
        // .mount("/clientes", routes![read])
        .mount("/hello", routes![hello])
        .launch();
}
