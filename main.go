package goym

import (
	"errors"

	"github.com/oklookat/goym/auth"
	"github.com/oklookat/goym/vantuz"
)

var (
	ErrNilAlbum    = errors.New("nil album")
	ErrNilArtist   = errors.New("nil artist")
	ErrNilPlaylist = errors.New("nil playlist")
	ErrNilTracks   = errors.New("nil tracks")
	ErrNilTrack    = errors.New("nil track")
	ErrNilTrackIds = errors.New("nil trackIds")
)

// Получить Client для запросов к API.
//
// Получить tokens можно войдя в аккаунт, используя пакет goymauth.
func New(tokens *auth.Tokens) (*Client, error) {
	if tokens == nil {
		return nil, errors.New("nil tokens")
	}

	var vCl = vantuz.C().
		SetGlobalHeaders(map[string]string{
			"User-Agent":    "oklookat/goym",
			"Authorization": "OAuth " + tokens.AccessToken,
		})
	var cl = &Client{
		self: vCl,
	}

	// get uid
	resp, err := cl.GetAccountStatus()
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("nil response")
	}

	if resp.Account == nil {
		return nil, errors.New("nil account")
	}

	cl.UserId = resp.Account.UID
	cl.userId = i2s(cl.UserId)
	return cl, err
}

// Клиент для запросов к API.
type Client struct {
	UserId int64

	// Для создания эндпоинтов.
	userId string
	self   *vantuz.Client
}

// Включить вывод HTTP запросов в консоль.
func (c Client) EnableDevMode() {
	c.self.EnableDevMode()
}

// Отключить вывод HTTP запросов в консоль.
func (c Client) DisableDevMode() {
	c.self.DisableDevMode()
}
