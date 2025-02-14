package repositories

import (
	"github.com/MohamedMosalm/Todo-App/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) TaskRepository {
	return &gormTaskRepository{db: db}
}

func (r *gormTaskRepository) CreateTask(task *models.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormTaskRepository) GetTasksByUserID(userID uuid.UUID) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *gormTaskRepository) UpdateTask(taskID uuid.UUID, updates map[string]interface{}) error {
	return r.db.Model(&models.Task{}).Where("id = ?", taskID).Updates(updates).Error
}

func (r *gormTaskRepository) DeleteTask(taskID, userID uuid.UUID) error {
	if err := r.db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormTaskRepository) GetTaskByID(taskID uuid.UUID) (*models.Task, error) {
	var task models.Task
	err := r.db.Where("id = ?", taskID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}
