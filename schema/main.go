package schema

import (
	"errors"
	"strconv"
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

const (
	// Сортировка по году.
	SortByYear SortBy = "year"

	// Сортировка по рейтингу.
	SortByRating SortBy = "rating"
)

// Сортировка по...
type SortOrder string

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

	// Ошибка. Например ошибка валидации.
	Error struct {
		// Например: validate.
		Name string `json:"name"`

		// Например: Parameters requirements are not met.
		Message string `json:"message"`
	}

	// Информация о страницах.
	Pager struct {
		// Текущая страница.
		Page uint16 `json:"page"`

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
		PerPage uint16 `json:"perPage"`

		// Общее кол-во элементов.
		Total uint16 `json:"total"`
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
