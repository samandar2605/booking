package postgres

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/booking/storage/repo"
)

type likeRepo struct {
	db *sqlx.DB
}

func NewLike(db *sqlx.DB) repo.LikeStorageI {
	return &likeRepo{
		db: db,
	}
}

func (lr *likeRepo) CreateOrUpdate(l *repo.Like) error {
	like, err := lr.Get(l.UserId, l.HotelId)
	if errors.Is(err, sql.ErrNoRows) {
		query := `
			INSERT INTO likes(user_id, post_id, status) 
			VALUES($1, $2, $3) RETURNING id
		`

		_, err := lr.db.Exec(query, l.UserId, l.HotelId, l.Status)
		if err != nil {
			return err
		}
	} else if like != nil {
		if like.Status == l.Status {
			_, err := lr.db.Exec(`DELETE FROM likes WHERE id=$1`, like.ID)
			if err != nil {
				return err
			}
		} else {
			query := `UPDATE likes SET status=$1 WHERE id=$2`
			_, err := lr.db.Exec(query, l.Status, like.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (cr *likeRepo) Get(userId, hotelId int) (*repo.Like, error) {
	var result repo.Like

	query := `
		SELECT
			id,
			user_id,
			post_id,
			status
		FROM likes
		WHERE user_id=$1 AND post_id=$2
	`

	row := cr.db.QueryRow(query, userId)
	err := row.Scan(
		&result.ID,
		&result.UserId,
		&result.HotelId,
		&result.Status,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *likeRepo) GetLikesDislikesCount(hotelId int) (repo.LikesDislikesCountsResult, error) {
	var result repo.LikesDislikesCountsResult

	query := `
		SELECT
			COUNT(1) FILTER (WHERE status=true) as likes_count,
			COUNT(1) FILTER (WHERE status=false) as dislikes_count
		FROM likes
		WHERE post_id=$1
	`

	row := cr.db.QueryRow(query, hotelId)
	err := row.Scan(
		&result.LikesCount,
		&result.DislikesCount,
	)
	if err != nil {
		return repo.LikesDislikesCountsResult{}, err
	}

	return result, nil
}
