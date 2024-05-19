CREATE TABLE IF NOT EXISTS events
(
    id           bigserial primary key,
    name         text   not null,
    description  text   not null,
    country_id   int,
    date         timestamptz,
    FOREIGN KEY (country_id) REFERENCES countries(id)
);

