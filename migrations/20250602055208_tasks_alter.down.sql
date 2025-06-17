-- 20250602055208_tasks_alter.down.sql
-- Откат добавления колонки user_id
ALTER TABLE tasks
  DROP CONSTRAINT IF EXISTS tasks_user_id_fkey;   -- если был FK
ALTER TABLE tasks
  DROP COLUMN IF EXISTS user_id;
