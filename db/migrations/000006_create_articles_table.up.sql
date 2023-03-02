CREATE TABLE IF NOT EXISTS articles (
  id SERIAL NOT NULL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  content VARCHAR(300) NOT NULL,
  author_id SERIAL REFERENCES users(id),
  topic_id SERIAL REFERENCES topics(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_articles
BEFORE UPDATE ON articles
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
