package schema

import "time"

type Album struct {
	// Идентификатор альбома.
	ID int64 `json:"id"`

	// Название альбома.
	Title string `json:"title"`

	// Мета тип (single, podcast, music, remix).
	MetaType string `json:"metaType"`

	// Год релиза.
	Year int `json:"year"`

	// Дата релиза в формате ISO 8601.
	ReleaseDate time.Time `json:"releaseDate"`

	// Ссылка на обложку.
	CoverURI string `json:"coverUri"`

	// Ссылка на превью Open Graph.
	OgImage string `json:"ogImage"`

	// Жанр музыки.
	Genre string `json:"genre"`

	// Количество треков.
	TrackCount int `json:"trackCount"`

	// Количество лайков.
	LikesCount int `json:"likesCount"`

	// Является ли альбом новым.
	Recent bool `json:"recent"`

	// Популярен ли альбом у слушателей.
	VeryImportant bool `json:"veryImportant"`

	// Артисты.
	Artists []*Artist `json:"artists"`

	// Доступен ли альбом.
	Available bool `json:"available"`

	// Доступен ли альбом для пользователей с подпиской.
	AvailableForPremiumUsers bool `json:"availableForPremiumUsers"`

	AvailableForOptions []string `json:"availableForOptions"`

	// Доступен ли альбом из приложения для телефона.
	AvailableForMobile bool `json:"availableForMobile"`

	// Доступен ли альбом частично для пользователей без подписки.
	AvailablePartially bool `json:"availablePartially"`

	// ID лучших треков альбома.
	Bests []int64 `json:"bests"`

	// Лейблы.
	//
	// Может быть как слайсом строк с названиями, так и слайсом структур Label.
	//
	// (?) Слайсы строк используются при поиске, а слайсы структур в остальных случаях.
	Labels []any `json:"labels"`

	StorageDir string `json:"storageDir"`

	TrackPosition struct {
		Volume int `json:"volume"`
		Index  int `json:"index"`
	} `json:"trackPosition"`

	Regions          []string      `json:"regions"`
	AvailableRegions []interface{} `json:"availableRegions"`

	// например: "single".
	Type *string `json:"type"`

	// например: "Remix".
	Version *string `json:"version"`

	// Ремиксы, и прочее. Не nil, например когда запрашивается альбом с треками.
	Duplicates []*Album `json:"duplicates"`

	// Треки альбома, разделенные по дискам.
	Volumes [][]*Track `json:"volumes"`
}

type AlbumShort struct {
	ID        int64     `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

// POST /albums
type GetAlbumsByIdsRequestBody struct {
	// ID альбомов.
	AlbumIds []int64 `url:",album-ids"`
}

// POST /users/{userId}/likes/albums/add
type LikeAlbumRequestBody struct {
	// ID альбома.
	AlbumId int64 `url:"album-id"`
}
