package project

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

func (pr *Repository) Create(title string, order int, link string, description string) (*model.Project, error) {
	var project model.Project

	q := `INSERT INTO Project ("order", "title", "link","description")
	VALUES($1, $2, $3, $4) RETURNING *`
	err := pr.db.QueryRow(context.Background(), q,
		order, title, link, description).Scan(
		&project.Id,
		&project.Order,
		&project.Title,
		&project.Link,
		&project.Description)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (pr *Repository) GetAll() ([]*model.Project, error) {
	var projects []*model.Project

	q := `SELECT * FROM Project`

	rows, err := pr.db.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row model.Project
		err := rows.Scan(&row.Id, &row.Order, &row.Title, &row.Link, &row.Description)

		if err != nil {
			pr.log.Error("Error to parse", err)
		}

		projects = append(projects, &row)
	}

	return projects, nil
}

func (pr *Repository) Update(id int, order int, title string, link string, description string) (*model.Project, error) {

	var project model.Project

	q := `UPDATE Project
	SET "order" = $2, title = $3, link = $4, description = $5
	WHERE id = $1 RETURNING *;
	`

	err := pr.db.QueryRow(context.Background(), q, id, order, title, link, description).Scan(
		&project.Id,
		&project.Order,
		&project.Title,
		&project.Link,
		&project.Description)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (pr *Repository) Delete(id int) (*model.Project, error) {
	var project model.Project

	q := `DELETE FROM Project WHERE id = $1 RETURNING *`

	err := pr.db.QueryRow(context.Background(), q, id).Scan(
		&project.Id,
		&project.Order,
		&project.Title,
		&project.Link,
		&project.Description)

	if err != nil {
		return nil, err
	}

	return &project, nil
}
