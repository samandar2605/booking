package postgres

import (
	"time"

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
		insert into orders(
			hotel_id,
			room_id,
			date_first,
			days,
			adults_count,
			children_count,
			user_id
		)values($1,$2,$3,$4,$5,$6,$7)
		RETURNING id,(select ($5*price_for_adults)+($6*price_for_children) from rooms where hotel_id=$1 and id=$2)
	`
	row := ur.db.QueryRow(
		query,
		order.HotelId,
		order.RoomId,
		order.DateFirst,
		order.Days,
		order.AdultsCount,
		order.ChildrenCount,
		order.UserId)
	if err := row.Scan(
		&order.Id,
		&order.Summa,
	); err != nil {
		return nil, err
	}
	_, err := ur.db.Exec("update order set summa=$1 where id=$2", order.Summa, order.Id)
	if err != nil {
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
			date_first,
			days,
			adults_count,
			children_count,
			user_id,
			summa
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
			&order.DateFirst,
			&order.Days,
			&order.AdultsCount,
			&order.ChildrenCount,
			&order.UserId,
			&order.Summa,
		); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (ur *orderRepo) CheckRoom(order *repo.Order) bool {
	var ls []*int
	query := `
		select 
			id
		from orders where hote_id=$1 and ((date_first<$2 and date_first<$3) or (date_first+interval '$4 days'<$2 and $3<date_first+interval '$4 days') )
	`
	rows, err := ur.db.Query(
		query,
		order.HotelId,
		order.DateFirst,
		order.DateFirst.Add(time.Duration(order.Days)*86400),
		order.Days,
	)
	if err != nil {
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var n *int
		err = rows.Scan(&n)
		if err != nil {
			return false
		}
		ls = append(ls, n)
	}
	return len(ls) > 0
}
