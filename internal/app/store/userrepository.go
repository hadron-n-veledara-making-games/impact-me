package store

import (
	"database/sql"

	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/models"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *models.User) (*models.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (telegram_id) VALUES ($1) RETURNING id",
		u.TelegramID,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByTelegramID(id int) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, telegram_id FROM users WHERE telegram_id = $1",
		id,
	).Scan(
		&u.ID,
		&u.TelegramID,
	); err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return u, nil
}
