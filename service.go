package goym

import (
	"fmt"
	"strconv"

	"github.com/oklookat/goym/schema"
	"github.com/oklookat/goym/vantuz"
)

// int в строку (десятичная система).
func i2s[T int | int64 | int32](val T) string {
	return strconv.FormatInt(int64(val), 10)
}

// строку в int64 (десятичная система).
func s2i64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

// Пример:
//
// genApiPath([]string{"users", i2s(1234), "playlists", "list"})
//
// Результат: https://api.music.yandex.net/users/1234/playlists/list
func genApiPath(paths []string) string {
	if paths == nil {
		return ""
	}

	base := schema.ApiUrl
	for i := range paths {
		base += "/" + paths[i]
	}

	return base
}

// Проверить Response на наличие ошибки (поле Error).
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
	if data.Error.Message == "session-expired" {
		return ErrTokensExpired
	}
	return fmt.Errorf(errPrefix+"%v: %v", data.Error.Name, data.Error.Message)
}
