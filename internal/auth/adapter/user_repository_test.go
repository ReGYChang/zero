package adapter

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"zero/internal/auth/domain"
	"zero/test"
)

func assertUser(t *testing.T, expected *domain.User, actual *domain.User) {
	require.NotNil(t, actual)
	assert.Equal(t, expected.UID, actual.UID)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.Name, actual.Name)
}

func TestPostgresRepository_CreateUser(t *testing.T) {
	db := getTestPostgresDB()
	repo := initRepository(t, db)

	// Args
	type Args struct {
		User domain.User
	}
	var args Args
	_ = faker.FakeData(&args)

	user, err := repo.CreateUser(context.Background(), args.User)

	require.NoError(t, err)
	assertUser(t, &args.User, user)

	// No duplicate
	_, err = repo.CreateUser(context.Background(), args.User)
	require.Error(t, err)
}

func TestPostgresRepository_GetUserByEmail(t *testing.T) {
	db := getTestPostgresDB()
	repo := initRepository(t, db, test.Path(test.TestDataUser))

	email := "user1@cresclab.com"

	_, err := repo.GetUserByEmail(context.Background(), email)
	require.NoError(t, err)
}
