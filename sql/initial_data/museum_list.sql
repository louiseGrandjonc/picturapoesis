INSERT INTO place_address (
    id,
    zip_code,
    address,
    city)
VALUES (
    1,
    75116,
    '11 avenue du Président Wilson',
    'Paris');

INSERT INTO place_place (
    id,
    address_id,
    infos,
    url,
    agenda_url,
    name,
    last_watched,
    default_lang)
VALUES (
    1,
    1,
    'TODO',
    'http://www.mam.paris.fr',
    'http://www.mam.paris.fr/fr/expositions',
    'Musée Art Moderne',
    now());


DELETE FROM place_museum;

INSERT INTO place_museum (
   id,
   schedule,
   place_id,
   exhibition_regex,
   default_lang)
VALUES (
    1,
    '{{"10:00:00", "18:00:00"}, {NULL, NULL}, {"10:00:00", "18:00:00"},  {"10:00:00", "18:00:00"},  {"10:00:00", "22:00:00"},  {"10:00:00", "18:00:00"},{"10:00:00", "18:00:00"}}',
    1,
    '"\/fr\/expositions\/exposition-([^"]*)"',
    'fr');
