package stack

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

func (pr *Repository) Create(order int, svg string, title string) (*model.Stack, error) {
	var stack model.Stack

	q := `INSERT INTO Stack ("order", "svg", "title")
	VALUES($1, $2, $3) RETURNING *`
	err := pr.db.QueryRow(context.Background(), q,
		order, svg, title).Scan(
		&stack.Id,
		&stack.Order,
		&stack.Svg,
		&stack.Title,
	)
	if err != nil {
		return nil, err
	}

	return &stack, nil
}

func (pr *Repository) GetAll() ([]*model.Stack, error) {
	var stacks []*model.Stack

	q := `SELECT * FROM Stack`

	rows, err := pr.db.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row model.Stack
		err := rows.Scan(&row.Id, &row.Order, &row.Title, &row.Svg)

		if err != nil {
			pr.log.Error("Error to parse", err)
		}

		stacks = append(stacks, &row)
	}
	return stacks, nil
}

func (pr *Repository) Delete(id int) (*model.Stack, error) {
	var stack model.Stack

	q := `DELETE FROM Stack WHERE id = $1 RETURNING *`

	err := pr.db.QueryRow(context.Background(), q, id).Scan(
		&stack.Id,
		&stack.Order,
		&stack.Svg,
		&stack.Title)

	if err != nil {
		return nil, err
	}

	return &stack, nil
}
