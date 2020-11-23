CREATE KEYSPACE toto
    WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE TABLE toto.main (
    id int,
    value double,
    PRIMARY KEY (id)
);
