-- +goose Up
create table "user" (
    id serial primary key,
    first_name text not null,
    last_name text not null,
    password text not null,
    phone_number text not null,
    email text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table "user";

