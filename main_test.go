package goym

import (
	"log"
	"os"
	"testing"
	"time"

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

	expiresIn, err := stringToInt64(os.Getenv("EXPIRES_IN"))
	require.Nil(err)
	refreshAfter, err := stringToInt64(os.Getenv("REFRESH_AFTER"))
	require.Nil(err)

	tok := &auth.Tokens{
		TokenType:    os.Getenv("TOKEN_TYPE"),
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		ExpiresIn:    expiresIn,
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
		RefreshAfter: refreshAfter,
	}

	cl, err := New(tok)
	require.Nil(err)

	cl.Http.SetLogger(loggerDefault{})
	cl.Http.SetRateLimit(1, time.Duration(1)*time.Second)

	return cl
}

type loggerDefault struct {
}

func (l loggerDefault) Debugf(msg string, args ...any) {
	log.Printf(msg, args...)
}

func (l loggerDefault) Err(msg string, err error) {
	if err == nil {
		log.Printf("%s", msg)
		return
	}
	log.Printf("%s. Err: %s", msg, err.Error())
}
