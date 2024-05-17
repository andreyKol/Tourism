CREATE TABLE IF NOT EXISTS policy_numbers
(
    id       bigserial primary key,
    number   text not null,
    filename text not null
);
