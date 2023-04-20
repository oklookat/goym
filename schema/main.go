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

// Тип поиска.
type SearchType string

const (
	// Поиск артистов.
	SearchTypeArtist SearchType = "artist"

	// Поиск альбомов.
	SearchTypeAlbum SearchType = "album"

	// Поиск треков.
	SearchTypeTrack SearchType = "track"

	// Поиск подкастов.
	SearchTypePodcast SearchType = "podcast"

	// Поиск плейлистов.
	SearchTypePlaylist SearchType = "playlist"

	// Поиск всего.
	SearchTypeAll SearchType = "all"
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
		InvocationInfo InvocationInfo `json:"invocationInfo"`

		// Если не nil, то поле result будет nil.
		Error *Error `json:"error"`

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

	// Ошибка. Ошибка валидации, например.
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
		// Допустим, вы отправили запрос, где указали perPage = 20.
		//
		// Вам пришел ответ с этой структурой (Pager). И вот это поле (PerPage)
		// может не быть равным 20. Такое может быть, когда всего элементов не намного больше, чем perPage.
		//
		// Например вы указали perPage = 20, и пришел ответ, где Total равен 22.
		// В таком случае это поле (PerPage) будет равно 22.
		PerPage uint16 `json:"perPage"`

		// Всего элементов.
		Total uint16 `json:"total"`
	}
)

// Уникальный ID.
type ID uint64

func (u ID) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

func (u *ID) FromString(val string) error {
	res, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}
	*u = ID(res)
	return nil
}
