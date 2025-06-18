-- 1) Добавляем колонку user_id (пока nullable)
ALTER TABLE tasks
  ADD COLUMN user_id INTEGER;

-- 2) Заполняем уже существующие строки (пример: всем ставим user_id = 1)
UPDATE tasks
  SET user_id = 1
  WHERE user_id IS NULL;

-- 3) Делаем колонку NOT NULL и вешаем внешний ключ
ALTER TABLE tasks
  ALTER COLUMN user_id SET NOT NULL,
  ADD CONSTRAINT fk_tasks_user
    FOREIGN KEY (user_id) REFERENCES users(id);