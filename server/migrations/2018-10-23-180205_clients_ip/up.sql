CREATE TABLE client_ip (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  client_id INT(11) NOT NULL,
  CONSTRAINT client_id
    FOREIGN KEY (client_id)
    REFERENCES clients (id),
  ip VARCHAR(20)
)
