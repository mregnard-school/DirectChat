use gotham::handler::IntoResponse;
use gotham::state::State;
use hyper::Response;
use gotham::http::response::create_response;
use hyper::StatusCode;

#[derive(Queryable, Serialize)]
pub struct ClientIp {
    pub id: i32,
    pub client_id: i32,
    pub ip: String,
}

/// Implements `gotham::handler::IntoResponse` trait for `Product`
///
/// `IntoResponse` represents a type which can be converted to a response. This trait is used in
/// converting the return type of a function into a response.
///
/// This trait implementation uses the Serde project when generating responses. You don't need to
/// know about Serde in order to understand the response that is being created here but if you're
/// interested you can learn more at `http://serde.rs`.
impl IntoResponse for ClientIp {
    fn into_response(self, state: &State) -> Response {
        create_response(
            state,
            StatusCode::Ok,
            Some((
                serde_json::to_string(&self)
                    .expect("serialized product")
                    .into_bytes(),
                mime::APPLICATION_JSON,
            )),
        )
    }
}
