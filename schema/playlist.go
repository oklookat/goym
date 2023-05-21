package schema

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type (
	Playlist struct {
		// Владелец плейлиста.
		Owner Owner `json:"owner"`

		// UID владельца плейлиста.
		UID ID `json:"uid"`

		// UUID.
		PlaylistUuid string `json:"playlistUuid"`

		// Уникальный идентификатор плейлиста.
		//
		// Обычно используется для операций над плейлистом.
		Kind ID `json:"kind"`

		// Название.
		Title string `json:"title"`

		// Описание.
		Description          string `json:"description"`
		DescriptionFormatted string `json:"descriptionFormatted,omitempty"`

		// Что-то типа версии плейлиста.
		// Если плейлист изменился: добавили/удалили треки,
		// то Revision прибавляется на 1.
		//
		// Может быть пуст, если плейлист не создан вами(?).
		Revision ID `json:"revision"`

		Available bool `json:"available"`

		// Совместный плейлист?
		Collective bool `json:"collective"`

		// Обложка.
		Cover Cover `json:"cover"`

		// Дата создания.
		Created time.Time `json:"created"`

		// Дата изменения.
		Modified time.Time `json:"modified"`

		// Общая длина в миллисекундах.
		DurationMs uint64 `json:"durationMs"`

		OgImage string `json:"ogImage"`

		// Количество треков.
		TrackCount uint64 `json:"trackCount"`

		// Количество лайков.
		LikesCount uint32 `json:"likesCount"`

		// Видимость.
		Visibility Visibility `json:"visibility"`

		// Треки.
		//
		// Может быть пустым. Зависит от метода, который вернул эту структуру.
		Tracks []TrackItem `json:"tracks"`
	}

	// Рекомендации для плейлиста
	PlaylistRecommendations struct {
		// Уникальный идентификатор партии треков
		BatchId string `json:"batch_id"`

		// Треки.
		Tracks []Track `json:"tracks"`
	}

	PlaylistId struct {
		// Уникальный идентификатор пользователя владеющим плейлистом.
		UID ID `json:"uid"`

		// Уникальный идентификатор плейлиста.
		Kind ID `json:"kind"`
	}

	// GET /users/{userId}/playlists
	GetUserPlaylistsQueryParams struct {
		// like 1000,1003
		Kinds []string `url:",kinds"`

		Mixed bool `url:"mixed"`

		RichTracks bool `url:"rich-tracks"`
	}

	// POST /users/{userId}/playlists/create
	CreatePlaylistRequestBody struct {
		// Название.
		Title string `url:"title"`

		// Видимость.
		Visibility Visibility `url:"visibility"`

		// Описание.
		Description string `url:"description"`
	}

	// POST /users/{userId}/playlists/{kind}/name
	//
	// POST /users/{userId}/playlists/{kind}/description
	ValuePlaylistRequestBody struct {
		// Новое значение.
		Value string `url:"value"`
	}

	// POST /users/{userId}/playlists/{kind}/visibility
	ChangePlaylistVisibilityRequestBody struct {
		Value Visibility `url:"value"`
	}

	// POST /users/{userId}/likes/playlists/add
	KindOwnerUidRequestBody struct {
		// Kind плейлиста.
		Kind ID `url:"kind"`

		// UID владельца плейлиста.
		OwnerUid ID `url:"owner-uid"`
	}

	ResponseLikedPlaylist struct {
		Playlist  Playlist  `json:"playlist"`
		Timestamp time.Time `json:"timestamp"`
	}

	PlaylistAddTracksOperation struct {
		// insert.
		Op     string `json:"op"`
		At     uint64 `json:"at"`
		Tracks []struct {
			ID      ID `json:"id"`
			AlbumID ID `json:"albumId"`
		} `json:"tracks"`
	}

	PlaylistDeleteTracksOperation struct {
		// delete.
		Op     string `json:"op"`
		From   uint64 `json:"from"`
		To     uint64 `json:"to"`
		Tracks []struct {
			ID      ID `json:"id"`
			AlbumID ID `json:"albumId"`
		} `json:"tracks"`
	}
)

func NewAddPlaylistTracksRequestBody(pl Playlist) AddPlaylistTracksRequestBody {
	return AddPlaylistTracksRequestBody{
		Revision: pl.Revision,
		pl:       pl,
	}
}

