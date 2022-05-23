CREATE TABLE IF NOT EXISTS stories (
  id serial PRIMARY KEY,
  date timestamptz NULL,
  place character varying(255) NULL,
  source character varying(255) NULL,
  text_ua text NULL,
  text_en text NULL,
  title_ua text NULL,
  title_en text NULL,
  updated_at timestamptz NULL,
  created_at timestamptz NULL,
  deleted_at timestamptz NULL
);

CREATE INDEX IF NOT EXISTS place_idx ON stories (place);