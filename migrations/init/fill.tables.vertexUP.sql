INSERT INTO "Status" ("name")
    VALUES ('ADMIN'),
           ('SUPER_ADMIN'),
           ('USER');

INSERT INTO "User"("login","name","email","password","status","date_registration")
    VALUES('service','Алукард','service@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',1,current_timestamp),
          ('wakadoo','wakadoo','wakadoo@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',2,current_timestamp),
          ('minecraft','minecraft','minecraft@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',3,current_timestamp),
          ('berserk','berserk','berserk@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',1,current_timestamp),
          ('huews','huews','huews@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',2,current_timestamp);
