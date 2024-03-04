package experience

import (
	"log/slog"

	"github.com/olegtemek/portfolio-go/internal/dto"
	"github.com/olegtemek/portfolio-go/internal/model"
	"github.com/olegtemek/portfolio-go/internal/repository"
)

type Service struct {
	log        *slog.Logger
	repository *repository.Experience
}

func NewService(log *slog.Logger, repository *repository.Experience) *Service {
	return &Service{
		log:        log,
		repository: repository,
	}
}

func (is *Service) Create(experienceDto *dto.ExperienceCreateDto) (*model.Experience, error) {

	experience, err := (*is.repository).Create(experienceDto.Title, experienceDto.Order, experienceDto.Position, experienceDto.Period, experienceDto.Description)

	if err != nil {
		return nil, err
	}

	return experience, nil
}

func (is *Service) GetAll() ([]*model.Experience, error) {

	experiences, err := (*is.repository).GetAll()

	if err != nil {
		return nil, err
	}

	return experiences, nil
}

func (is *Service) Update(id int, experienceDto *dto.ExperienceUpdateDto) (*model.Experience, error) {

	experience, err := (*is.repository).Update(id, experienceDto.Order, experienceDto.Title, experienceDto.Position, experienceDto.Period, experienceDto.Description)

	if err != nil {
		return nil, err
	}

	return experience, nil
}

func (is *Service) Delete(id int) (*model.Experience, error) {

	experience, err := (*is.repository).Delete(id)

	if err != nil {
		return nil, err
	}

	return experience, nil
}
