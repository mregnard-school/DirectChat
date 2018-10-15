const date = new Date();

module.exports = [
  {
    "client": {
      "id": 1,
      "pseudo": "billy",
      "ips": [
        "127.0.0.1",
        //"127.0.0.1",
      ]
    },
    "queue": [
      {
        "message": {
          "id": 1,
          "chat_id": 1,
          "author_id": 1,
          "date": date,
          "content": "Hello",
        }
      }
    ]
  }
];