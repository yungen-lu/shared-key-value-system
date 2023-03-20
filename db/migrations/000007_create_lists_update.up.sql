CREATE OR REPLACE FUNCTION update_latest_page_on_insert()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.next_page_key IS NOT NULL AND EXISTS(SELECT 1 FROM pages WHERE key = NEW.next_page_key) THEN
    UPDATE pages SET next_page_key = NEW.key WHERE next_page_key = NEW.next_page_key AND list_key = NEW.list_key AND key <> NEW.key;
    UPDATE lists SET next_page_key = NEW.key WHERE key = NEW.list_key AND next_page_key = NEW.next_page_key;
  ELSE
    UPDATE pages SET next_page_key = NEW.key WHERE list_key = NEW.list_key AND key = (SELECT latest_page_key FROM lists WHERE key = NEW.list_key);
    UPDATE lists SET latest_page_key = NEW.key WHERE key = NEW.list_key;
    UPDATE lists SET next_page_key = NEW.key WHERE key = NEW.list_key AND next_page_key IS NULL;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_latest_page_on_delete()
RETURNS TRIGGER AS $$
BEGIN
  IF pg_trigger_depth() > 1 THEN
    RETURN OLD;
  END IF;
  UPDATE lists SET latest_page_key = (SELECT key FROM pages WHERE list_key = OLD.list_key AND next_page_key = OLD.key LIMIT 1) WHERE key = OLD.list_key AND latest_page_key = OLD.key;
  UPDATE lists SET next_page_key = OLD.next_page_key WHERE key = OLD.list_key AND next_page_key = OLD.key;
  UPDATE pages SET next_page_key = OLD.next_page_key WHERE list_key = OLD.list_key AND next_page_key = OLD.key;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_latest_page_after
AFTER INSERT ON pages
FOR EACH ROW
EXECUTE PROCEDURE update_latest_page_on_insert();
CREATE TRIGGER trigger_latest_page_before
BEFORE DELETE ON pages
FOR EACH ROW
EXECUTE PROCEDURE update_latest_page_on_delete();