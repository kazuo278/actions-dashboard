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
  histories(repository_id, run_id, status, started_at, finished_at)
VALUES
  ('test1_id', 'run_id1', 'STARTED', '2023-02-22T03:00Z', null);
INSERT INTO
  histories(repository_id, run_id, status, started_at, finished_at)
VALUES
  ('test1_id', 'run_id2', 'FINISHED', '2023-02-22T03:05Z', '2023-02-22T03:10Z');

-- test2リポジトリの実行履歴
INSERT INTO
  histories(repository_id, run_id, status, started_at, finished_at)
VALUES
  ('test2_id', 'run_id1', 'STARTED', '2023-02-22T03:10Z', null);
INSERT INTO
  histories(repository_id, run_id, status, started_at, finished_at)
VALUES
  ('test2_id', 'run_id2', 'FINISHED', '2023-02-22T03:15Z', '2023-02-22T03:20Z');

