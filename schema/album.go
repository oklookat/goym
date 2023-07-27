package schema

import "time"

type (
	// Альбом.
	Album struct {
		// Идентификатор альбома.
		ID ID `json:"id"`

		// Название альбома.
		Title string `json:"title"`

		// Мета тип (single, podcast, music, remix).
		MetaType AlbumMetaType `json:"metaType"`

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

		// Является новинкой.
		Recent bool `json:"recent"`

		// Является важным.
		VeryImportant bool `json:"veryImportant"`

		// Исполнители альбома, в минимальной информации.
		Artists []Artist `json:"artists"`

		// Лейблы.
		//
		// Может быть как слайсом строк с названиями, так и слайсом структур Label.
		//
		// (?) Слайсы строк используются при поиске, а слайсы структур в остальных случаях.
		Labels []any `json:"labels"`

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
		Bests []ID `json:"bests"`

		// например: "single".
		Type string `json:"type"`

		// Ремиксы, и прочее. Не пуст, например когда запрашивается альбом с треками.
		Duplicates []Album `json:"duplicates"`

		StorageDir string `json:"storageDir"`

		TrackPosition struct {
			Volume int `json:"volume"`
			Index  int `json:"index"`
		} `json:"trackPosition"`

		Regions          []string      `json:"regions"`
		AvailableRegions []interface{} `json:"availableRegions"`

		// например: "Remix".
		Version string `json:"version"`

		// Треки альбома, разделенные по дискам.
		Volumes [][]Track `json:"volumes"`
	}

	AlbumShort struct {
		ID        ID        `json:"id"`
		Timestamp time.Time `json:"timestamp"`
	}

	// POST /users/{userId}/likes/albums/add
	AlbumIdRequestBody struct {
		// ID альбома.
		AlbumId ID `url:"album-id"`
	}
)

// POST /users/{userId}/likes/albums/add-multiple
//
// POST /users/{userId}/likes/albums/remove
type AlbumIdsRequestBody struct {
	// ID альбомов.
	AlbumIds []ID `url:",albumIds"`
}

func (l *AlbumIdsRequestBody) SetIds(ids ...ID) {
	l.AlbumIds = []ID{}
	l.AlbumIds = append(l.AlbumIds, ids...)
}

// Тип альбома.
type AlbumMetaType string

const (
	AlbumMetaTypeSingle  AlbumMetaType = "single"
	AlbumMetaTypePodcast AlbumMetaType = "podcast"
	AlbumMetaTypeMusic   AlbumMetaType = "music"
	AlbumMetaTypeRemix   AlbumMetaType = "remix"
)
