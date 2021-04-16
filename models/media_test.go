package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertNewMedia(t *testing.T) {
	err := DeleteAllUsers()
	assert.NoError(t, err)

	err = DeleteAllMedias()
	assert.NoError(t, err)

	testUser, err := CreateTestUser()
	assert.NoError(t, err)

	mediaTest := &Media{
		ID:       uuid.New(),
		Name:     "Testing Picture.png",
		Type:     "image/png",
		OwnerID:  testUser.ID,
		FileHash: "0a4a712e4dceafd5b96b2ddb6372cd19ef94a6ab79a04a210682f73ba763dd14",
		Size:     420,
	}

	t.Run("Should successfully insert a new media", func(t *testing.T) {
		_, err := mediaTest.InsertNewMedia()
		assert.NoError(t, err, "Calling InsertNewMedia should have no error")
	})
}

func TestMediaHashExist(t *testing.T) {
	err := DeleteAllUsers()
	assert.NoError(t, err)

	err = DeleteAllMedias()
	assert.NoError(t, err)

	testUser, err := CreateTestUser()
	assert.NoError(t, err)

	mediaTest := &Media{
		ID:       uuid.New(),
		Name:     "Testing Picture.png",
		Type:     "image/png",
		OwnerID:  testUser.ID,
		FileHash: "0a4a712e4dceafd5b96b2ddb6372cd19ef94a6ab79a04a210682f73ba763dd14",
		Size:     420,
	}

	t.Run("Should return true with no error if there's an existing picture with the same hash under an user", func(t *testing.T) {
		_, err := mediaTest.InsertNewMedia()
		assert.NoError(t, err)

		result, err := MediaHashExist(mediaTest.OwnerID, mediaTest.FileHash)
		assert.NoError(t, err, "Calling MediaHashExist should have no error")
		assert.True(t, result, "Expect MediaHashExist to be true")
	})

	t.Run("Should return false with no error if there's no matching picture for a filehash and ownerid", func(t *testing.T) {
		result, err := MediaHashExist(mediaTest.OwnerID, "not-a-valid-hash")
		assert.NoError(t, err, "Calling MediaHashExist should have no error in that case")
		assert.False(t, result, "Expect MediaHashExist to be false")
	})
}

func TestGetMediaByID(t *testing.T) {
	err := DeleteAllUsers()
	assert.NoError(t, err)

	err = DeleteAllMedias()
	assert.NoError(t, err)

	testUser, err := CreateTestUser()
	assert.NoError(t, err)

	mediaTest := &Media{
		ID:       uuid.New(),
		Name:     "Darude sandstorm",
		Type:     "image/png",
		OwnerID:  testUser.ID,
		FileHash: "0a4a712e4dceafd5b96b2ddb6372cd19ef94a6ab79a04a210682f73ba763dd14",
		Size:     420,
	}

	t.Run("Should return a valid picture", func(t *testing.T) {
		mediaid, err := mediaTest.InsertNewMedia()
		assert.NoError(t, err)

		mediaInDatabase, err := GetMediaByID(mediaid, mediaTest.OwnerID)
		assert.NoError(t, err, "Calling GetMediaByID should have no error")
		assert.Equal(t, mediaInDatabase.ID, mediaTest.ID)
		assert.Equal(t, mediaInDatabase.FileHash, mediaTest.FileHash)
	})

	t.Run("Should return error with ErrNoRecord if media doesn't exist", func(t *testing.T) {
		err := DeleteAllMedias()
		assert.NoError(t, err)
		_, err = mediaTest.InsertNewMedia()
		assert.NoError(t, err)

		mediaInDatabase, err := GetMediaByID(uuid.MustParse("47a4c648-6d15-4d8b-8be6-49e55219b89d"), mediaTest.OwnerID)
		assert.Equal(t, &Media{}, mediaInDatabase, "Returned media should be nil in that case")
		assert.EqualError(t, err, ErrRecordNotFound.Error(), "Returned error should be \"Record not found\"")
	})
}
