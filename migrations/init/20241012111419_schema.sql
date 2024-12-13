-- +goose Up
-- +goose StatementBegin
CREATE TABLE "User"(
    uuid UUID PRIMARY KEY ,
    name varchar,
    login varchar UNIQUE,
    email varchar UNIQUE ,
    password varchar,
    status int,
    date_registration timestamp,
);

CREATE TABLE "Status"(
    id serial PRIMARY KEY ,
    name varchar UNIQUE
);

CREATE TABLE "Token" (
    id serial PRIMARY KEY ,
    user UUID,
    access_token TEXT NOT NULL UNIQUE,
    refresh_token TEXT NOT NULL UNIQUE,
    access_token_expiration timestamp,
    refresh_token_expiration timestamp
) ;

ALTER TABLE "User"
    ADD CONSTRAINT "fk_User_0" FOREIGN KEY ("status") REFERENCES "Status" ("id");

ALTER TABLE "Token"
    ADD CONSTRAINT "fk_Token_0" FOREIGN KEY ("user")  REFERENCES "User" ("id") ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "User" DROP CONSTRAINT "fk_User_0" CASCADE ;
ALTER TABLE "Token" DROP CONSTRAINT "fk_Token_0" CASCADE ;

DROP TABLE "User" cascade ;
DROP TABLE "Token" cascade ;
DROP TABLE "Status" cascade ;
-- +goose StatementEnd
