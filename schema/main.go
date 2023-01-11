package schema

import (
	"errors"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Visibility string
type SearchType string

const (
	errPrefix = "goym/schema: "

	ApiUrl = "https://api.music.yandex.net"

	VisibilityPrivate  Visibility = "private"
	VisibilityPublic   Visibility = "public"
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
	InvocationInfo *InvocationInfo `json:"invocationInfo"`

	// Если не nil, то поле result будет nil.
	Error *Error `json:"error"`

	Result T `json:"result"`
}

// Что-то техническое.
type InvocationInfo struct {
	// (?) Время выполнения запроса в миллисекундах.
	//
	// string | int
	ExecDurationMillis any `json:"exec-duration-millis"`

	// Адрес какого-то сервера Яндекс.Музыки.
	Hostname string `json:"hostname"`

	// ID запроса.
	ReqID string `json:"req-id"`
}

// Ошибка. Ошибка валидации, например.
type Error struct {
	// example: validate.
	Name string `json:"name"`

	// example: Parameters requirements are not met.
	Message string `json:"message"`
}

// Преобразовать struct (НЕ указатель на struct) в url.Values.
//
// Доступно для структур, название которых заканчивается на "Params" и "Body".
//
// Но не всегда. В некоторых структурах есть дополнительные методы. Читайте доки (c).
//
// После получения Values можно сделать Encode(), и отправить GET или POST (request body).
func ParamsToValues(s any) (url.Values, error) {
	return query.Values(s)
}
