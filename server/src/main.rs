//! An introduction to the Gotham web framework's `IntoResponse` trait.
extern crate diesel;
extern crate futures;
extern crate gotham;
extern crate hyper;
extern crate mime;
extern crate serde;
extern crate server;

use diesel::prelude::*;
use gotham::handler::IntoResponse;
use gotham::http::response::create_response;
use gotham::router::builder::*;
use gotham::router::Router;
use gotham::state::State;
use gotham::state::FromState;
use hyper::{Response, StatusCode};
use server::establish_connection;
use server::get_database;
use server::IdPath;
use server::client::Client;


fn get_product_handler(mut state: State) -> (State, Response) {

    let IdPath { id } = IdPath::take_from(&mut state);
    print!("Called with id");
    //let clients = Client::get_all(&get_database());
    let clients:Vec<Client> = Vec::new();
    let json = serde_json::to_string(&clients).unwrap();
    let response = create_response(
        &state,
        StatusCode::Ok,
        Some((json.into_bytes(), mime::APPLICATION_JSON)),
    );
    (state, response)
}

fn router() -> Router {
    build_simple_router(|route| {
        route.get("/clients/:id/")
            .with_path_extractor::<IdPath>()
            .to(get_product_handler);
    })
}

fn main() {
    let connection = establish_connection();
    /// Start a server and use a `Router` to dispatch requests
    let addr = "0.0.0.0:8000";
    println!("Listening for requests at http://{}", addr);
    gotham::start(addr, router())

}

#[cfg(test)]
mod tests {
    use gotham::test::TestServer;
    use super::*;

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
