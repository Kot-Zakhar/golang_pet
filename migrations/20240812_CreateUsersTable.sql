CREATE TABLE IF NOT EXISTS Users (
    Id INT NOT NULL PRIMARY KEY,
    Name VARCHAR,
    Login VARCHAR,
    Password VARCHAR,
    CreatedAt TIMESTAMT NOT NULL
)