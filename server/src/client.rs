use gotham::handler::IntoResponse;
use gotham::state::State;
use hyper::Response;
use gotham::http::response::create_response;
use hyper::StatusCode;
use diesel::mysql::MysqlConnection;
use schema::clients;
use diesel::query_dsl::RunQueryDsl;

#[derive(Queryable, Serialize)]
pub struct Client {
    pub id: i32,
    pub pseudo: String,
    pub password: String,
    pub email: String,
}

impl IntoResponse for Client {
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

impl Client {
    pub fn get_all(connection: &MysqlConnection) -> Vec<Client> {
        clients::table.load::<Client>(connection).unwrap()
    }
}