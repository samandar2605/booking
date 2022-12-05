package postgres_test

import (
	"testing"

	"github.com/samandar2605/booking/storage/repo"
	"github.com/stretchr/testify/require"
)

func createLike(t *testing.T) (repo.Like, error) {
	u := createUser(t)
	h := createHotel(t)
	like := repo.Like{
		UserId:  u.Id,
		HotelId: h.Id,
		Status:  true,
	}
	err := strg.Like().CreateOrUpdate(&like)
	return like, err
}
func TestGetLike(t *testing.T) {
	n, err := createLike(t)
	require.NoError(t, err)
	note, err := strg.Like().Get(n.UserId, n.HotelId)
	require.NoError(t, err)
	require.NotEmpty(t, note)
}

func TestCreateLike(t *testing.T) {
	err := strg.Like().CreateOrUpdate(&repo.Like{
		UserId:  1,
		HotelId: 1,
		Status:  true,
	})
	require.NoError(t, err)
}

func TestDeleteLike(t *testing.T) {
	u, err := createLike(t)
	require.NoError(t, err)
	require.NotEmpty(t, u)
}

func TestGetAllInfo(t *testing.T) {
	l, err := createLike(t)
	require.NoError(t, err)
	result, err := strg.Like().GetLikesDislikesCount(l.HotelId)
	require.NotEmpty(t, result)
	require.NoError(t, err)
}
