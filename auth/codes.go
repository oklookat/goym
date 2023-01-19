package auth

import (
	"context"
	"fmt"

	"github.com/oklookat/goym/vantuz"
)

// Первый шаг.
//
// Приложение запрашивает два кода — device_code для устройства и user_code для пользователя.
//
// Время жизни предоставленных кодов — 10 минут. По истечении этого времени коды нужно запросить заново.
//
// https://yandex.ru/dev/id/doc/dg/oauth/reference/simple-input-client.html#simple-input-client__get-codes
type confirmationCodes struct {
	// Идентификатор приложения.
	// Доступен в свойствах приложения (нажмите название приложения, чтобы открыть его свойства).
	clientId string

	// Уникальный идентификатор устройства, для которого запрашивается токен.
	// Чтобы обеспечить уникальность, достаточно один раз сгенерировать UUID
	// и использовать его при каждом запросе нового токена с данного устройства.
	// Идентификатор должен быть не короче 6 символов и не длиннее 50.
	// Допускается использовать только печатаемые ASCII-символы (с кодами от 32 до 126).
	deviceId string

	// Имя устройства, которое следует показывать пользователям. Не длиннее 100 символов.
	deviceName string

	// Логин или почта на Яндексе.
	login string

	// Был ли вызван метод New()?
	isNewCalled bool

	// Форма для отправки POST запроса. Создается при вызове New().
	form map[string]string
}

// Яндекс.OAuth возвращает код для пользователя и информацию для запроса токена.
type confirmationCodesResponse struct {
	// Код, с которым следует запрашивать OAuth-токен на следующем шаге.
	DeviceCode string `json:"device_code"`

	// Код, который должен ввести пользователь, чтобы разрешить доступ к своим данным.
	UserCode string `json:"user_code"`

	// Адрес страницы, на которой пользователь должен ввести код из свойства user_code.
	VerificationUrl string `json:"verification_url"`

	// Минимальный интервал, с которым приложение должно запрашивать OAuth-токен.
	// Если запросы будут приходить чаще, Яндекс.OAuth может ответить ошибкой.
	Interval uint8 `json:"interval"`

	// Срок действия пары кодов.
	// По истечению этого срока получить токен для них будет невозможно — нужно будет начать процедуру сначала.
	ExpiresIn uint32 `json:"expires_in"`
}

func (c *confirmationCodes) New(login string) (err error) {
	c.clientId = client_id
	c.login = login
	c.deviceId = generateUlid()
	c.deviceName, err = getHostname()
	c.makeForm()
	c.isNewCalled = true
	return
}

// Отправить запрос.
func (c confirmationCodes) Send(ctx context.Context) (*confirmationCodesResponse, error) {
	if !c.isNewCalled {
		return nil, nil
	}

	var codes = &confirmationCodesResponse{}
	var tokensErr = &tokensError{}
	var request = vantuz.C().R().
		SetFormUrlMap(c.form).
		SetResult(codes).
		SetError(tokensErr)

	if ctx.Err() != nil {
		return nil, ErrCancelled
	}

	resp, err := request.Post(ctx, code_endpoint)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		err = fmt.Errorf(errPrefix+"confirmCodes (%s)", tokensErr.Error)
	}

	return codes, err
}

// Создать форму, чтобы отправить ее вместе с запросом.
func (c *confirmationCodes) makeForm() {
	var form = make(map[string]string)
	form["client_id"] = c.clientId
	form["device_id"] = c.deviceId
	form["device_name"] = c.deviceName
	c.form = form
}
