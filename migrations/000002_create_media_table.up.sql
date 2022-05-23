CREATE TABLE IF NOT EXISTS media (
  id serial PRIMARY KEY,
  url text NULL,
  type character varying(50) NULL,
  story_id integer NULL,
  updated_at timestamptz NULL,
  created_at timestamptz NULL,
  deleted_at timestamptz NULL,
  FOREIGN KEY (story_id) REFERENCES stories(id) ON DELETE SET NULL
);
