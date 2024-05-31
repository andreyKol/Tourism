CREATE TABLE IF NOT EXISTS messages (
    id   bigserial primary key,
    content text not null,
    room_id bigint not null references rooms(id),
    client_id bigint not null references users(id),
    created_at   timestamptz not null
);
