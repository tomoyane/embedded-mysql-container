-- name: sample
CREATE DATABASE sample;

-- name: sample
USE sample;

-- name: tests
DROP TABLE IF EXISTS tests;

-- name: tests
CREATE TABLE tests (
  id int(11) NOT NULL AUTO_INCREMENT,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
