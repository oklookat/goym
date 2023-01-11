package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить плейлисты текущего пользователя.
//
// Доступно поле Tracks.
func (c Client) GetMyPlaylists(ctx context.Context) ([]*schema.Playlist, error) {
	// GET /users/{userId}/playlists/list
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", "list"})
	var data = &schema.TypicalResponse[[]*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получить плейлист текущего пользователя по kind.
//
// Доступно поле Tracks.
func (c Client) GetMyPlaylistByKind(ctx context.Context, playlistKind int64) (*schema.Playlist, error) {
	// GET /users/{userId}/playlists/{kind}
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(playlistKind)})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Создать плейлист.
func (c Client) CreatePlaylist(ctx context.Context, name string, vis schema.Visibility) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/create
	var body = schema.CreatePlaylistRequestBody{
		Title:      name,
		Visibility: vis,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "playlists", "create"})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Переименовать плейлист.
func (c Client) RenamePlaylist(ctx context.Context, pl *schema.Playlist, newName string) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/name
	if pl == nil {
		return nil, ErrNilPlaylist
	}
	var body = schema.RenamePlaylistRequestBody{
		Value: newName,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "name"})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Удалить плейлист.
func (c Client) DeletePlaylist(ctx context.Context, pl *schema.Playlist) error {
	// POST /users/{userId}/playlists/{kind}/delete
	if pl == nil {
		return ErrNilPlaylist
	}
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "delete"})
	var data = &schema.TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return err
}

// Получить рекомендации на основе плейлиста.
//
// Только для плейлистов, созданных пользователем.
//
// Если в плейлисте нет треков, рекомендаций не будет.
func (c Client) GetPlaylistRecommendations(ctx context.Context, pl *schema.Playlist) (*schema.PlaylistRecommendations, error) {
	// GET /users/{userId}/playlists/{kind}/recommendations
	if pl == nil {
		return nil, ErrNilPlaylist
	}
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "recommendations"})
	var data = &schema.TypicalResponse[*schema.PlaylistRecommendations]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Изменить видимость плейлиста.
func (c Client) ChangePlaylistVisibility(ctx context.Context, pl *schema.Playlist, vis schema.Visibility) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/visibility
	if pl == nil {
		return nil, ErrNilPlaylist
	}
	var body = schema.ChangePlaylistVisibilityRequestBody{
		Value: vis,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "visibility"})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Добавить треки в плейлист.
//
// Возвращает плейлист без поля Tracks.
func (c Client) AddTracksToPlaylist(ctx context.Context, pl *schema.Playlist, tracks []*schema.Track) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/change-relative
	// ||
	// POST /users/{userId}/playlists/{kind}/change
	if pl == nil {
		return nil, ErrNilPlaylist
	}
	if len(tracks) == 0 {
		return nil, ErrNilTracks
	}

	var body = schema.AddDeleteTracksToPlaylistRequestBody{}
	if err := body.Add(pl, tracks); err != nil {
		return nil, err
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "change"})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Удалить трек из плейлиста.
//
// Возвращает плейлист без поля Tracks.
//
// track - TrackItem из плейлиста.
func (c Client) DeleteTrackFromPlaylist(ctx context.Context, pl *schema.Playlist, track *schema.TrackItem) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/change-relative
	//
	// POST /users/{userId}/playlists/{kind}/change
	if pl == nil {
		return nil, ErrNilPlaylist
	}
	if track == nil {
		return nil, ErrNilTrack
	}

	var body = schema.AddDeleteTracksToPlaylistRequestBody{}
	if err := body.Delete(pl, track); err != nil {
		return nil, err
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "change"})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}
