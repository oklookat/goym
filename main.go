package goym

import (
	"context"
	"errors"

	"github.com/oklookat/goym/auth"
	"github.com/oklookat/goym/schema"
	"github.com/oklookat/goym/vantuz"
)

const (
	errPrefix      = "goym: "
	errNotProvided = " not provided"
)

var (
	ErrNilTokens   = errors.New(errPrefix + "tokens" + errNotProvided)
	ErrNilAlbum    = errors.New(errPrefix + "album" + errNotProvided)
	ErrNilArtist   = errors.New(errPrefix + "artist" + errNotProvided)
	ErrNilPlaylist = errors.New(errPrefix + "playlist" + errNotProvided)
	ErrNilTracks   = errors.New(errPrefix + "tracks" + errNotProvided)
	ErrNilTrack    = errors.New(errPrefix + "track" + errNotProvided)
	ErrNilTrackIds = errors.New(errPrefix + "track ids" + errNotProvided)
	ErrNilAlbumIds = errors.New(errPrefix + "album ids" + errNotProvided)
	ErrNilUidKind  = errors.New(errPrefix + "uid-kind map" + errNotProvided)
	ErrNilStation  = errors.New(errPrefix + "station" + errNotProvided)
	//
	ErrNilHttpResponse  = errors.New(errPrefix + "nil http.response (dev error?)")
	ErrNilResponse      = errors.New(errPrefix + "nil Response (dev error?)")
	ErrNilResponseError = errors.New(errPrefix + "nil Response.Error (API changed?)")
	ErrNilStatus        = errors.New(errPrefix + "nil Status (bad auth or API changed?)")
	ErrNilAccount       = errors.New(errPrefix + "nil Status.Account (API changed?)")
)

// Получить Client для запросов к API.
//
// Получить tokens можно войдя в аккаунт, используя пакет auth.
func New(tokens *auth.Tokens) (*Client, error) {
	if tokens == nil {
		return nil, ErrNilTokens
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
	status, err := cl.GetAccountStatus(context.Background())
	if err != nil {
		return nil, err
	}
	if status == nil {
		return nil, ErrNilStatus
	}

	if status.Account == nil {
		return nil, ErrNilAccount
	}

	cl.UserId = status.Account.UID
	cl.userId = cl.UserId.String()
	return cl, err
}

// Клиент для запросов к API.
type Client struct {
	// ID текущего пользователя.
	UserId schema.UniqueID

	// Для создания эндпоинтов.
	// Чтоб не конвертировать по 100 раз UserId.
	userId string

	// Отправляет запросы.
	self *vantuz.Client
}

// Включить вывод HTTP запросов в консоль.
func (c Client) EnableDevMode() {
	c.self.EnableDevMode()
}

// Отключить вывод HTTP запросов в консоль.
func (c Client) DisableDevMode() {
	c.self.DisableDevMode()
}
