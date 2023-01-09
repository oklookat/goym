package goymauth

import (
	"context"
	"errors"
)

/**
Большинство описаний взяты из документации Яндекса.
https://yandex.ru/dev/id/doc/dg
**/

var (
	// Авторизация отменена с помощью контекста.
	ErrCancelled = errors.New("auth cancelled by context")
)

const (
	// ID приложения Яндекс.Музыки для Windows.
	client_id string = "23cabbbdc6cd418abb4b39c32c41195d"

	// Secret приложения Яндекс.Музыки для Windows.
	client_secret string = "53bc75238f0c4d08a118e51fe9203300"

	token_endpoint = "https://oauth.yandex.ru/token"

	code_endpoint = "https://oauth.yandex.ru/device/code"
)

// Последовательность получения токена в этом случае:
//
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
// ========
//
// ctx - если передадите context with cancel, и отмените,
// то при следующем запросе к Яндексу, авторизация будет отменена. Будет возвращена ошибка ErrCancelled.
func New(ctx context.Context,
	login string,
	code func(url string, code string)) (*Tokens, error) {
	if ctx == nil {
		return nil, errors.New("nil ctx")
	}
	if code == nil {
		return nil, errors.New("nil code")
	}

	// запрашиваем коды.
	var firstStep = confirmationCodes{}
	if err := firstStep.New(login); err != nil {
		return nil, err
	}

	codes, err := firstStep.Send(ctx)
	if err != nil {
		return nil, err
	}

	// пользователь идет вводить код на странице Яндекса...
	go code(codes.VerificationUrl, codes.UserCode)

	// проверяем ввод. Если пользователь ввел верный код, выдаем токен.
	var tokens = &Tokens{}
	err = tokens.Request(ctx, codes)
	return tokens, err
}
