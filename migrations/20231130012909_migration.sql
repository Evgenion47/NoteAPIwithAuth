-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table "notes"
(
  "noteid" serial PRIMARY KEY,
  "ownerid" VARCHAR(32) NOT NULL,
  "title" VARCHAR(255) NOT NULL,
  "description" VARCHAR(255) NOT NULL,
  "completion" boolean default false,
  "createdat" timestamp default now()
);
create table "users"
(
  "userid" varchar(32) primary key,
  "password" varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists "notes";
drop table if exists "users";
-- +goose StatementEnd
