CREATE TABLE IF NOT EXISTS clients (
    client_id bigint not null references users(id),
    room_id bigint not null references rooms(id)
);


