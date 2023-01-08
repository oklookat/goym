package goym

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/oklookat/goym/vantuz"
)

// int в строку (десятичная система).
func i2s[T int | int64 | int32 | uint](val T) string {
	return strconv.FormatInt(int64(val), 10)
}

// строку в int64 (десятичная система).
func s2i64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

// bool в строку.
func b2s(val bool) string {
	return strconv.FormatBool(val)
}

// true = VisibilityPublic
//
// false = VisibilityPrivate
func visibilityToString(public bool) string {
	var visibility = VisibilityPrivate
	if public {
		visibility = VisibilityPublic
	}
	return visibility
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

// Например:
//
// val = map[int64]int64{29238706: 7019818, 29238706: 83063895}
//
// результат: "29238706:7019818,29238706:83063895"
// func i64ColonJoin(data map[int64]int64) string {
// 	if data == nil {
// 		return ""
// 	}

// 	var pairs = []string{}
// 	for id, albumId := range data {
// 		var idStr = i2s(id)
// 		var albumIdStr = i2s(albumId)
// 		pairs = append(pairs, idStr+":"+albumIdStr)
// 	}

// 	return strings.Join(pairs, ",")
// }

// Проверить TypicalResponse на наличие ошибки (поле Error).
//
// Если ошибка есть, возвращает error с сообщением.
func checkTypicalResponse[T any](resp *vantuz.Response, data *TypicalResponse[T]) error {
	if resp == nil {
		return errors.New("nil response")
	}
	if data == nil {
		return errors.New("nil data")
	}
	if resp.IsSuccess() {
		return nil
	}
	if data.Error == nil {
		return errors.New("nil data.Error")
	}
	return fmt.Errorf("%v: %v", data.Error.Name, data.Error.Message)
}
