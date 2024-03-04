package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/portfolio-go/internal/model"
	"github.com/olegtemek/portfolio-go/internal/repository/experience"
	"github.com/olegtemek/portfolio-go/internal/repository/info"
	"github.com/olegtemek/portfolio-go/internal/repository/project"
	"github.com/olegtemek/portfolio-go/internal/repository/stack"
)

type Info interface {
	CreateOrUpdate(email string, github string, telegram string, location string, position string, about string) (*model.Info, error)
	Get() (*model.Info, error)
}
type Experience interface {
	Create(title string, order int, position string, period string, description string) (*model.Experience, error)
	Update(id int, order int, title string, position string, period string, description string) (*model.Experience, error)
	GetAll() ([]*model.Experience, error)
	Delete(id int) (*model.Experience, error)
}
type Project interface {
	Create(title string, order int, link string, description string) (*model.Project, error)
	Update(id int, order int, title string, link string, description string) (*model.Project, error)
	GetAll() ([]*model.Project, error)
	Delete(id int) (*model.Project, error)
}
type Stack interface {
	Create(order int, svg string, title string) (*model.Stack, error)
	GetAll() ([]*model.Stack, error)
	Delete(id int) (*model.Stack, error)
}

type Repository struct {
	Info
	Experience
	Stack
	Project
}

func New(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		Info:       info.NewRepository(log, db),
		Experience: experience.NewRepository(log, db),
		Project:    project.NewRepository(log, db),
		Stack:      stack.NewRepository(log, db),
	}
}
