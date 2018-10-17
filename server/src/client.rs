#[derive(Serialize, Deserialize)]
pub struct Client {
    pub id: Option<i32>,
    pub pseudo: String,
    pub email: String,
    pub friends: Vec<Client>,
    pub ips: Vec<String>
}

impl Client {
    //fonction ajout ami
    //fonction suppression ami
    //fonction create
    //fonction update
    //fonction delete
}
