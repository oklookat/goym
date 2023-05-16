package goym

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oklookat/goym/schema"
	"github.com/oklookat/goym/vantuz"
)

// iToString преобразует число в строку (десятичная система).
func iToString[T int | int64 | int32](val T) string {
	return strconv.FormatInt(int64(val), 10)
}

// stringToInt64 преобразует строку в int64 (десятичная система).
func stringToInt64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
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
