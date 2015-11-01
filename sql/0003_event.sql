CREATE TABLE IF NOT EXISTS event_keyword (
     id serial,
     name varchar(250) NOT NULL,
     lookup tsvector NOT NULL UNIQUE,
     CONSTRAINT event_keyword_pkey PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS event_exhibition (
    id serial,
    url varchar(250) UNIQUE,
    museum_id INTEGER REFERENCES "place_museum" ("id"),
    date_begin date,
    date_end date,
    description text,
    description_vector tsvector,
    title varchar(250),
    lang varchar(10),
    CONSTRAINT event_exhibition_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS event_exhibition_keyword (
    id serial,
    event_id INTEGER NOT NULL REFERENCES "event_exhibition" ("id") ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED,
    keyword_id INTEGER NOT NULL REFERENCES "event_keyword" ("id") ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED,
    UNIQUE(event_id, keyword_id),
    CONSTRAINT event_exhibition_keyword_pkey PRIMARY KEY (id)
);
