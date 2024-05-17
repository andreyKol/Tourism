CREATE TABLE IF NOT EXISTS patients
(
    id           bigserial primary key,
    name         text        not null,
    phone        text        not null,
    password_enc text        not null,
    created_at   timestamptz not null,
    surname      text,
    patronymic   text,
    age          int2,
    gender       int2,
    email        text,
    image_id     int8 references images (id),
    last_online  timestamptz,
    deleted_at   timestamptz,
    policy_number int8 references policy_numbers (id),
    medical_card int8 references medical_cards (id)
);