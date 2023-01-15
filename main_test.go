package goym

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/oklookat/goym/auth"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestAll(t *testing.T) {
	TestService(t)
	TestAccount(t)
	TestAlbum(t)
	TestArtist(t)
	TestPlaylist(t)
	TestSearch(t)
	TestTrack(t)
	TestRotor(t)
}

func TestService(t *testing.T) {
	suite.Run(t, &ServiceTestSuite{})
}

func TestAccount(t *testing.T) {
	suite.Run(t, &AccountTestSuite{})
}

func TestAlbum(t *testing.T) {
	suite.Run(t, &AlbumTestSuite{})
}

func TestArtist(t *testing.T) {
	suite.Run(t, &ArtistTestSuite{})
}

func TestPlaylist(t *testing.T) {
	suite.Run(t, &PlaylistTestSuite{})
}

func TestSearch(t *testing.T) {
	suite.Run(t, &SearchTestSuite{})
}

func TestTrack(t *testing.T) {
	suite.Run(t, &TrackTestSuite{})
}

func TestRotor(t *testing.T) {
	suite.Run(t, &RotorTestSuite{})
}

// Получить клиент для запросов к API.
func getClient(t *testing.T) *Client {
	require := require.New(t)
	err := godotenv.Load()
	require.Nil(err)

	expiresIn, err := s2i64(os.Getenv("EXPIRES_IN"))
	require.Nil(err)
	refreshAfter, err := s2i64(os.Getenv("REFRESH_AFTER"))
	require.Nil(err)

	var tok = &auth.Tokens{
		TokenType:    os.Getenv("TOKEN_TYPE"),
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		ExpiresIn:    expiresIn,
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
		RefreshAfter: refreshAfter,
	}

	cl, err := New(tok)
	require.Nil(err)

	//cl.EnableDevMode()
	return cl
}
