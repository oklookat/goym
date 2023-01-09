package goym

import (
	"encoding/json"
	"errors"
)

// {"album-ids": "id,id,id,id,id"}.
func formAlbumIds(albumIds []int64) map[string]string {
	return map[string]string{
		"album-ids": i64Join(albumIds),
	}
}

// {"album-id": "id"}.
func formAlbumId(albumId int64) map[string]string {
	return map[string]string{
		"album-id": i2s(albumId),
	}
}

// {"artist-id": "id"}.
func formArtistId(artistId int64) map[string]string {
	return map[string]string{
		"artist-id": i2s(artistId),
	}
}

// Параметры для поиска.
func searchQuery(text string, page uint32, _type string, nocorrect bool) map[string]string {
	return map[string]string{
		"text":      text,
		"page":      i2s(page),
		"type":      _type,
		"nocorrect": b2s(nocorrect),
	}
}

// Параметры для поисковых подсказок.
func searchSuggestQuery(part string) map[string]string {
	return map[string]string{
		"part": part,
	}
}

// {"track-ids": "id,id,id,id,id"}.
func formTrackIds(trackIds []int64) map[string]string {
	return map[string]string{
		"track-ids": i64Join(trackIds),
	}
}

// {"track-id": "id"}.
func formTrackId(trackId int64) map[string]string {
	return map[string]string{
		"track-id": i2s(trackId),
	}
}

// {"title": "hello", "visibility": "public"}
func formTitleVisibility(title string, public bool) map[string]string {
	return map[string]string{
		"title":      title,
		"visibility": visibilityToString(public),
	}
}

// {"value": "hello"}
func formValue(val string) map[string]string {
	return map[string]string{
		"value": val,
	}
}

type playlistChange struct {
	// ID плейлиста.
	Kind int64 `json:"kind"`

	// Ревизия API (?). Можно получить из Playlist.Revision.
	Revision int `json:"revision"`

	// Операции над плейлистом.
	Diff []*playlistDiff `json:"diff"`
}

func (p *playlistChange) New(pl *Playlist) {
	if pl == nil {
		return
	}
	p.Kind = pl.Kind
	p.Revision = pl.Revision
}

func (p *playlistChange) AddDiff(d *playlistDiff) {
	if d == nil {
		return
	}
	if len(p.Diff) == 0 {
		p.Diff = make([]*playlistDiff, 0)
	}
	p.Diff = append(p.Diff, d)
}

func (p *playlistChange) GetForm() (map[string]string, error) {
	diffByte, err := json.Marshal(p.Diff)
	if err != nil {
		return nil, err
	}
	var mapped = map[string]string{
		"kind":     i2s(p.Kind),
		"revision": i2s(p.Revision),
		"diff":     string(diffByte),
	}
	return mapped, err
}

// /users/{userId}/playlists/{kind}/change
//
// [{"op":"insert","at":4,"tracks":[{"id":"40315071","albumId":"5236619"}]}]
type playlistDiff struct {
	// insert | delete
	Op string `json:"op"`

	// При добавлении трека. Позиция. Например, если в плейлисте 4 трека, то при insert = 4,
	// трек будет помещен в конец плейлиста
	At *int `json:"at,omitempty"`

	// При удалении. Позиция в плейлисте, с которой начать удалять треки.
	From *int `json:"from,omitempty"`

	// При удалении. Позиция в плейлисте, где закончить удалять треки.
	To *int `json:"to,omitempty"`

	// Треки, которые надо добавить/удалить из плейлиста.
	Tracks []struct {
		ID      string `json:"id"`
		AlbumID string `json:"albumId"`
	} `json:"tracks"`
}

// Новая операция над плейлистом.
func (p *playlistDiff) NewInsert(pl *Playlist, tracks []*Track) error {
	if pl == nil {
		return errors.New("nil playlist")
	}
	if len(tracks) == 0 {
		return errors.New("empty tracks")
	}
	if len(p.Tracks) == 0 {
		p.Tracks = make([]struct {
			ID      string "json:\"id\""
			AlbumID string "json:\"albumId\""
		}, 0)
	}

	p.Op = "insert"
	p.At = &pl.TrackCount

	for i := range tracks {
		if len(tracks[i].Albums) == 0 {
			return errors.New("track without album")
		}
		var id = i2s(tracks[i].ID)
		var albumId = i2s(tracks[i].Albums[0].ID)
		var trc = struct {
			ID      string "json:\"id\""
			AlbumID string "json:\"albumId\""
		}{
			ID:      id,
			AlbumID: albumId,
		}
		p.Tracks = append(p.Tracks, trc)
	}
	return nil
}

func (p *playlistDiff) NewDelete(pl *Playlist, from int, to int) error {
	if len(pl.Tracks) == 0 {
		return errors.New("no tracks in playlist")
	}

	p.Op = "delete"
	p.From = &from
	p.To = &to

	// ищем треки
	for _, ti := range pl.Tracks {
		var origIndex = ti.OriginalIndex
		if !(origIndex >= from && origIndex <= to) {
			continue
		}
		var trc = struct {
			ID      string "json:\"id\""
			AlbumID string "json:\"albumId\""
		}{
			ID:      i2s(ti.Track.ID),
			AlbumID: i2s(ti.Track.Albums[0].ID),
		}
		p.Tracks = append(p.Tracks, trc)
	}

	return nil
}
