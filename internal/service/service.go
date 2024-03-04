package service

import (
	"log/slog"

	"github.com/olegtemek/portfolio-go/internal/dto"
	"github.com/olegtemek/portfolio-go/internal/model"
	"github.com/olegtemek/portfolio-go/internal/repository"
	"github.com/olegtemek/portfolio-go/internal/service/experience"
	"github.com/olegtemek/portfolio-go/internal/service/info"
	"github.com/olegtemek/portfolio-go/internal/service/project"
	"github.com/olegtemek/portfolio-go/internal/service/stack"
)

type Info interface {
	CreateOrUpdate(info *dto.InfoCreateOrUpdateDto) (*model.Info, error)
	Get() (*model.Info, error)
}

type Experience interface {
	Create(experienceDto *dto.ExperienceCreateDto) (*model.Experience, error)
	Update(id int, experienceDto *dto.ExperienceUpdateDto) (*model.Experience, error)
	GetAll() ([]*model.Experience, error)
	Delete(id int) (*model.Experience, error)
}
type Project interface {
	Create(ProjectDto *dto.ProjectCreateDto) (*model.Project, error)
	Update(id int, ProjectDto *dto.ProjectUpdateDto) (*model.Project, error)
	GetAll() ([]*model.Project, error)
	Delete(id int) (*model.Project, error)
}
type Stack interface {
	Create(stackDto *dto.StackCreateDto) (*model.Stack, error)
	GetAll() ([]*model.Stack, error)
	Delete(id int) (*model.Stack, error)
}

type Service struct {
	Info
	Experience
	Stack
	Project
}

func NewService(log *slog.Logger, repositories *repository.Repository) *Service {
	return &Service{
		Info:       info.NewService(log, &repositories.Info),
		Experience: experience.NewService(log, &repositories.Experience),
		Project:    project.NewService(log, &repositories.Project),
		Stack:      stack.NewService(log, &repositories.Stack),
	}
}
