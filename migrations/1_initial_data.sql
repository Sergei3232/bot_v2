-- +goose Up
-- +goose StatementBegin
create table users
(
    id SERIAL primary key,
    name varchar not null,
    telegram_id integer not null
);

create table rool
(
    id  SERIAL primary key,
    name varchar not null
);

create table rool_users
(
    id  SERIAL primary key,
    user_id integer not null,
    rool_id integer not null
);

INSERT INTO rool (name)
VALUES ('admin');

INSERT INTO rool (name)
VALUES ('user');

INSERT INTO users (name, telegram_id)
VALUES ( 'MrS','519588080');

INSERT INTO rool_users (user_id, rool_id)
VALUES ('1', '1');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
drop table if exists rool;
drop table if exists rool_users;
-- +goose StatementEnd