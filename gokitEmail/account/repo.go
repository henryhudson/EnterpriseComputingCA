package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `INSERT INTO users2 (id, email, password, body, toEmail)
			VALUES ($1, $2, $3, $4, $5)`

	if user.Email == "" || user.Password == "" || user.Body == "" || user.ToEmail == "" {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password, user.Body, user.ToEmail)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	var body string
	var toEmail string

	err := repo.db.QueryRow("SELECT email, body, toEmail FROM users2 WHERE id=$1", id).Scan(&email, &body, &toEmail)
	if err != nil {
		return "", RepoErr
	}

	return "To: " + toEmail + " Body: " + body + " From: " + email, nil
}
