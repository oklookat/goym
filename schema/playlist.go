package schema

import "fmt"

type Playlist struct {
	// Владелец плейлиста.
	Owner *Owner `json:"owner"`

	// UUID.
	PlaylistUuid string `json:"playlistUuid"`

	// ???
	Uid int64 `json:"uid"`

	// Обычно используется для операций над плейлистом.
	Kind int64 `json:"kind"`

	// Меняется плейлист - меняется revision (+1).
	Revision int `json:"revision"`

	// Описание.
	Description string `json:"description"`
	Available   bool   `json:"available"`

	// Название.
	Title      string `json:"title"`
	Collective bool   `json:"collective"`

	// Обложка.
	Cover *Cover `json:"cover"`

	Created    string `json:"created"`
	Modified   string `json:"modified"`
	DurationMs int64  `json:"durationMs"`
	OgImage    string `json:"ogImage"`

	// Количество треков.
	TrackCount int `json:"trackCount"`

	// Количество лайков.
	LikesCount int `json:"likesCount"`

	// Видимость.
	//
	// public | private
	Visibility string `json:"visibility"`

	// Треки.
	Tracks []*TrackItem `json:"tracks"`
}

// Рекомендации для плейлиста
type PlaylistRecommendations struct {
	// Уникальный идентификатор партии треков
	BatchId string `json:"batch_id"`

	Tracks []*Track `json:"tracks"`
}

type PlaylistId struct {
	// Уникальный идентификатор пользователя владеющим плейлистом.
	Uid int64 `json:"uid"`

	// Уникальный идентификатор плейлиста.
	Kind int64 `json:"kind"`
}

// GET /users/{userId}/playlists
type GetUserPlaylistsQueryParams struct {
	// like 1000,1003
	Kinds []string `url:",kinds"`

	Mixed bool `url:"mixed"`

	RichTracks bool `url:"rich-tracks"`
}

// POST /users/{userId}/playlists/create
type CreatePlaylistRequestBody struct {
	// Название плейлиста.
	Title string `url:"title"`

	// Видимость плейлиста.
	Visibility Visibility `url:"visibility"`
}

// POST /users/{userId}/playlists/{kind}/name
type RenamePlaylistRequestBody struct {
	// Новое имя плейлиста.
	Value string `url:"value"`
}

// POST /users/{userId}/playlists/{kind}/change-relative
// TODO
type AddTracksToPlaylistRequestBody struct {
	// Используй '{"diff":{"op":"insert","at":0,"tracks":[{"id":"20599729","albumId":"2347459"}]}}' - для добавления,
	// {"diff":{"op":"delete","from":0,"to":1,"tracks":[{"id":"20599729","albumId":"2347459"}]}} - для удаления треков
	Diff     string `url:"diff"`
	Revision string `url:"revision"`
}

// POST /users/{userId}/playlists/{kind}/visibility
type ChangePlaylistVisibilityRequestBody struct {
	Value Visibility `url:"value"`
}

// POST /playlists/list
//
// Доступен метод Add()
type GetPlaylistsByIdsRequestBody struct {
	// uid владельца плейлиста и kind плейлиста через двоеточие и запятую
	PlaylistIds []string `url:",playlistIds"`
}

// Добавить в PlaylistIds.
//
// owner - владелец плейлиста
//
// kind - kind плейлиста
func (g *GetPlaylistsByIdsRequestBody) Add(owner int64, kind int64) {
	if g.PlaylistIds == nil {
		g.PlaylistIds = make([]string, 0)
	}
	g.PlaylistIds = append(g.PlaylistIds, fmt.Sprintf("%d:%d", owner, kind))
}
