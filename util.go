package goym

import (
	"context"
	"errors"
	"net/url"

	"github.com/oklookat/goym/schema"
	"github.com/oklookat/vantuz"
)

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
	// Странная ошибка.
	ErrNilResponse = errors.New("nil http or schema response")
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
		return schema.NewErrWithStatusCode(resp.StatusCode)
	}
	return data.Error
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
