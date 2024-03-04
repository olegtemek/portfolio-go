package info

import (
	"log/slog"

	"github.com/olegtemek/portfolio-go/internal/dto"
	"github.com/olegtemek/portfolio-go/internal/model"
	"github.com/olegtemek/portfolio-go/internal/repository"
)

type Service struct {
	log        *slog.Logger
	repository *repository.Info
}

func NewService(log *slog.Logger, repository *repository.Info) *Service {
	return &Service{
		log:        log,
		repository: repository,
	}
}

func (is *Service) CreateOrUpdate(info *dto.InfoCreateOrUpdateDto) (*model.Info, error) {

	Info, err := (*is.repository).CreateOrUpdate(info.Email, info.Github, info.Telegram, info.Location, info.Position, info.About)

	if err != nil {
		return nil, err
	}

	return Info, nil
}

func (is *Service) Get() (*model.Info, error) {

	Info, err := (*is.repository).Get()

	if err != nil {
		return nil, err
	}

	return Info, nil
}
