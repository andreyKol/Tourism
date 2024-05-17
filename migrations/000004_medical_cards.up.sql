CREATE TABLE IF NOT EXISTS medical_cards
(
    id       bigserial primary key,
    number   text not null,
    filename text not null
);
