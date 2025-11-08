package interfaces

import (
	"context"

	"github.com/Sahil2k07/graphql/internal/graphql/generated"
	"github.com/Sahil2k07/graphql/internal/models"
	"github.com/Sahil2k07/graphql/internal/utils"
)

type TodoRepository interface {
	Create(todo *models.Todo) error
	Update(todo *models.Todo) error
	Delete(id uint, userID uint) error
	GetByID(id uint, userID uint) (*models.Todo, error)
	GetAll(userID uint, pageFilter utils.PageFilter, sortFilter utils.SortFilter) ([]models.Todo, int64, error)
}

type TodoService interface {
	CreateTodo(ctx context.Context, input generated.CreateTodoInput) (*generated.Todo, error)
	UpdateTodo(ctx context.Context, input generated.UpdateTodoInput) (*generated.Todo, error)
	DeleteTodo(ctx context.Context, id string) (string, error)
	GetTodoByID(ctx context.Context, id string) (*generated.Todo, error)
	GetTodos(ctx context.Context, page *int, limit *int) (*generated.TodoPage, error)
}
