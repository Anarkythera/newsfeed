-- create sources table
CREATE TABLE if NOT EXISTS sources (
       id serial PRIMARY KEY,
       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
       url VARCHAR(255) NOT NULL,
       source_name VARCHAR(255) NOT NULL
);
