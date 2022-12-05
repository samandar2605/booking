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
			date_last,
			adults_count,
			children_count,
			user_id
		)values($1,$2,$3,$4,$5,$6,$7)
		RETURNING id
	`

	order.DateLast = order.DateFirst.Add(time.Duration(time.Duration(order.Days).Hours()) * 24)

	row := ur.db.QueryRow(
		query,
		order.HotelId,
		order.RoomId,
		order.DateFirst,
		order.DateLast,
		order.AdultsCount,
		order.ChildrenCount,
		order.UserId,
	)
	if err := row.Scan(
		&order.Id,
	); err != nil {
		return nil, err
	}

	var PriceForChildren, PriceForAdults float64
	row = ur.db.QueryRow("select price_for_adults,price_for_children from rooms where id=$1 and hotel_id=$2", order.RoomId, order.HotelId)
	err := row.Scan(&PriceForAdults, &PriceForChildren)
	if err != nil {
		return nil, err
	}

	_, err = ur.db.Exec("update orders set summa=$1 where id=$2", (PriceForAdults*float64(order.AdultsCount))+(PriceForChildren*float64(order.ChildrenCount)), order.HotelId)
	if err != nil {
		return nil, err
	}
	order.Summa = (PriceForAdults * float64(order.AdultsCount)) + (PriceForChildren * float64(order.ChildrenCount))
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
			date_last,
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
			&order.DateLast,
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
		from orders where hote_id=$1 and ((date_first<$2 and date_first<$3) or (date_last<$2 and $3<date_last) )
	`
	rows, err := ur.db.Query(
		query,
		order.HotelId,
		order.DateFirst,
		order.DateFirst.Add(time.Duration(time.Duration(order.Days).Hours())*24),
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
