CREATE TABLE IF NOT EXISTS portfolios
(
    id       bigserial primary key,
    experience   text not null,
    filename text not null
);
