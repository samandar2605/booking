package postgres_test

import (
	"strconv"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/samandar2605/booking/storage/repo"
	"github.com/stretchr/testify/require"
)

func createHotel(t *testing.T) *repo.Hotel {
	hotel, err := strg.Hotel().Create(&repo.Hotel{
		Name:        faker.Name(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		ImageUrl:    faker.URL(),
		Address:     faker.MacAddress(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, hotel)
	return hotel
}

func createRoom(t *testing.T) *repo.Room {
	h := createHotel(t)
	n, _ := strconv.Atoi(faker.CCNumber())
	hotel, err := strg.Hotel().AddRoom(&repo.Room{
		HotelId:          h.Id,
		ImageUrl:         faker.URL(),
		IsActive:         true,
		RoomType:         faker.Sentence(),
		PriceForChildren: float64(n),
		PriceForAdults:   float64(n) + 200,
	})
	require.NoError(t, err)
	require.NotEmpty(t, hotel)
	return hotel
}

func deleteHotel(id int, t *testing.T) {
	err := strg.Hotel().DeleteHotel(id)
	require.NoError(t, err)
}
func deleteRoom(id int, t *testing.T) {
	err := strg.Hotel().DeleteRoom(id)
	require.NoError(t, err)
}

func TestCreateHotel(t *testing.T) {
	h := createHotel(t)
	deleteHotel(h.Id, t)
}

func TestAddRoom(t *testing.T) {
	room := createRoom(t)
	err := strg.Hotel().DeleteRoom(room.Id)
	require.NotEmpty(t, room)
	require.NoError(t, err)
}
func TestAddRoomImage(t *testing.T) {
	h := createHotel(t)
	r := createRoom(t)
	n, _ := strconv.Atoi(faker.CCNumber())
	roomImage, err := strg.Hotel().AddRoomsImage(&repo.RoomsImage{
		RoomId:         r.Id,
		HotelId:        h.Id,
		ImageUrl:       faker.URL(),
		SequenceNumber: n,
	})
	require.NotEmpty(t, roomImage)
	require.NoError(t, err)
}


