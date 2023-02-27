DROP TABLE IF EXISTS "histories" CASCADE;
CREATE TABLE IF NOT EXISTS "histories" (
  "repository_id" varchar(100) NOT NULL,
  "run_id" varchar(100) NOT NULL,
  "status" varchar(10) NOT NULL DEFAULT 'STARTED',
  "started_at" timestamp NOT NULL DEFAULT 'now',
  "finished_at" timestamp,
  PRIMARY KEY("repository_id","run_id")
);

DROP TABLE IF EXISTS "repositories" CASCADE;
CREATE TABLE IF NOT EXISTS "repositories" (
  "repository_id" varchar(100) NOT NULL PRIMARY KEY,
  "repository_name" varchar(100) NOT NULL
);