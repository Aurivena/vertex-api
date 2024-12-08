CREATE TABLE "User"(
    name varchar,
    login varchar UNIQUE,
    email varchar UNIQUE ,
    password varchar,
    status int,
    date_registration timestamp,
    PRIMARY KEY (login)
) WITHOUT OIDS ;

CREATE TABLE "Status"(
    id serial PRIMARY KEY ,
    name varchar UNIQUE
) WITHOUT OIDS ;

CREATE TABLE "Token" (
    id serial PRIMARY KEY ,
    login varchar,
    access_token TEXT NOT NULL UNIQUE,
    refresh_token TEXT NOT NULL UNIQUE,
    token_expiration timestamp,
    refresh_token_expiration timestamp,
    is_revoked bool default false
) WITHOUT OIDS ;

ALTER TABLE "User"
    ADD CONSTRAINT "fk_User_0" FOREIGN KEY ("status") REFERENCES "Status" ("id");

ALTER TABLE "Token"
    ADD CONSTRAINT "fk_Token_0" FOREIGN KEY ("login")  REFERENCES "User" ("login") ON DELETE CASCADE
