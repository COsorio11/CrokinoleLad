-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE individuals(id SERIAL UNIQUE, name varchar(50) NOT NULL, wins int, losses int, password varchar(50) NOT NULL);

INSERT INTO individuals (name, password) VALUES ('Gojira', 'D3stroyToky0');
INSERT INTO individuals (name, password) VALUES ('Ryhorn', 'first created');
INSERT INTO individuals (name, password) VALUES ('Agumon', 'Tai');


CREATE TABLE teams(id SERIAL UNIQUE, name varchar(50) NOT NULL, wins int, losses int, password text NOT NULL, playerone varchar(50) NOT NULL, playertwo varchar(50) NOT NULL);

INSERT INTO teams (name, password, playerone, playertwo) VALUES ('There Goes Tokyo', 'D3stroyToky0', 'Mothra', 'King Ghidorah');

CREATE TABLE games(id SERIAL UNIQUE, playerone int NOT NULL, playertwo int NOT NULL, winner int NOT NULL);

INSERT INTO games (playerone, playertwo, winner) VALUES (1, 2, 1);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE individuals;
DROP TABLE teams;
DROP TABLE games;
