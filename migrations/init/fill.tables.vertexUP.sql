INSERT INTO "Status" ("name")
    VALUES ('ADMIN'),
           ('SUPER_ADMIN'),
           ('USER');

INSERT INTO "User"("login","name","email","password",status)
    VALUES('service','Алукард','service@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',1),
          ('wakadoo','wakadoo','wakadoo@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',2),
          ('minecraft','minecraft','minecraft@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',3),
          ('berserk','berserk','berserk@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',1),
          ('huews','huews','huews@gmail.com','489cd5dbc708c7e541de4d7cd91ce6d0f1613573b7fc5b40d3942ccb9555cf35',2);
