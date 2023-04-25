CREATE OR REPLACE FUNCTION update_latest_page_on_insert()
RETURNS TRIGGER AS $$
BEGIN
  -- insert in the middle of the list
  IF NEW.next_page_key IS NOT NULL AND EXISTS(SELECT 1 FROM pages WHERE key = NEW.next_page_key) THEN
    -- find the page that is pointing to the same page as the new page pointing, because this triggers after the update
    -- exclude the new page itself
    UPDATE pages SET next_page_key = NEW.key WHERE next_page_key = NEW.next_page_key AND list_key = NEW.list_key AND key <> NEW.key;
    -- if we inserted a page in the front of the list, update the list's next_page_key
    UPDATE lists SET next_page_key = NEW.key WHERE key = NEW.list_key AND next_page_key = NEW.next_page_key;
  -- insert at the end of the list 
  ELSE
    -- update the previous last page to point to the new page
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
  -- ignore trigger when on cascade delete
  IF pg_trigger_depth() > 1 THEN
    RETURN OLD;
  END IF;
  -- if the deleted page is the latest page, update the list's latest_page_key
  UPDATE lists SET latest_page_key = (SELECT key FROM pages WHERE list_key = OLD.list_key AND next_page_key = OLD.key LIMIT 1) WHERE key = OLD.list_key AND latest_page_key = OLD.key;
  -- if the deleted page is the first page, update the list's next_page_key
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