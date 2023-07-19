package goym

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/oklookat/goym/schema"
	"github.com/oklookat/vantuz"
)

// iToString преобразует число в строку (десятичная система).
func iToString[T int | int64 | int32](val T) string {
	return strconv.FormatInt(int64(val), 10)
}

// genApiPath создает URL-адрес для запроса к API, используя заданный путь.
//
// Пример использования: genApiPath([]string{"users", iToString(1234), "playlists", "list"})
//
// Результат: https://api.music.yandex.net/users/1234/playlists/list
func genApiPath(paths ...string) string {
	if len(paths) == 0 {
		return schema.ApiUrl
	}

	base := schema.ApiUrl + "/" + paths[0]
	for i := 1; i < len(paths); i++ {
		base += "/" + paths[i]
	}

	return base
}

var (
	// Нужно обновить токен.
	ErrTokensExpired = errors.New(errPrefix + "tokens expired. You need to refresh tokens")

	// Странная ошибка.
	ErrNilResponse = errors.New(errPrefix + "nil http or schema response (???)")

	// Ответ с ошибкой, но поля Error в ответе нет.
	ErrNilResponseError = errors.New(errPrefix + "nil Response.Error (API changed?)")
)

// checkResponse проверяет наличие ошибки в ответе API.
// Возвращает nil, если ошибки нет.
//
// Если ошибка есть, возвращает error с сообщением.
func checkResponse[T any](resp *vantuz.Response, data *schema.Response[T]) error {
	if resp == nil || data == nil {
		return ErrNilResponse
	}
	if resp.IsSuccess() {
		return nil
	}
	if data.Error == nil {
		return ErrNilResponseError
	}
	if strings.EqualFold(data.Error.Message, "session-expired") {
		return ErrTokensExpired
	}
	return fmt.Errorf(errPrefix+"%v: %v", data.Error.Name, data.Error.Message)
}

func addRemoveMultiple(
	ctx context.Context,
	client *Client,
	vals url.Values,
	add bool,
	entityName string,
) (schema.Response[string], error) {
	// POST /users/{userId}/likes/ENTITIES/add-multiple
	// ||
	// POST /users/{userId}/likes/ENTITIES/remove
	endEndPoint := "add-multiple"
	if !add {
		endEndPoint = "remove"
	}
	endpoint := genApiPath("users", string(client.UserId), "likes", entityName, endEndPoint)
	data := &schema.Response[string]{}

	resp, err := client.Http.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// If add: use "id" request body.
//
// If remove: use "ids" request body.
func addRemove(
	ctx context.Context,
	client *Client,
	vals url.Values,
	add bool,
	entityName string) (schema.Response[string], error) {
	// POST /users/{userId}/likes/ENTITIES/add
	// ||
	// POST /users/{userId}/likes/ENTITIES/remove
	endEndPoint := "add"
	if !add {
		endEndPoint = "remove"
	}

	endpoint := genApiPath("users", string(client.UserId), "likes", entityName, endEndPoint)
	data := &schema.Response[string]{}
	resp, err := client.Http.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

func likesDislikes[T any](ctx context.Context, client *Client, likes bool, entityName string) (schema.Response[T], error) {
	// GET /users/{userId}/likes/ENTITY
	// ||
	// GET /users/{userId}/dislikes/ENTITY
	lord := "likes"
	if !likes {
		lord = "dislikes"
	}
	endpoint := genApiPath("users", string(client.UserId), lord, entityName)
	data := &schema.Response[T]{}

	resp, err := client.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return *data, err
}
