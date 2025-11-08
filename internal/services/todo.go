package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Sahil2k07/graphql/internal/enums"
	errz "github.com/Sahil2k07/graphql/internal/errors"
	"github.com/Sahil2k07/graphql/internal/graphql/generated"
	"github.com/Sahil2k07/graphql/internal/interfaces"
	"github.com/Sahil2k07/graphql/internal/models"
	"github.com/Sahil2k07/graphql/internal/utils"
)

type todoService struct {
	repo interfaces.TodoRepository
}

func NewTodoService(repo interfaces.TodoRepository) interfaces.TodoService {
	return &todoService{repo: repo}
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func formatTimePtr(t any) *string {
	if tm, ok := t.(interface{ Format(string) string }); ok {
		formatted := tm.Format("2006-01-02T15:04:05Z")
		return &formatted
	}
	return nil
}

func (s *todoService) CreateTodo(ctx context.Context, input generated.CreateTodoInput) (*generated.Todo, error) {
	user, err := utils.GetUserClaims(ctx)
	if err != nil {
		return nil, errz.NewUnauthorized("unauthorized")
	}

	todo := &models.Todo{
		Title:       input.Title,
		Description: input.Description,
		UserID:      user.ID,
	}

	if err := s.repo.Create(todo); err != nil {
		return nil, errz.NewInternalError(fmt.Sprintf("failed to create todo: %v", err))
	}

	return &generated.Todo{
		ID:          fmt.Sprintf("%d", todo.ID),
		Title:       todo.Title,
		Description: strPtr(todo.Description),
		Status:      string(todo.Status),
		CreatedAt:   formatTimePtr(todo.CreatedAt),
		UpdatedAt:   formatTimePtr(todo.UpdatedAt),
	}, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, input generated.UpdateTodoInput) (*generated.Todo, error) {
	user, err := utils.GetUserClaims(ctx)
	if err != nil {
		return nil, errz.NewUnauthorized("unauthorized")
	}

	id, err := strconv.ParseUint(input.ID, 10, 64)
	if err != nil {
		return nil, errz.NewValidation("invalid todo id")
	}

	todo, err := s.repo.GetByID(uint(id), user.ID)
	if err != nil {
		return nil, errz.NewNotFound("todo not found")
	}

	todo.Title = input.Title
	if input.Description != nil {
		todo.Description = *input.Description
	}
	if input.Status != nil {
		todo.Status = enums.TodoStatus(*input.Status)
	}

	if err := s.repo.Update(todo); err != nil {
		return nil, errz.NewInternalError(fmt.Sprintf("failed to update todo: %v", err))
	}

	return &generated.Todo{
		ID:          fmt.Sprintf("%d", todo.ID),
		Title:       todo.Title,
		Status:      string(todo.Status),
		Description: strPtr(todo.Description),
		CreatedAt:   formatTimePtr(todo.CreatedAt),
		UpdatedAt:   formatTimePtr(todo.UpdatedAt),
	}, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, id string) (string, error) {
	user, err := utils.GetUserClaims(ctx)
	if err != nil {
		return "", errz.NewUnauthorized("unauthorized")
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return "", errz.NewValidation("invalid todo id")
	}

	if err := s.repo.Delete(uint(idUint), user.ID); err != nil {
		return "", errz.NewInternalError(fmt.Sprintf("failed to delete todo: %v", err))
	}

	return "Todo deleted successfully", nil
}

func (s *todoService) GetTodoByID(ctx context.Context, id string) (*generated.Todo, error) {
	user, err := utils.GetUserClaims(ctx)
	if err != nil {
		return nil, errz.NewUnauthorized("unauthorized")
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, errz.NewValidation("invalid todo id")
	}

	todo, err := s.repo.GetByID(uint(idUint), user.ID)
	if err != nil {
		return nil, errz.NewNotFound("todo not found")
	}

	return &generated.Todo{
		ID:          fmt.Sprintf("%d", todo.ID),
		Title:       todo.Title,
		Status:      string(todo.Status),
		Description: strPtr(todo.Description),
		CreatedAt:   formatTimePtr(todo.CreatedAt),
		UpdatedAt:   formatTimePtr(todo.UpdatedAt),
	}, nil
}

func (s *todoService) GetTodos(ctx context.Context, page *int, limit *int) (*generated.TodoPage, error) {
	user, err := utils.GetUserClaims(ctx)
	if err != nil {
		return nil, errz.NewUnauthorized("unauthorized")
	}

	pf := utils.PageFilter{Page: 1, Limit: 10}
	if page != nil {
		pf.Page = *page
	}
	if limit != nil {
		pf.Limit = *limit
	}

	todos, total, err := s.repo.GetAll(user.ID, pf, utils.SortFilter{})
	if err != nil {
		return nil, errz.NewInternalError(fmt.Sprintf("failed to fetch todos: %v", err))
	}

	gqlTodos := make([]*generated.Todo, 0, len(todos))
	for _, t := range todos {
		gqlTodos = append(gqlTodos, &generated.Todo{
			ID:          fmt.Sprintf("%d", t.ID),
			Title:       t.Title,
			Status:      string(t.Status),
			Description: strPtr(t.Description),
			CreatedAt:   formatTimePtr(t.CreatedAt),
			UpdatedAt:   formatTimePtr(t.UpdatedAt),
		})
	}

	return &generated.TodoPage{
		Todos:      gqlTodos,
		TotalCount: int(total),
		Page:       pf.Page,
		Limit:      pf.Limit,
	}, nil
}
