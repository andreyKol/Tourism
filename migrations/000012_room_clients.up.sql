CREATE TABLE IF NOT EXISTS room_clients (
    room_id   bigint,
    client_id bigint,
    PRIMARY KEY (room_id, client_id),
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);
