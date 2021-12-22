-- create sources table
CREATE TABLE if NOT EXISTS sources (
       id serial PRIMARY KEY,
       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
       url VARCHAR(255) NOT NULL,
       source_name VARCHAR(255) NOT NULL
);

-- create news table
CREATE TABLE if NOT EXISTS news (
       id serial PRIMARY KEY,
       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
       url VARCHAR(255) NOT NULL,
       title VARCHAR(255) NOT NULL,
       rss_source VARCHAR(255) NOT NULL,
       publish_date TIMESTAMPTZ NOT NULL
);
