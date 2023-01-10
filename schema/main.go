package schema

type Visibility string
type SearchType string

const (
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
