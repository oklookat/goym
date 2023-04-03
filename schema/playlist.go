package schema

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	Playlist struct {
		// Владелец плейлиста.
		Owner *Owner `json:"owner"`

		// UID владельца плейлиста.
		UID ID `json:"uid"`

		// UUID.
		PlaylistUuid *string `json:"playlistUuid"`

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
		// Может быть nil, если плейлист не создан вами(?).
		Revision *ID `json:"revision"`

		Available *bool `json:"available"`

		Collective bool `json:"collective"`

		// Обложка.
		Cover *Cover `json:"cover"`

		Created    *string `json:"created"`
		Modified   *string `json:"modified"`
		DurationMs *uint64 `json:"durationMs"`
		OgImage    *string `json:"ogImage"`

		// Количество треков.
		TrackCount uint32 `json:"trackCount"`

		// Количество лайков.
		LikesCount *uint32 `json:"likesCount"`

		// Видимость.
		Visibility *Visibility `json:"visibility"`

		// Треки.
		//
		// Может быть nil. Зависит от метода, который вернул эту структуру.
		Tracks []*TrackItem `json:"tracks"`
	}

	// Рекомендации для плейлиста
	PlaylistRecommendations struct {
		// Уникальный идентификатор партии треков
		BatchId string `json:"batch_id"`

		// Треки.
		Tracks []*Track `json:"tracks"`
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
		// Название плейлиста.
		Title string `url:"title"`

		// Видимость плейлиста.
		Visibility Visibility `url:"visibility"`
	}

	// POST /users/{userId}/playlists/{kind}/name
	RenamePlaylistRequestBody struct {
		// Новое имя плейлиста.
		Value string `url:"value"`
	}

	// POST /users/{userId}/playlists/{kind}/visibility
	ChangePlaylistVisibilityRequestBody struct {
		Value Visibility `url:"value"`
	}

	// POST /users/{userId}/likes/playlists/add
	LikePlaylistRequestBody struct {
		// Kind плейлиста.
		Kind ID `url:"kind"`

		// UID владельца плейлиста.
		OwnerUid ID `url:"owner-uid"`
	}

	// POST /users/{userId}/likes/playlists/add
	GetPlaylistByUidKindQueryParams struct {
		LikePlaylistRequestBody
	}
)

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

	trackObjs := []string{}
	for _, t := range tracks {
		if len(t.Albums) == 0 {
			return fmt.Errorf(errPrefix+"track (id %d) without albums", t.ID)
		}
		trackObjs = append(trackObjs, a.getTrackObj(t.ID, t.Albums[0].ID))
	}

	at := strconv.FormatUint(uint64(pl.TrackCount), 10) // добавить треки в конец плейлиста
	tracksObj := strings.Join(trackObjs, ",")           // trackobj,trackobj,trackobj

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
	trackObj := ""
	var from uint16 = 0
	var to uint16 = 0

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
	if pl.Revision == nil {
		return ErrNilPlaylist
	}
	a.Revision = pl.Revision.String()
	return nil
}

// {"id":"1234","albumId":"1234"}
func (a AddDeleteTracksToPlaylistRequestBody) getTrackObj(id ID, albumId ID) string {
	idStr := id.String()
	obj := `{"id":`           // {"id":
	obj += `"` + idStr + `",` // {"id":"1234",
	albumIdStr := albumId.String()
	obj += `"albumId":"` + albumIdStr + `"}` // {"id":"1234","albumId":"1234"}
	return obj
}

// POST /playlists/list
type PlaylistsIdsRequestBody struct {
	// uid владельца плейлиста и kind плейлиста через двоеточие и запятую
	PlaylistIds []string `url:",playlistIds"`
}

// Добавить в PlaylistIds.
//
// owner - владелец плейлиста
//
// kind - kind плейлиста
func (g *PlaylistsIdsRequestBody) Add(kind ID, owner ID) {
	if g.PlaylistIds == nil {
		g.PlaylistIds = make([]string, 0)
	}
	dat := owner.String() + ":" + kind.String()
	g.PlaylistIds = append(g.PlaylistIds, dat)
}

// Добавить в PlaylistIds.
//
// map[kind плейлиста]uid_владельца
func (g *PlaylistsIdsRequestBody) AddMany(kindUid map[ID]ID) {
	if len(kindUid) == 0 {
		return
	}
	if g.PlaylistIds == nil {
		g.PlaylistIds = make([]string, 0)
	}
	for k, v := range kindUid {
		g.PlaylistIds = append(g.PlaylistIds, fmt.Sprintf("%d:%d", v, k))
	}
}
