-- repository
INSERT INTO
  repositories(repository_id, repository_name)
VALUES
  ('test1_id', 'test1_name');
INSERT INTO
  repositories(repository_id, repository_name)
VALUES
  ('test2_id', 'test2_name');

-- test1リポジトリの実行履歴
INSERT INTO
  histories(repository_id, run_id, workflow_ref, job_name, run_attempt, status, started_at, finished_at)
VALUES
  ('test1_id', 'run_id1-1', '.github/workflows/workflow1-1@ref1-1', 'job1-1', '1', 'STARTED', '2023-02-22T03:00Z', null);
INSERT INTO
  histories(repository_id, run_id, workflow_ref, job_name, run_attempt, status, started_at, finished_at)
VALUES
  ('test1_id', 'run_id1-2', '.github/workflows/workflow1-2@ref1-2', 'job1-2', '1', 'FINISHED', '2023-02-22T03:05Z', '2023-02-22T03:10Z');
INSERT INTO
  histories(repository_id, run_id, workflow_ref, job_name, run_attempt, status, started_at, finished_at)
VALUES
  ('test1_id', 'run_id1-3', '.github/workflows/workflow1-2@ref1-2', 'job1-2', '1', 'FINISHED', '2023-02-22T04:05Z', '2023-02-23T04:05Z');
-- test2リポジトリの実行履歴
INSERT INTO
  histories(repository_id, run_id, workflow_ref, job_name, run_attempt, status, started_at, finished_at)
VALUES
  ('test2_id', 'run_id2-1', '.github/workflows/workflow2-1@ref2-1', 'job2-1', '1', 'STARTED', '2023-02-22T03:10Z', null);
INSERT INTO
  histories(repository_id, run_id, workflow_ref, job_name, run_attempt, status, started_at, finished_at)
VALUES
  ('test2_id', 'run_id2-2', '.github/workflows/workflow2-2@ref2-2', 'job2-2', '1', 'FINISHED', '2023-02-22T03:15Z', '2023-02-22T03:20Z');

