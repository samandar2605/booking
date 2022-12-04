package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/booking/storage/repo"
)

type orderRepo struct {
	db *sqlx.DB
}

func NewOrder(db *sqlx.DB) repo.OrderStorageI {
	return &orderRepo{db: db}
}

func (ur *orderRepo) CreateOrder(order *repo.Order) (*repo.Order, error) {

	query := `
		insert into order(
			hotel_id,
			room_id,
			full_name,
			date_first,
			date_last,
			adults_count,
			children_count,
			user_id
		)values($1,$2,$3,$4,$5,$6,$7,$8)
		where ((date_first>=(select date_last from order) 
			and date_last<=(select date_fist from order))or
		((date_first<(select date_last from order) 
			and date_last<=(select date_fist from order))) 
		and 
		returning id
	`

	row := ur.db.QueryRow(
		query,
		order.HotelId,
		order.RoomId,
		order.FullName,
		order.DateFirst,
		order.DateLast,
		order.AdultsCount,
		order.ChildrenCount,
		order.UserId,
	)

	if err := row.Scan(&order.Id); err != nil {
		return nil, err
	}

	return order, nil
}

func (ur *orderRepo) GetOrder(userId int) ([]*repo.Order, error) {
	var orders []*repo.Order

	query := `
		SELECT 
			id,
			hotel_id,
			room_id,
			full_name,
			date_first,
			date_last,
			adults_count,
			children_count,
			user_id
		from orders
		where user_id=$1
		order by id desc
	`
	rows, err := ur.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order repo.Order
		if err := rows.Scan(
			&order.Id,
			&order.HotelId,
			&order.RoomId,
			&order.FullName,
			&order.DateFirst,
			&order.DateLast,
			&order.AdultsCount,
			&order.ChildrenCount,
			&order.UserId,
		); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}
