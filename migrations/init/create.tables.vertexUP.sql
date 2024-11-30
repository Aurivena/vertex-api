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


ALTER TABLE "User"
    ADD CONSTRAINT "fk_User_0" FOREIGN KEY ("status") REFERENCES "Status" ("id")
