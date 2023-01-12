package schema

import (
	"errors"
)

type Visibility string
type SearchType string
type Theme string

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
)

var (
	ErrNilTrack    = errors.New(errPrefix + "nil track")
	ErrNilTracks   = errors.New(errPrefix + "nil tracks")
	ErrNilPlaylist = errors.New(errPrefix + "nil playlist")
)

// Обычно ответ выглядит так.
type TypicalResponse[T any] struct {
	InvocationInfo InvocationInfo `json:"invocationInfo"`

	// Если не nil, то поле result будет nil.
	Error *Error `json:"error"`

	Result T `json:"result"`
}

// Что-то техническое.
type InvocationInfo struct {
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
type Error struct {
	// example: validate.
	Name string `json:"name"`

	// example: Parameters requirements are not met.
	Message string `json:"message"`
}
