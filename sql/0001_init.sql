--
-- Conversions
--
CREATE EXTENSION IF NOT EXISTS hstore;
-- Convert small country code to well known postgresql configuration
CREATE OR REPLACE  FUNCTION lang_conversion(lang text) RETURNS text AS $BODY$
BEGIN
    RETURN 'fr=>french, en=>english'::hstore -> lang;
END;
$BODY$ LANGUAGE plpgsql;
