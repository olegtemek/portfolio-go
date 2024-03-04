package info

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/portfolio-go/internal/model"
)

type Reository struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewRepository(log *slog.Logger, db *pgxpool.Pool) *Reository {
	return &Reository{
		log: log,
		db:  db,
	}
}

func (cr *Reository) CreateOrUpdate(email string, github string, telegram string, location string, position string, about string) (*model.Info, error) {
	info := &model.Info{}

	err := cr.db.QueryRow(context.Background(),
		`INSERT INTO Info (id, email, github, telegram, location, position, about)
			VALUES(1, $1, $2, $3, $4, $5, $6)
			ON CONFLICT (id) DO UPDATE SET
				email = $1,
				github = $2,
				telegram = $3,
				location = $4,
				position = $5,
				about = $6
			RETURNING *;
		`, email, github, telegram, location, position, about).Scan(&info.Id, &info.Email, &info.Github, &info.Telegram, &info.Location, &info.Position, &info.About)

	if err != nil {
		return nil, err
	}

	return info, nil
}

func (cr *Reository) Get() (*model.Info, error) {
	info := &model.Info{}

	err := cr.db.QueryRow(context.Background(),
		`SELECT * FROM Info WHERE id = 1 LIMIT 1;`).Scan(&info.Id, &info.Email, &info.Github, &info.Telegram, &info.Location, &info.Position, &info.About)

	if err != nil {
		return nil, err
	}

	return info, nil
}
