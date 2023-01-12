CREATE ROLE root superuser;

CREATE TABLE pods (
    Id varchar(50) NOT NULL,
    Name varchar(100) NOT NULL,
    Namespace varchar(100) NOT NULL,
    Count NUMERIC(12, 2),
    created timestamp
);