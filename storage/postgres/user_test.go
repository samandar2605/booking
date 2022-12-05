package postgres_test

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/samandar2605/booking/storage/repo"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *repo.User {
	User, err := strg.User().Create(&repo.User{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		UserName:  faker.Username(),
		Type:      "user",
	})
	require.NoError(t, err)
	require.NotEmpty(t, User)
	return User
}

func deleteUser(id int, t *testing.T) {
	err := strg.User().Delete(id)
	require.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	n := createUser(t)
	note, err := strg.User().Get(n.Id)
	require.NoError(t, err)
	require.NotEmpty(t, note)

	deleteUser(note.Id, t)
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestUpdateUser(t *testing.T) {
	n := createUser(t)

	n.FirstName = faker.FirstName()
	n.LastName = faker.LastName()
	n.Email = faker.Email()
	n.UserName = faker.Username()
	n.Type = "user"

	User, err := strg.User().Update(n)
	require.NoError(t, err)
	require.NotEmpty(t, User)

	deleteUser(User.Id, t)
}

func TestDeleteUser(t *testing.T) {
	u := createUser(t)
	deleteUser(u.Id, t)
}

func TestGetAllUser(t *testing.T) {
	u := createUser(t)
	n, _ := faker.RandomInt(100)
	_, err := strg.User().GetAll(repo.GetUserQuery{
		Page:  n[0],
		Limit: n[0],
	})

	require.NoError(t, err)
	deleteUser(u.Id, t)
}

func TestGetByEmail(t *testing.T) {
	u := createUser(t)
	usr, err := strg.User().GetByEmail(u.Email)
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	deleteUser(u.Id, t)
}

func TestUpdatePassword(t *testing.T) {
	u := createUser(t)
	err := strg.User().UpdatePassword(&repo.UpdatePassword{
		UserID:   int64(u.Id),
		Password: faker.Password(),
	})
	require.NoError(t, err)
	deleteUser(u.Id, t)
}

func TestCheckInfo(t *testing.T) {
	u := createUser(t)
	user, err := strg.User().CheckInfo(u.Email, u.UserName)
	require.NotEmpty(t, user)
	require.NoError(t, err)
	deleteUser(u.Id, t)
}

func TestGetUserProfileInfo(t *testing.T) {
	u := createUser(t)
	user, err := strg.User().GetUserProfileInfo(u.Id)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	deleteUser(u.Id, t)
}
