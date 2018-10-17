-- Your SQL goes here
CREATE TABLE friend_binding (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  client_id INT(11) NOT NULL,
  friend_id INT(11) NOT NULL,
  FOREIGN KEY (client_id) REFERENCES clients(id),
  FOREIGN KEY (friend_id) REFERENCES clients(id)
)
