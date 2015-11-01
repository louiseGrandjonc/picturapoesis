--
-- INSERT/UPDATE document_fulltext
--
CREATE OR REPLACE FUNCTION description_vector_trg() RETURNS TRIGGER AS $BODY$
DECLARE eventlang text;
BEGIN
  SELECT lang_conversion(NEW.lang) INTO eventlang;
  NEW.description_vector = to_tsvector(eventlang::regconfig, NEW.description);
  RETURN NEW;
END;
$BODY$ LANGUAGE plpgsql;


--
-- Declaration des triggers
--
DROP TRIGGER IF EXISTS event_exhibition_description_vector_trigger ON event_exhibition;

--
CREATE TRIGGER event_exhibition_description_vector_trigger
    BEFORE INSERT OR UPDATE ON event_exhibition FOR EACH ROW EXECUTE PROCEDURE description_vector_trg();    
