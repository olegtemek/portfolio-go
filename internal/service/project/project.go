package project

import (
	"log/slog"

	"github.com/olegtemek/portfolio-go/internal/dto"
	"github.com/olegtemek/portfolio-go/internal/model"
	"github.com/olegtemek/portfolio-go/internal/repository"
)

type Service struct {
	log        *slog.Logger
	repository *repository.Project
}

func NewService(log *slog.Logger, repository *repository.Project) *Service {
	return &Service{
		log:        log,
		repository: repository,
	}
}

func (is *Service) Create(projectDto *dto.ProjectCreateDto) (*model.Project, error) {

	project, err := (*is.repository).Create(projectDto.Title, projectDto.Order, projectDto.Link, projectDto.Description)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (is *Service) GetAll() ([]*model.Project, error) {

	projects, err := (*is.repository).GetAll()

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (is *Service) Update(id int, projectDto *dto.ProjectUpdateDto) (*model.Project, error) {

	project, err := (*is.repository).Update(id, projectDto.Order, projectDto.Title, projectDto.Link, projectDto.Description)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (is *Service) Delete(id int) (*model.Project, error) {

	project, err := (*is.repository).Delete(id)

	if err != nil {
		return nil, err
	}

	return project, nil
}
