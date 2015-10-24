CREATE TABLE IF NOT EXISTS place_address(
       id serial,
       zip_code INTEGER NOT NULL DEFAULT 75000,
       address varchar(250),
       city varchar(60),
       CONSTRAINT place_address_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS place_place (
    id serial,
    address_id INTEGER REFERENCES "place_address" ("id"),
    infos text,
    url varchar(250) NOT NULL,
    agenda_url varchar(250) NOT NULL,
    last_watch timestamp,
    image varchar(250),
    name varchar(250),
    CONSTRAINT place_place_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS place_museum (
    id serial,
    schedule time[][],
    place_id INTEGER REFERENCES "place_place" ("id") ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED,
    exhibition_regex varchar(250),
    tag_regex varchar(250),
    CONSTRAINT place_museum_pkey PRIMARY KEY (id)
);
