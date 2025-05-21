package db_test

import (
	"context"
	db "github/ronnie5562/fingreat_backend/db/sqlc"
	"github/ronnie5562/fingreat_backend/utils"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func clean_up() {
	err := testQuery.DeleteAllUsers(context.Background())

	if err != nil {
		log.Fatal("Unable to delete all users", err)
	}
}

func createRandomUser(t *testing.T) db.User {
	hashedPassword, err := utils.GenerateHashedPassword(utils.RandomString(8))

	if err != nil {
		log.Fatal("Unable to generate hash password", err)
	}

	arg := db.CreateUserParams{
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
	}

	user, err := testQuery.CreateUser(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.NotZero(t, user.ID)
	assert.Equal(t, arg.Email, user.Email)
	assert.Equal(t, arg.HashedPassword, user.HashedPassword)

	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2*time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)

	return user
}

func TestCreateUser(t *testing.T) {
	user1 := createRandomUser(t)
	defer clean_up()

	user2, err := testQuery.CreateUser(context.Background(), db.CreateUserParams{
		Email:          user1.Email,
		HashedPassword: user1.HashedPassword,
	})
	assert.Error(t, err)
	assert.Empty(t, user2)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)
	defer clean_up()

	newPassword, err := utils.GenerateHashedPassword(utils.RandomString(8))

	if err != nil {
		log.Fatal("Unable to generate hash password", err)
	}

	arg := db.UpdateUserPasswordParams{
		HashedPassword: newPassword,
		ID:             user.ID,
		UpdatedAt:      time.Now(),
	}

	updatedUser, err := testQuery.UpdateUserPassword(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, updatedUser)
	assert.Equal(t, updatedUser.Email, user.Email)
	assert.Equal(t, arg.HashedPassword, updatedUser.HashedPassword)
}

func TestGetUserByID(t *testing.T) {
	user := createRandomUser(t)
	defer clean_up()

	queried_user, err := testQuery.GetUserByID(context.Background(), user.ID)

	if err != nil {
		log.Fatal("Unable to get user by ID", err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, queried_user)
	assert.Equal(t, user.ID, queried_user.ID)
	assert.Equal(t, user.Email, queried_user.Email)
	assert.Equal(t, user.HashedPassword, queried_user.HashedPassword)
}

func TestGetUserByEmail(t *testing.T) {
	user := createRandomUser(t)
	defer clean_up()

	queried_user, err := testQuery.GetUserByEmail(context.Background(), user.Email)

	if err != nil {
		log.Fatal("Unable to get user by email", err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, queried_user)
	assert.Equal(t, user.ID, queried_user.ID)
	assert.Equal(t, user.Email, queried_user.Email)
	assert.Equal(t, user.HashedPassword, queried_user.HashedPassword)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	defer clean_up()

	err := testQuery.DeleteUser(context.Background(), user.ID)

	if err != nil {
		log.Fatal("Unable to delete user", err)
	}

	assert.NoError(t, err)

	// Check if the deleted user is still saved on the database
	deleted_user, err := testQuery.GetUserByID(context.Background(), user.ID)

	assert.Error(t, err)
	assert.Empty(t, deleted_user)
}

func TestListUsers(t *testing.T) {
	defer clean_up()
	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			createRandomUser(t)
		}()
	}

	wg.Wait()

	arg := db.ListUsersParams{
		Offset: 0,
		Limit:  30,
	}

	users, err := testQuery.ListUsers(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	assert.Len(t, users, 30)
}
