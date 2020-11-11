package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertNewSong(t *testing.T) {
	err := DeleteAllUsers()
	assert.NoError(t, err)
	err = DeleteAllSongs()
	assert.NoError(t, err)

	testUser, err := CreateTestUser()
	assert.NoError(t, err)

	songTest := &Song{
		Name:     "Darude sandstorm",
		ObjectID: uuid.New().String(),
		Ext:      "mp3",
		OwnerID:  testUser.ID,
		FileHash: "0a4a712e4dceafd5b96b2ddb6372cd19ef94a6ab79a04a210682f73ba763dd14",
		Size:     420,
	}

	t.Run("Should successfully insert a new song", func(t *testing.T) {
		_, err := songTest.InsertNewSong()
		assert.NoError(t, err, "Calling InsertNewSong should have no error")
	})
}

func TestSongHashExist(t *testing.T) {
	err := DeleteAllUsers()
	assert.NoError(t, err)
	err = DeleteAllSongs()
	assert.NoError(t, err)

	testUser, err := CreateTestUser()
	assert.NoError(t, err)

	songTest := &Song{
		Name:     "Darude sandstorm",
		ObjectID: uuid.New().String(),
		Ext:      "mp3",
		OwnerID:  testUser.ID,
		FileHash: "0a4a712e4dceafd5b96b2ddb6372cd19ef94a6ab79a04a210682f73ba763dd14",
		Size:     420,
	}

	t.Run("Should return true with no error if there's an existing song with the same hash under an user", func(t *testing.T) {
		_, err := songTest.InsertNewSong()
		assert.NoError(t, err)

		result, err := SongHashExist(songTest.OwnerID, songTest.FileHash)
		assert.NoError(t, err, "Calling SongHashExist should have no error")
		assert.True(t, result, "Expect SongHashExist to be true")
	})

	t.Run("Should return false with no error if there's no matching song for a filehash and ownerid", func(t *testing.T) {
		result, err := SongHashExist(songTest.OwnerID, "no-valid-hash")
		assert.NoError(t, err, "Calling SongHashExist should have no error in that case")
		assert.False(t, result, "Expect SongHashExist to be false")
	})
}

func TestGetSongByID(t *testing.T) {
	err := DeleteAllUsers()
	assert.NoError(t, err)
	err = DeleteAllSongs()
	assert.NoError(t, err)

	testUser, err := CreateTestUser()
	assert.NoError(t, err)

	songTest := &Song{
		Name:     "Darude sandstorm",
		ObjectID: uuid.New().String(),
		Ext:      "mp3",
		OwnerID:  testUser.ID,
		FileHash: "0a4a712e4dceafd5b96b2ddb6372cd19ef94a6ab79a04a210682f73ba763dd14",
		Size:     420,
	}

	t.Run("Should return a valid song", func(t *testing.T) {
		songid, err := songTest.InsertNewSong()
		assert.NoError(t, err)

		songInDatabase, err := GetSongByID(songid, songTest.OwnerID)
		assert.NoError(t, err, "Calling GetSongByID should have no error")
		assert.Equal(t, songInDatabase.ObjectID, songTest.ObjectID)
		assert.Equal(t, songInDatabase.FileHash, songTest.FileHash)
	})

	t.Run("Should return error with ErrNoRecord if song doesn't exist", func(t *testing.T) {
		err := DeleteAllSongs()
		assert.NoError(t, err)
		_, err = songTest.InsertNewSong()
		assert.NoError(t, err)

		songInDatabase, err := GetSongByID(1234563, songTest.OwnerID)
		assert.Equal(t, &Song{}, songInDatabase, "Returned song should be nil in that case")
		assert.EqualError(t, err, ErrRecordNotFound.Error(), "Returned error should be \"Record not found\"")
	})
}
