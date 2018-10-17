#[derive(Serialize, Deserialize)]
pub struct Client {
    pub id: Option<i32>,
    pub pseudo: String,
    pub email: String
}
