//! An introduction to the Gotham web framework's `IntoResponse` trait.
extern crate diesel;
extern crate server;

use diesel::prelude::*;
use server::establish_connection;
use server::schema::clients::dsl::*;
use server::client::Client;
extern crate futures;
extern crate gotham;
extern crate hyper;
extern crate mime;
extern crate serde;

use hyper::{Response, StatusCode};

use gotham::http::response::create_response;
use gotham::router::Router;
use gotham::router::builder::*;
use gotham::state::State;
use gotham::handler::IntoResponse;

/// Function to handle the `GET` requests coming to `/products/t-shirt`
///
/// Note that this function returns a `(State, Product)` instead of the usual `(State, Response)`.
/// As we've implemented `IntoResponse` above Gotham will correctly handle this and call our
/// `into_response` method when appropriate.
fn get_product_handler(state: State) -> (State, Client) {
    let product = Client {
        id: 2,
        pseudo: "t-shirt".to_string(),
        password: "t-shirt".to_string(),
        email: "t-shirt".to_string(),
    };

    (state, product)
}

/// Create a `Router`
///
/// /products/t-shirt            --> GET
fn router() -> Router {
    build_simple_router(|route| {
        route.get("/client").to(get_product_handler);
    })
}


fn main() {
// /
    let connection = establish_connection();
    let results = clients
        .load::<Client>(&connection)
        .expect("Error loading clients");
//
    println!("Displaying {} posts", results.len());
    for post in results {
        println!("{}", post.pseudo);
        println!("-----------\n");
        println!("{}", post.email);
    }

    /// Start a server and use a `Router` to dispatch requests
    let addr = "0.0.0.0:8000";
    println!("Listening for requests at http://{}", addr);
    gotham::start(addr, router())

}

#[cfg(test)]
mod tests {
    use super::*;
    use gotham::test::TestServer;

    #[test]
    fn get_product_response() {
        let test_server = TestServer::new(router()).unwrap();
        let response = test_server
            .client()
            .get("http://localhost/products/t-shirt")
            .perform()
            .unwrap();

        assert_eq!(response.status(), StatusCode::Ok);

        let body = response.read_body().unwrap();
        let expected_product = Product {
            name: "t-shirt".to_string(),
        };
        let expected_body = serde_json::to_string(&expected_product).expect("serialized product");
        assert_eq!(&body[..], expected_body.as_bytes());
    }
}
