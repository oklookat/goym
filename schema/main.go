package schema

import (
	"errors"
	"strconv"
	"strings"
)

const (
	errPrefix = "goym/schema: "

	ApiUrl = "https://api.music.yandex.net"
)

var (
	ErrNilTrack    = errors.New(errPrefix + "nil track")
	ErrNilTracks   = errors.New(errPrefix + "nil tracks")
	ErrNilPlaylist = errors.New(errPrefix + "nil playlist")
)

// Сортировка по...
type SortBy string

func (e SortBy) String() string {
	return string(e)
}

const (
	// Сортировка по году.
	SortByYear SortBy = "year"

	// Сортировка по рейтингу.
	SortByRating SortBy = "rating"
)

// Сортировка по...
type SortOrder string

func (e SortOrder) String() string {
	return string(e)
}

const (
	// Сортировка по убыванию.
	SortOrderDesc SortOrder = "desc"

	// Сортировка по возрастанию.
	SortOrderAsc SortOrder = "asc"
)

// Тема.
type Theme string

const (
	// Темная.
	ThemeBlack Theme = "black"

	// Светлая.
	ThemeWhite Theme = "white"

	// По умолчанию.
	ThemeDefault Theme = "default"
)

// Видимость.
type Visibility string

func (e Visibility) String() string {
	return string(e)
}

const (
	// Приватная.
	VisibilityPrivate Visibility = "private"

	// Публичная.
	VisibilityPublic Visibility = "public"
)

type (
	// Обычно ответ выглядит так.
	Response[T any] struct {
		// Информация о запросе.
		InvocationInfo InvocationInfo `json:"invocationInfo"`

		// Если не nil, то поле result будет nil.
		Error *Error `json:"error"`

		// Результат запроса.
		Result T `json:"result"`

		// Может быть при некоторых запросах.
		// Например при получении лайкнутых плейлистов.
		Pager *Pager `json:"pager"`
	}

	// Что-то техническое.
	InvocationInfo struct {
		// Адрес какого-то сервера Яндекс.Музыки.
		Hostname string `json:"hostname"`

		// ID запроса.
		ReqID string `json:"req-id"`

		// (?) Время выполнения запроса в миллисекундах.
		//
		// string | int
		ExecDurationMillis any `json:"exec-duration-millis"`
	}

	// Информация о страницах.
	Pager struct {
		// Текущая страница.
		Page int `json:"page"`

		// Сколько элементов на странице.
		//
		// Обратите внимание:
		//
		// Допустим вы отправили запрос с perPage = 20.
		//
		// Вам пришел ответ с этой структурой (Pager). И вот это поле (PerPage)
		// может не быть равным 20. Такое может быть когда элементов чуть больше чем perPage.
		//
		// Например вы указали perPage = 20, и пришел ответ, где Total равен 22.
		// В таком случае это поле (PerPage) будет равно 22.
		PerPage int `json:"perPage"`

		// Общее кол-во элементов.
		Total int `json:"total"`
	}
)

type ID string

func (i ID) String() string {
	return string(i)
}

func (i *ID) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	str, err := strconv.Unquote(string(data))
	if err != nil {
		*i = ID(data)
		return nil
	}
	*i = ID(str)
	return nil
}

// Ошибка. Например ошибка валидации.
type Error struct {
	// Например: validate.
	Name string `json:"name"`

	// Например: Parameters requirements are not met.
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Name + ": " + e.Message
}

// Ошибка валидации может быть в этих случаях:
//
// 1. Изменилось API.
//
// 2. Я допустил ошибку.
//
// 3. Неверные данные в теле запроса.
// Например вы пытаетесь получить артиста с ID "-1", "0" и так далее.
// Хотя в таких случаях нужно отдавать 404, но имеем что имеем.
func (e Error) IsValidate() bool {
	return strings.EqualFold(e.Name, "validate")
}

// Вероятно если ID валидный (см. IsValidate),
// то в каких-то случаях может быть и ошибка not found.
func (e Error) IsNotFound() bool {
	return strings.Contains(e.Name, "not-found")
}

// Нужно обновить access token.
func (e Error) IsSessionExpired() bool {
	return strings.EqualFold(e.Message, "session-expired")
}

func NewErrWithStatusCode(statusCode int) ErrWithStatusCode {
	return ErrWithStatusCode{
		StatusCode: statusCode,
	}
}

type ErrWithStatusCode struct {
	StatusCode int
}

func (e ErrWithStatusCode) Error() string {
	return strconv.Itoa(e.StatusCode)
}
