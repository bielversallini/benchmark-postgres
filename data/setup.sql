DROP SCHEMA IF EXISTS bench CASCADE;

CREATE SCHEMA bench;

CREATE EXTENSION "pgcrypto";

CREATE TABLE bench.sample_one (
  id UUID NOT NULL DEFAULT gen_random_uuid(),
  name varchar(255) default NULL,
  email varchar(255) default NULL,
  country varchar(100) default NULL,
  CONSTRAINT id_pk PRIMARY KEY (id)
);

CREATE TABLE bench.sample_two (
  id SERIAL,
  name varchar(255) default NULL,
  email varchar(255) default NULL,
  country varchar(100) default NULL,
  CONSTRAINT id_pkk PRIMARY KEY (id)
);

CREATE INDEX ON bench.sample_two (country);