// POST /users/{userId}/playlists/{kind}/change-relative
//
// POST /users/{userId}/playlists/{kind}/change
type AddPlaylistTracksRequestBody struct {
	Revision ID     `url:"revision"`
	Diff     string `url:"diff"`

	// Не входит в API.
	hDiff []PlaylistAddTracksOperation `url:"-"`
	pl    Playlist                     `url:"-"`
}

func (a *AddPlaylistTracksRequestBody) ParamsToValues() (url.Values, error) {
	diffBytes, err := json.Marshal(a.hDiff)
	if err != nil {
		return nil, err
	}
	a.Diff = string(diffBytes)
	return ParamsToValues(*a)
}

func (a *AddPlaylistTracksRequestBody) AddTracks(item []Track) {
	for _, track := range item {
		if len(track.Albums) == 0 {
			continue
		}
		if len(a.hDiff) > 0 {
			a.hDiff[0].Tracks = append(a.hDiff[0].Tracks, struct {
				ID      ID "json:\"id\""
				AlbumID ID "json:\"albumId\""
			}{
				ID:      track.ID,
				AlbumID: track.Albums[0].ID,
			})
			continue
		}
		a.hDiff = append(a.hDiff, PlaylistAddTracksOperation{
			Op: "insert",
			At: a.pl.TrackCount,
			Tracks: []struct {
				ID      ID "json:\"id\""
				AlbumID ID "json:\"albumId\""
			}{
				{
					ID:      track.ID,
					AlbumID: track.Albums[0].ID,
				},
			},
		})
	}
}

func NewDeletePlaylistTracksRequestBody(pl Playlist) DeletePlaylistTracksRequestBody {
	return DeletePlaylistTracksRequestBody{
		Kind:     pl.Kind,
		Revision: pl.Revision,
		pl:       pl,
	}
}

type DeletePlaylistTracksRequestBody struct {
	Kind     ID     `url:"kind"`
	Revision ID     `url:"revision"`
	Diff     string `url:"diff"`

	// Не входит в API.
	hDiff []PlaylistDeleteTracksOperation `url:"-"`
	pl    Playlist                        `url:"-"`
}

func (d *DeletePlaylistTracksRequestBody) ParamsToValues() (url.Values, error) {
	diffBytes, err := json.Marshal(d.hDiff)
	if err != nil {
		return nil, err
	}
	d.Diff = string(diffBytes)
	return ParamsToValues(*d)
}

func (d *DeletePlaylistTracksRequestBody) AddTrack(item TrackItem) {
	d.hDiff = append(d.hDiff, PlaylistDeleteTracksOperation{
		Op:   "delete",
		From: item.OriginalIndex,
		To:   item.OriginalIndex + 1,
		Tracks: []struct {
			ID      ID "json:\"id\""
			AlbumID ID "json:\"albumId\""
		}{
			{
				ID:      item.ID,
				AlbumID: item.Track.Albums[0].ID,
			},
		},
	})
}

// POST /playlists/list
//
// POST /users/{userId}/likes/playlists/add-multiple
type PlaylistsIdsRequestBody struct {
	// uid владельца плейлиста и kind плейлиста через двоеточие и запятую.
	//
	// "123:9482", "999:8888".
	PlaylistIds []string `url:",playlistIds"`
}

// Добавить в PlaylistIds.
//
// owner - uid владелеца плейлиста
//
// kind - kind плейлиста
func (g *PlaylistsIdsRequestBody) Add(kind ID, uid ID) {
	if len(g.PlaylistIds) == 0 {
		g.PlaylistIds = []string{}
	}
	dat := string(uid) + ":" + string(kind)
	g.PlaylistIds = append(g.PlaylistIds, dat)
}

// Добавить в PlaylistIds.
//
// map[kind плейлиста]uid_владельца
func (g *PlaylistsIdsRequestBody) AddMany(kindUid map[ID]ID) {
	if len(kindUid) == 0 {
		return
	}
	if len(g.PlaylistIds) == 0 {
		g.PlaylistIds = []string{}
	}
	for k, v := range kindUid {
		g.PlaylistIds = append(g.PlaylistIds, fmt.Sprintf("%s:%s", v, k))
	}
}
