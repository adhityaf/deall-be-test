DROP TABLE IF EXISTS users;
CREATE SEQUENCE USER_ID START 1;
CREATE TABLE users (
    USER_ID serial PRIMARY KEY,
    NAME VARCHAR(256) NOT NULL,
    USERNAME VARCHAR(256) UNIQUE NOT NULL,
    PASSWORD TEXT NOT NULL,
    ROLE VARCHAR(5) NOT NULL,
    CREATED_AT TIMESTAMPTZ DEFAULT Now(),
    UPDATED_AT TIMESTAMPTZ DEFAULT Now() 
);

INSERT INTO "users" ("username","password","name","role","created_at","updated_at") 
VALUES ('aditadmin','$2a$08$/nXfxNCYSHDbriSPy65Taugs2DRFyKNKgxe2EhJFFZ5.Nl54ivPua','Adhitya Admin','admin',NOW(),NOW());

INSERT INTO "users" ("username","password","name","role","created_at","updated_at") 
VALUES ('adhityaf','$2a$08$/nXfxNCYSHDbriSPy65Taugs2DRFyKNKgxe2EhJFFZ5.Nl54ivPua','Adhitya Febhiakbar','user',NOW(),NOW());