DROP TRIGGER IF EXISTS trigger_latest_page_after ON pages;
DROP TRIGGER IF EXISTS trigger_latest_page_before ON pages;
DROP FUNCTION IF EXISTS update_latest_page_on_delete;
DROP FUNCTION IF EXISTS update_latest_page_on_insert;
