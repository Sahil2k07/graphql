package repositories

import (
	"errors"

	"github.com/Sahil2k07/graphql/internal/database"
	"github.com/Sahil2k07/graphql/internal/interfaces"
	"github.com/Sahil2k07/graphql/internal/models"
	"github.com/Sahil2k07/graphql/internal/utils"
	"gorm.io/gorm"
)

type todoRepository struct {
}

func NewTodoRepository() interfaces.TodoRepository {
	return &todoRepository{}
}

func (r *todoRepository) Create(todo *models.Todo) error {
	return database.DB.Create(todo).Error
}

func (r *todoRepository) Update(todo *models.Todo) error {
	return database.DB.Save(todo).Error
}

func (r *todoRepository) Delete(id uint, userID uint) error {
	res := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{})
	if res.RowsAffected == 0 {
		return errors.New("todo not found or not owned by user")
	}
	return res.Error
}

func (r *todoRepository) GetByID(id uint, userID uint) (*models.Todo, error) {
	var todo models.Todo
	err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &todo, err
}

func (r *todoRepository) GetAll(userID uint, pageFilter utils.PageFilter, sortFilter utils.SortFilter) ([]models.Todo, int64, error) {
	var todos []models.Todo
	var total int64

	query := database.DB.Model(&models.Todo{}).Where("user_id = ?", userID)
	query.Count(&total)

	query = utils.AddPagination(query, pageFilter, sortFilter)
	err := query.Find(&todos).Error

	return todos, total, err
}
