CREATE TABLE IF NOT EXISTS portfolios
(
    id       bigserial primary key,
    number   text not null,
    filename text not null
);
