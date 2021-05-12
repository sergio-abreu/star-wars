CREATE TABLE planets
(
    id   UUID    NOT NULL
        CONSTRAINT planets_pk
            PRIMARY KEY,
    name varchar(50) NOT NULL,
    climates varchar(255) NOT NULL,
    terrains varchar(255) NOT NULL
);
