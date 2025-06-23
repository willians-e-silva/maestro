package repository

import (
	task "maestro/internal/domain/task"

	"gorm.io/gorm"
)

type PostgresTaskRepository struct {
	db *gorm.DB
}

func NewPostgresTaskRepository(db *gorm.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) CreateTask(task *task.Task) (*task.Task, error) {
	result := r.db.Create(task)

	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (r *PostgresTaskRepository) GetAllTasks() ([]task.Task, error) {
	var tasks []task.Task
	result := r.db.Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
