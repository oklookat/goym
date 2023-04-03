package auth

import (
	"context"
	"errors"
)

/**
Большинство описаний взяты из документации Яндекса.
https://yandex.ru/dev/id/doc/dg
**/

const (
	_errPrefix = "goym/auth: "

	// ID приложения Яндекс.Музыки для Windows.
	_clientID string = "23cabbbdc6cd418abb4b39c32c41195d"

	// Secret приложения Яндекс.Музыки для Windows.
	_clientSecret string = "53bc75238f0c4d08a118e51fe9203300"

	_tokenEndpoint = "https://oauth.yandex.ru/token"

	_codeEndpoint = "https://oauth.yandex.ru/device/code"
)

var (
	ErrInvalidGrant       = errors.New(_errPrefix + "incorrect or expired confirmation code")
	ErrTokensRefreshAfter = errors.New(_errPrefix + "empty Tokens.RefreshAfter (broken token?)")
	// bad
	ErrBrokenTokensErr = errors.New(_errPrefix + "statusCode != 200, but tokensError is empty (API changed?)")
	ErrBrokenClient    = errors.New(_errPrefix + "broken client_id or client_secret (OAuth App changed?)")
)

// 1. Приложение запрашивает два кода — device_code для устройства и user_code для пользователя.
// Время жизни предоставленных кодов — 10 минут. По истечении этого времени коды нужно запросить заново.
//
// 2. Приложение одновременно:
//
// - предлагает пользователю ввести user_code на странице Авторизация на устройстве;
//
// - начинает периодически запрашивать OAuth-токен, передавая device_code.
//
// 3. Пользователь вводит правильный код до истечения времени его жизни.
//
// 4. Яндекс.OAuth возвращает токен в ответ на следующий запрос приложения.
//
// hostnamePostfix:
//
// Устанавливает постфикс имени устройства, которое будет отображаться
// в аккаунте Яндекса после авторизации.
// Например, если постфикс будет "hello",
// то имя устройства будет таким: "имяустройства (hello)".
//
// Если постфикс не указан, будет использоваться значение по умолчанию.
func New(
	ctx context.Context,
	login string,
	onUrlCode func(url string, code string),
	hostnamePostfix *string) (*Tokens, error) {

	if ctx == nil || onUrlCode == nil {
		return nil, nil
	}

	if hostnamePostfix == nil || len(*hostnamePostfix) == 0 {
		_hostnamePostfix = "goym"
	}

	// запрашиваем коды.
	firstStep := confirmationCodes{}
	if err := firstStep.New(login); err != nil {
		return nil, err
	}

	codes, err := firstStep.Send(ctx)
	if err != nil {
		return nil, err
	}

	// пользователь идет вводить код на странице Яндекса...
	go onUrlCode(codes.VerificationUrl, codes.UserCode)

	// проверяем ввод. Если пользователь ввел верный код, выдаем токен.
	tokens := &Tokens{}
	err = tokens.Request(ctx, codes)

	_hostnamePostfix = ""
	return tokens, err
}
