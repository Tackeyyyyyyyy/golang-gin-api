create database name_list;


-- auto-generated definition
CREATE TABLE names
(
  id         INT AUTO_INCREMENT
    PRIMARY KEY,
  first_name VARCHAR(40) NULL,
  last_name  VARCHAR(40) NULL
)
  ENGINE = InnoDB;

INSERT INTO name_list.names (id, first_name, last_name) VALUES (1, 'test', 'hoge');
INSERT INTO name_list.names (id, first_name, last_name) VALUES (2, 'test2', 'hoge2');
INSERT INTO name_list.names (id, first_name, last_name) VALUES (3, 'test3', 'hoge3');
INSERT INTO name_list.names (id, first_name, last_name) VALUES (4, 'test4', 'hoge4');
