DROP TABLE IF EXISTS "histories" CASCADE;
CREATE TABLE IF NOT EXISTS "histories" (
  "repository_id" varchar(100) NOT NULL,
  "run_id" varchar(100) NOT NULL,
  "workflow_ref"  varchar(100) NOT NULL,
  "job_name" varchar(100) NOT NULL,
  "run_attempt" varchar(10) NOT NULL,
  "status" varchar(10) NOT NULL DEFAULT 'STARTED',
  "started_at" timestamp WITH TIME ZONE NOT NULL DEFAULT 'now',
  "finished_at" timestamp WITH TIME ZONE,
  PRIMARY KEY("repository_id","run_id","job_name","run_attempt")
);

DROP TABLE IF EXISTS "repositories" CASCADE;
CREATE TABLE IF NOT EXISTS "repositories" (
  "repository_id" varchar(100) NOT NULL PRIMARY KEY,
  "repository_name" varchar(100) NOT NULL
);