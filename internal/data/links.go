package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Imprezzaa/linkshortener/internal/validator"
)

type Link struct {
	// Temoprarily removing userID until the users table is completed
	// UserID  int64  `json:"created_by"`
	ShortID   string    `json:"short_id"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
}

type LinkModel struct {
	DB *sql.DB
}

// need to recreate the shortID generator
func (l LinkModel) Insert(link *Link) error {
	link.ShortID = MakeString(5)

	// readd created_by later
	query := `
	INSERT INTO links (link, shortid)
	VALUES ($1, $2)
	RETURNING created_at`

	args := []any{link.Link, link.ShortID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return l.DB.QueryRowContext(ctx, query, args...).Scan(&link.CreatedAt)
}

func (l LinkModel) Get(shortid string) (*Link, error) {
	if len(shortid) != 5 {
		return nil, ErrRecordNotFound
	}

	query := `
	SELECT shortid, link, created_at
	FROM links
	WHERE shortid = $1`

	var link Link

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := l.DB.QueryRowContext(ctx, query, shortid).Scan(&link.ShortID, &link.Link, &link.CreatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &link, nil
}

func (l LinkModel) Patch(link *Link) error {
	query := `
	UPDATE links
	SET link = $1
	WHERE shortid = $2`

	args := []any{link.Link, link.ShortID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return l.DB.QueryRowContext(ctx, query, args...).Err()
}

func (l LinkModel) Delete(shortid string) error {
	if len(shortid) != 5 {
		return ErrRecordNotFound
	}

	query := `
	DELETE FROM links
	WHERE shortid = $1`

	result, err := l.DB.Exec(query, shortid)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func ValidateLink(v *validator.Validator, link *Link) {
	v.Check(link.Link != "", "link", "must provide a link")
}
