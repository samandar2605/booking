package postgres_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/samandar2605/booking/storage/repo"
	"github.com/stretchr/testify/require"
)

func createOrder(t *testing.T) (*repo.Order, error) {
	h := createHotel(t)
	r := createRoom(t)
	u := createUser(t)
	n, _ := faker.RandomInt(100)
	Order := repo.Order{
		UserId:        u.Id,
		HotelId:       h.Id,
		RoomId:        r.Id,
		DateFirst:     time.Now(),
		Days:          n[0] % 10,
		AdultsCount:   n[0] % 10,
		ChildrenCount: n[0] % 10,
	}
	order, err := strg.Order().CreateOrder(&Order)
	return order, err
}
func TestGetOrder(t *testing.T) {
	o, err := createOrder(t)
	fmt.Println(err)
	require.NoError(t, err)
	note, err := strg.Order().GetOrder(o.Id)
	require.NoError(t, err)
	require.NotEmpty(t, note)
}

func TestCreateOrder(t *testing.T) {
	order, err := createOrder(t)
	require.NoError(t, err)
	require.NotEmpty(t, order)
}

func TestCheckRoom(t *testing.T) {
	order, err := createOrder(t)
	n := strg.Order().CheckRoom(order)
	require.NoError(t, err)
	require.NotEmpty(t, n)
}
