CREATE TABLE IF NOT EXISTS clients (
    id bigserial primary key,
    client_id text not null,
    room_id bigint not null references rooms(id)
);


