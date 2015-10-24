CREATE TABLE IF NOT EXISTS event_keyword (
     id serial,
     name varchar(250) NOT NULL UNIQUE
     CONTRAINT event_hashtags_pkey PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS event_exhibition (
    id serial,
    url varchar(250) UNIQUE,
    museum_id INTEGER REFERENCES "place_museum" ("id"),
    date_begin time,
    date_end time,
    description text,
    title varchar(250),
    CONTRAINT event_event_pkey PRIMARY KEY (id)
)
CREATE TABLE IF NOT EXISTS event_exhibition_keyword (
    id serial,
    event_id INTEGER NOT NULL REFERENCES "event_exhibition" ("id") ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED,
    hashtag_id INTEGER NOT NULL REFERENCES "event_keyword" ("id") ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED,
    UNIQUE(event_id, hashtag_id)
    CONTRAINT event_event_hashtags_pkey PRIMARY KEY (id)
);
