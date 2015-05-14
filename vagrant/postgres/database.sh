#!/bin/bash

dropdb --if-exists -h localhost -p 5432 -U takasing -e toyo
dropuser -h localhost -p 5432 -U postgres -e takasing
createuser -P -d -U postgres -e takasing
createdb -h localhost -p 5432 -U takasing -E UTF8 -e toyo

sh -c 'exec psql -h localhost -p 5432 -U takasing -d toyo' <<EOF
DROP TABLE IF EXISTS articles;
CREATE TABLE articles ( \
  id SERIAL PRIMARY KEY, \
  title varchar, \
  content text, \
  created_at timestamp \
  );
DROP TABLE IF EXISTS users;
CREATE TABLE users ( \
  id SERIAL PRIMARY KEY, \
  name varchar, \
  created_at timestamp \
  );
INSERT INTO users (name, created_at) values ('toyo', now());
EOF
