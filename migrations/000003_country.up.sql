CREATE TABLE IF NOT EXISTS countries
(
    id           bigserial primary key,
    name         text   not null,
    description  text   not null
);