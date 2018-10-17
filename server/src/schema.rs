table! {
    clients (id) {
        id -> Nullable<Integer>,
        pseudo -> Varchar,
        password -> Varchar,
        email -> Varchar,
    }
}

table! {
    client_ip (id) {
        id -> Integer,
        client_id -> Integer,
        ip -> Nullable<Varchar>,
    }
}

table! {
    friend_binding (id) {
        id -> Integer,
        client_id -> Integer,
        friend_id -> Integer,
    }
}

joinable!(client_ip -> clients (client_id));

allow_tables_to_appear_in_same_query!(
    clients,
    client_ip,
    friend_binding,
);
