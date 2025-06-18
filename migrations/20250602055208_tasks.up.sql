CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    user_id INTEGER NOT NULL,
    CONSTRAINT fk_tasks_user FOREIGN KEY (user_id) REFERENCES users(id)
);