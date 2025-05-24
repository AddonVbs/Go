package taskservers

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) error
	GetAllTask() ([]Task, error)
	GetTaskByID(id int) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id int) error
}

type RepositorysTasks struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &RepositorysTasks{db: db}
}

func (r *RepositorysTasks) CreateTask(task Task) error {
	return r.db.Create(&task).Error
}

func (r *RepositorysTasks) GetAllTask() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err

}

func (r *RepositorysTasks) GetTaskByID(id int) (Task, error) {
	var t Task
	err := r.db.First(&t, "id = ?", id).Error
	return t, err
}

func (r *RepositorysTasks) UpdateTask(task Task) error {
	return r.db.Save(&task).Error

}

func (r *RepositorysTasks) DeleteTask(id int) error {
	return r.db.Delete(&Task{}, "id = ? ", id).Error
}
