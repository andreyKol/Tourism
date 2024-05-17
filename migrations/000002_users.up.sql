CREATE TABLE IF NOT EXISTS users
(
    id           bigserial primary key,
    name         text        not null,
    phone        text        not null,
    password_enc text        not null,
    role         int2,
    created_at   timestamptz not null,
    surname      text,
    patronymic   text,
    age          int2,
    gender       int2,
    email        text,
    image_id     text,
    last_online  timestamptz,
    deleted_at   timestamptz
);