package schema

import (
	"fmt"
	"strconv"
	"strings"
)

type Playlist struct {
	// Владелец плейлиста.
	Owner *Owner `json:"owner"`

	// ?
	PlaylistUuid string `json:"playlistUuid"`

	// ?
	UID int64 `json:"uid"`

	// Обычно используется для операций над плейлистом.
	Kind int64 `json:"kind"`

	// Что-то типа версии плейлиста.
	// Если плейлист изменился: добавили, удалили треки,
	// то Revision прибавляется на 1.
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
	//
	// Может быть nil. Зависит от метода, которым получен плейлист.
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
	UID int64 `json:"uid"`

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

// Доступны методы Add() и Delete()
//
// POST /users/{userId}/playlists/{kind}/change-relative
//
// POST /users/{userId}/playlists/{kind}/change
type AddDeleteTracksToPlaylistRequestBody struct {
	// Playlist difference. Операция над плейлистом.
	Diff string `url:"diff"`

	// см. Playlist.Revision.
	Revision string `url:"revision"`
}

// Добавить треки в плейлист.
func (a *AddDeleteTracksToPlaylistRequestBody) Add(pl *Playlist, tracks []*Track) error {
	if len(tracks) == 0 {
		return ErrNilTracks
	}
	if err := a.fillBase(pl); err != nil {
		return err
	}

	var trackObjs = []string{}
	for _, t := range tracks {
		if len(t.Albums) == 0 {
			return fmt.Errorf(errPrefix+"track (id %d) without albums", t.ID)
		}
		trackObjs = append(trackObjs, a.getTrackObj(t.ID, t.Albums[0].ID))
	}

	var at = strconv.Itoa(pl.TrackCount)      // добавить треки в конец плейлиста
	tracksObj := strings.Join(trackObjs, ",") // trackobj,trackobj,trackobj

	// {"diff":{"op":"insert","at":144,"tracks":[{"id":"20599729","albumId":"2347459"}]}}
	a.Diff = fmt.Sprintf(`{"diff":{"op":"insert","at":%s,"tracks":[%s]}}`, at, tracksObj)
	return nil
}

// Удалить трек из плейлиста.
//
// track - TrackItem из Playlist.Tracks
func (a *AddDeleteTracksToPlaylistRequestBody) Delete(pl *Playlist, track *TrackItem) error {
	if track == nil {
		return ErrNilTrack
	}
	if err := a.fillBase(pl); err != nil {
		return err
	}
	var trackObj = ""
	var from = 0
	var to = 0

	for _, t := range pl.Tracks {
		if t.Track == nil {
			return ErrNilTrack
		}
		if len(t.Track.Albums) == 0 {
			return fmt.Errorf(errPrefix+"track (id %d) without albums", t.ID)
		}
		if track.ID != t.ID {
			continue
		}
		from = track.OriginalIndex
		to = from + 1
		trackObj = a.getTrackObj(t.Track.ID, t.Track.Albums[0].ID)
	}

	// {"diff":{"op":"delete","from":0,"to":1,"tracks":[{"id":"20599729","albumId":"2347459"}]}}
	a.Diff = fmt.Sprintf(`{"diff":{"op":"delete","from":%d,"to":%d,"tracks":[%s]}}`, from, to, trackObj)
	return nil
}

func (a *AddDeleteTracksToPlaylistRequestBody) fillBase(pl *Playlist) error {
	if pl == nil {
		return ErrNilPlaylist
	}
	a.Revision = strconv.Itoa(pl.Revision)
	return nil
}

// {"id":"1234","albumId":"1234"}
func (a AddDeleteTracksToPlaylistRequestBody) getTrackObj(id int64, albumId int64) string {
	var idStr = strconv.FormatInt(id, 10)
	var obj = `{"id":`        // {"id":
	obj += `"` + idStr + `",` // {"id":"1234",
	var albumIdStr = strconv.FormatInt(albumId, 10)
	obj += `"albumId":"` + albumIdStr + `"}` // {"id":"1234","albumId":"1234"}
	return obj
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
