package experience

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/portfolio-go/internal/model"
)

type Repository struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewRepository(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		log: log,
		db:  db,
	}
}

func (pr *Repository) Create(title string, order int, position string, period string, description string) (*model.Experience, error) {
	var experience model.Experience

	q := `INSERT INTO Experience ("order", "title", "position", "period", "description")
	VALUES($1, $2, $3, $4, $5) RETURNING *`
	err := pr.db.QueryRow(context.Background(), q,
		order, title, position, period, description).Scan(
		&experience.Id,
		&experience.Order,
		&experience.Title,
		&experience.Position,
		&experience.Period,
		&experience.Description)
	if err != nil {
		return nil, err
	}

	return &experience, nil
}

func (pr *Repository) GetOne(id int) (*model.Experience, error) {
	var experience model.Experience

	q := `SELECT * FROM Experience WHERE id = $1 LIMIT 1`

	err := pr.db.QueryRow(context.Background(), q, id).Scan(
		&experience.Id,
		&experience.Order,
		&experience.Title,
		&experience.Position,
		&experience.Period,
		&experience.Description)
	if err != nil {
		return nil, err
	}

	return &experience, nil
}

func (pr *Repository) GetAll() ([]*model.Experience, error) {
	var experiences []*model.Experience

	q := `SELECT * FROM Experience`

	rows, err := pr.db.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row model.Experience
		err := rows.Scan(&row.Id, &row.Order, &row.Title, &row.Position, &row.Period, &row.Description)

		if err != nil {
			pr.log.Error("Error to parse", err)
		}

		experiences = append(experiences, &row)
	}

	return experiences, nil
}

func (pr *Repository) Update(id int, order int, title string, position string, period string, description string) (*model.Experience, error) {

	var experience model.Experience

	q := `UPDATE Experience
	SET "order" = $2, title = $3, position = $4, period = $5, description = $6
	WHERE id = $1 RETURNING *;
	`

	err := pr.db.QueryRow(context.Background(), q, id, order, title, position, period, description).Scan(
		&experience.Id,
		&experience.Order,
		&experience.Title,
		&experience.Position,
		&experience.Period,
		&experience.Description)

	if err != nil {
		return nil, err
	}

	return &experience, nil
}

func (pr *Repository) Delete(id int) (*model.Experience, error) {
	var experience model.Experience

	q := `DELETE FROM Experience WHERE id = $1 RETURNING *`

	err := pr.db.QueryRow(context.Background(), q, id).Scan(
		&experience.Id,
		&experience.Order,
		&experience.Title,
		&experience.Position,
		&experience.Period,
		&experience.Description)

	if err != nil {
		return nil, err
	}

	return &experience, nil
}
