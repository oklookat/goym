package goym

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/oklookat/goym/holly"
)

// int в строку (десятичная система).
func i2s[T int | int64 | int32](val T) string {
	return strconv.FormatInt(int64(val), 10)
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

	var base = ApiUrl
	for i := range paths {
		base += "/" + paths[i]
	}

	return base
}

// Например:
//
// val = []int64{7019818,29238706,83063895}
//
// результат: "7019818,29238706,83063895"
func i64Join(data []int64) string {
	if data == nil {
		return ""
	}

	var converted = make([]string, len(data))
	for i := range data {
		converted[i] = i2s(data[i])
	}

	return strings.Join(converted, ",")
}

// Проверить GetResponse на наличие ошибки. Если есть, возвращает error, в которой будет сообщение.
func checkGetResponse[T any](resp *holly.Response, data *GetResponse[T]) (err error) {
	if resp == nil {
		err = errors.New("nil response")
		return
	}
	if data == nil {
		err = errors.New("nil data")
		return
	}
	if resp.IsSuccess() {
		return
	}
	if data.Error == nil {
		err = errors.New("nil data.Error")
		return
	}
	err = fmt.Errorf("%v: %v", data.Error.Name, data.Error.Message)
	return
}
