CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    age INT,
    score FLOAT,
    active BOOLEAN,
    created_at TIMESTAMP
);

INSERT INTO users(name,age, score, active, created_at) values ('Voja',NULL, NULL, NULL, NULL);
