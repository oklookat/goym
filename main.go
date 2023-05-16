package goym

import (
	"context"
	"errors"

	"github.com/oklookat/goym/auth"
	"github.com/oklookat/goym/schema"
	"github.com/oklookat/goym/vantuz"
)

const (
	errPrefix = "goym: "
)

var (
	// Токены авторизации истекли. Нужно вызывать Tokens.Refresh.
	ErrTokensExpired = errors.New(errPrefix + "tokens expired. You need to refresh your current tokens")

	//
	ErrNilResponse      = errors.New(errPrefix + "nil http or schema response (dev error?)")
	ErrNilResponseError = errors.New(errPrefix + "nil Response.Error (API changed?)")
	ErrNilStatus        = errors.New(errPrefix + "nil Status (bad auth or API changed?)")
)

// Получить Client для запросов к API.
//
// Получить tokens можно войдя в аккаунт, используя пакет auth.
//
// ==== accountID
//
// Для запросов к API помимо токенов нужен ID аккаунта, которому они принадлежат.
//
// Если accountID будет nil, то при вызове метода New будет запрос к API для получения информации об аккаунте, чтобы получить ID.
//
// В целом этот аргумент нужен чтобы не делать лишних запросов к API, если вы уже знаете ID аккаунта которому принадлежат токены.
func New(tokens *auth.Tokens, accountID *schema.ID) (*Client, error) {
	if tokens == nil {
		return nil, nil
	}

	httpCl := vantuz.C().
		SetGlobalHeaders(map[string]string{
			"Authorization": "OAuth " + tokens.AccessToken,
		})
	cl := &Client{
		Http: httpCl,
	}

	if accountID == nil {
		// get uid
		status, err := cl.AccountStatus(context.Background())
		if err != nil {
			return nil, err
		}
		cl.UserId = status.Result.Account.UID
	} else {
		cl.UserId = *accountID
	}

	return cl, nil
}

// Клиент для запросов к API.
type Client struct {
	// ID текущего пользователя.
	UserId schema.ID

	// Отправляет запросы.
	Http *vantuz.Client
}
