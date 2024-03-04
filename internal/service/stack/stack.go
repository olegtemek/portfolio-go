package stack

import (
	"log/slog"

	"github.com/olegtemek/portfolio-go/internal/dto"
	"github.com/olegtemek/portfolio-go/internal/model"
	"github.com/olegtemek/portfolio-go/internal/repository"
)

type Service struct {
	log        *slog.Logger
	repository *repository.Stack
}

func NewService(log *slog.Logger, repository *repository.Stack) *Service {
	return &Service{
		log:        log,
		repository: repository,
	}
}

func (is *Service) Create(stackDto *dto.StackCreateDto) (*model.Stack, error) {

	stack, err := (*is.repository).Create(stackDto.Order, stackDto.Svg, stackDto.Title)

	if err != nil {
		return nil, err
	}

	return stack, nil
}

func (is *Service) GetAll() ([]*model.Stack, error) {

	stacks, err := (*is.repository).GetAll()

	if err != nil {
		return nil, err
	}

	return stacks, nil
}

func (is *Service) Delete(id int) (*model.Stack, error) {

	stack, err := (*is.repository).Delete(id)

	if err != nil {
		return nil, err
	}

	return stack, nil
}
