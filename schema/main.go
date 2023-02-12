package schema

import (
	"errors"
	"strconv"
)

const (
	errPrefix = "goym/schema: "

	ApiUrl = "https://api.music.yandex.net"

	VisibilityPrivate  Visibility = "private"
	VisibilityPublic   Visibility = "public"
	ThemeBlack         Theme      = "black"
	ThemeWhite         Theme      = "white"
	ThemeDefault       Theme      = "default"
	SearchTypeArtist   SearchType = "artist"
	SearchTypeAlbum    SearchType = "album"
	SearchTypeTrack    SearchType = "track"
	SearchTypePodcast  SearchType = "podcast"
	SearchTypePlaylist SearchType = "playlist"
	SearchTypeAll      SearchType = "all"
	SortByYear         SortBy     = "year"
	SortByRating       SortBy     = "rating"
)

var (
	ErrNilTrack    = errors.New(errPrefix + "nil track")
	ErrNilTracks   = errors.New(errPrefix + "nil tracks")
	ErrNilPlaylist = errors.New(errPrefix + "nil playlist")
)

type (
	Visibility string
	SearchType string
	Theme      string
	SortBy     string

	// Обычно ответ выглядит так.
	Response[T any] struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`

		// Если не nil, то поле result будет nil.
		Error *Error `json:"error"`

		Result T `json:"result"`
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
		// example: validate.
		Name string `json:"name"`

		// example: Parameters requirements are not met.
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

type UniqueID uint64

func (u UniqueID) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

func (u *UniqueID) FromString(val string) error {
	res, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}
	*u = UniqueID(res)
	return nil
}

type KindID uint32

func (c KindID) String() string {
	return strconv.FormatUint(uint64(c), 10)
}

type RevisionID uint32

func (r RevisionID) String() string {
	return strconv.FormatUint(uint64(r), 10)
}
