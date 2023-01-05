package auth

import (
	"context"
	"errors"
	"time"

	"github.com/oklookat/goym/holly"
)

const (
	// Пользователь еще не ввел код подтверждения.
	errAuthorizationPending = "authorization_pending"

	// Приложение с указанным идентификатором (параметр client_id) не найдено или заблокировано.
	//
	// Этот код также возвращается, если в параметре client_secret передан неверный пароль приложения.
	//
	// P.S: в нашем случае эта ошибка может появиться,
	// если Яндекс сменит коды (id, secret) для своего приложения под Windows.
	//
	// В таком случае надо брать в руки анализатор трафика,
	// и идти искать новые коды.
	errInvalidClient = "invalid_client"

	// Неверный или просроченный код подтверждения.
	errInvalidGrant = "invalid_grant"
)

// Если выдать токен не удалось, то ответ содержит описание ошибки.
type tokensError struct {
	// Описание ошибки.
	ErrorDescription string `json:"error_description"`

	// Код ошибки.
	Error string `json:"error"`
}

// Яндекс.OAuth возвращает OAuth-токен, refresh-токен и время их жизни в JSON-формате.
//
// https://yandex.ru/dev/id/doc/dg/oauth/reference/simple-input-client.html#simple-input-client__token-body-title
type Tokens struct {
	// Тип выданного токена. Всегда принимает значение «bearer».
	TokenType string `json:"token_type"`

	// OAuth-токен с запрошенными правами или с правами, указанными при регистрации приложения.
	AccessToken string `json:"access_token"`

	// Время жизни токена в секундах.
	ExpiresIn int64 `json:"expires_in"`

	// Токен, который можно использовать для продления срока жизни соответствующего OAuth-токена.
	// Время жизни refresh-токена совпадает с временем жизни OAuth-токена.
	RefreshToken string `json:"refresh_token"`

	// Это поле не входит в ответ Яндекса.
	//
	// Дата в формате unix.
	// После этой даты надо обновить токены.
	RefreshAfter int64 `json:"refresh_after"`
}

// Приложение начинает периодически запрашивать OAuth-токен, передавая device_code.
func (t *Tokens) Request(ctx context.Context, codes *confirmationCodesResponse) (err error) {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if codes == nil {
		err = errors.New("nil codes")
		return
	}

	var form = map[string]string{
		// Способ запроса OAuth-токена.
		// Если вы используете код подтверждения, укажите значение «authorization_code».
		"grant_type": "device_code",

		// Код подтверждения, полученный от Яндекс.OAuth.
		// Время жизни предоставленного кода — 10 минут. По истечении этого времени код нужно запросить заново.
		"code": codes.DeviceCode,

		"client_id":     client_id,
		"client_secret": client_secret,
	}

	var resp *holly.Response
	var tokensErr = &tokensError{}
	var request = holly.C().R().
		SetFormData(form).
		SetResult(t).SetError(tokensErr)

	// таймер когда токены истекут
	var expiredDur = time.Duration(codes.ExpiresIn-8) * time.Second
	var tokensExpired = time.NewTimer(expiredDur)
	defer tokensExpired.Stop()

	// ждем *интервал* перед отправкой нового запроса...
	// +2 секунды на всякий случай
	var sleepFor = time.Duration(codes.Interval+2) * time.Second
	var requestSleep = time.NewTicker(sleepFor)
	defer requestSleep.Stop()

	for {
		select {
		case <-tokensExpired.C:
			err = errors.New("tokens expired")
			return
		case <-requestSleep.C:

			if ctx.Err() != nil {
				err = ErrCancelled
				return
			}

			if resp, err = request.Post(token_endpoint); err != nil {
				return
			}

			if resp.IsSuccess() {
				t.setRefreshAfter()
				return
			}

			if len(tokensErr.Error) < 1 {
				err = errors.New("statusCode != 200, but tokensErr is empty")
				return
			}

			switch tokensErr.Error {
			default:
				err = errors.New(tokensErr.Error)
				return
			case errAuthorizationPending:
				continue
			case errInvalidClient:
				err = errors.New("broken client_id or client_secret")
				return
			case errInvalidGrant:
				err = errors.New("incorrect or expired confirmation code")
				return
			}
		}
	}
}

// Обновить токены.
//
// Если обновление не требуется, возвращает nil.
//
// https://yandex.ru/dev/id/doc/dg/oauth/reference/refresh-client.html
func (t *Tokens) Refresh() (refreshed *Tokens, err error) {
	if t.RefreshAfter < 1 {
		err = errors.New("empty RefreshAfter (broken token?)")
		return
	}

	if !t.isNeedToRefresh() {
		return
	}

	var form = map[string]string{
		// Способ запроса OAuth-токена.
		// Если вы используете refresh-токен, укажите значение «refresh_token»
		"grant_type": "refresh_token",

		// Refresh-токен, полученный от Яндекс.OAuth вместе с OAuth-токеном. Время жизни токенов совпадает.
		"refresh_token": t.RefreshToken,
		"client_id":     client_id,
		"client_secret": client_secret,
	}

	var tokenErr = &tokensError{}
	var request = holly.C().R().
		SetFormData(form).
		SetResult(refreshed).SetError(tokenErr)

	var resp *holly.Response
	if resp, err = request.Post(token_endpoint); err != nil {
		refreshed = nil
		return
	}

	if resp.IsError() {
		err = errors.New(tokenErr.Error)
		return
	}

	if refreshed == nil {
		err = errors.New("nil tokens")
		return
	}

	refreshed.setRefreshAfter()

	return
}

// Устанавливает RefreshAfter.
func (t *Tokens) setRefreshAfter() {
	var now = time.Now()

	// Основной токен может не обновиться, если оставшийся срок его жизни достаточно длительный
	// и выдавать новый токен нет необходимости. Рекомендуем обновлять долгоживущие токены раз в три месяца.
	var after = now.AddDate(0, 3, 5)

	t.RefreshAfter = after.Unix()
}

// Сравнивает текущую дату и RefreshAfter.
func (t *Tokens) isNeedToRefresh() bool {
	var now = time.Now().Unix()
	return now > t.RefreshAfter
}
