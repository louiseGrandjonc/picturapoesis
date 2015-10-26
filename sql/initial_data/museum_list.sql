DELETE FROM place_museum;
DELETE FROM place_place;
DELETE FROM place_address;


INSERT INTO place_address (
    id,
    zip_code,
    address,
    city)
VALUES (
    1,
    75116,
    '11 avenue du Président Wilson',
    'Paris'),(
    2,
    75008,
    '3 Avenue du Général Eisenhower',
    'Paris'
    ),(
    3,
    75004,
    'Place Georges-Pompidou',
    'Paris'
    );

INSERT INTO place_place (
    id,
    address_id,
    infos,
    url,
    agenda_url,
    name,
    last_watched)
VALUES (
    1,
    1,
    'TODO',
    'http://www.mam.paris.fr',
    'http://www.mam.paris.fr/fr/expositions',
    'Musée Art Moderne',
    now()),
    (
    2,
    2,
    'Grand palais',
    'http://www.grandpalais.fr',
    'http://www.grandpalais.fr/fr/programmation?state=now&type[0]=14&term_node_tid_depth[0]=1575',
    'Grand palais',
    now()),
    (
    3,
    3,
    'Centre Pompidou',
    'https://www.centrepompidou.fr',
    'https://www.centrepompidou.fr/cpv/agenda/expositions',
    'Centre Pompidou',
    now());


INSERT INTO place_museum (
   id,
   schedule,
   place_id,
   exhibition_regex,
   lang)
VALUES (
    1,
    '{{"10:00:00", "18:00:00"}, {NULL, NULL}, {"10:00:00", "18:00:00"},  {"10:00:00", "18:00:00"},  {"10:00:00", "22:00:00"},  {"10:00:00", "18:00:00"},{"10:00:00", "18:00:00"}}',
    1,
    '"\/fr\/expositions\/exposition-([^"]*)"',
    'fr'),
    (
    2,
    '{{"10:00:00", "22:00:00"}, {"10:00:00", "20:00:00"}, {NULL, NULL}, {"10:00:00", "22:00:00"},  {"10:00:00", "22:00:00"},  {"10:00:00", "22:00:00"},{"10:00:00", "22:00:00"}}',
    2,
    '"\/fr\/evenement\/([^"]*)"',
    'fr'
    ),
    (
    3,
    '{{"11:00:00", "22:00:00"}, {"11:00:00", "22:00:00"}, {NULL, NULL}, {"11:00:00", "22:00:00"},  {"11:00:00", "22:00:00"},  {"11:00:00", "22:00:00"},{"11:00:00", "22:00:00"}}',
    3,
    '"\/cpv\/ressource\.action\?param\.id=FR_R([^"]*)param\.seance=seance"',
    'fr'
    );
