-- create sources table
CREATE TABLE if NOT EXISTS sources (
       id serial PRIMARY KEY,
       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
       url VARCHAR(255) NOT NULL,
       provider VARCHAR(255) NOT NULL,
       category VARCHAR(255) NOt NULL
);

-- create news table
CREATE TABLE if NOT EXISTS news (
       id serial PRIMARY KEY,
       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
       url VARCHAR(255) NOT NULL,
       title VARCHAR(255) NOT NULL,
       provider VARCHAR(255) NOT NULL,
       category VARCHAR(255) NOT NULL,
       publish_date TIMESTAMPTZ NOT NULL,
       thumbnail Text
);
