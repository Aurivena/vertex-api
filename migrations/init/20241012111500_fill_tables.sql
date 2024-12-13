-- +goose Up
-- +goose StatementBegin
INSERT INTO "Status" ("name")
    VALUES ('ADMIN'),
           ('SUPER_ADMIN'),
           ('USER');

INSERT INTO "User"("uuid","login","name","email","password","status","date_registration")
    VALUES('0193be23-3373-7120-9d5a-dfbd91d4272b','service','Алукард','service@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',1,current_timestamp),
          ('0193be23-7d7b-7b26-b3d6-35c47ce3cf36','wakadoo','wakadoo','wakadoo@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',2,current_timestamp),
          ('0193be23-b9be-7d79-8910-d2d6f26be8a3','minecraft','minecraft','minecraft@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',3,current_timestamp),
          ('0193be23-d069-78f5-9f81-7fea2b5eeaa6','berserk','berserk','berserk@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',1,current_timestamp),
          ('0193be23-f045-70eb-a3c8-20d402b6b645','huews','huews','huews@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',2,current_timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin


DELETE FROM "User" WHERE login in (
     'service',
     'wakadoo',
     'minecraft',
     'berserk',
     'huews'
    );

DELETE FROM "Status" WHERE name in (
     'ADMIN',
     'SUPER_ADMIN',
     'USER'
    );


-- +goose StatementEnd
