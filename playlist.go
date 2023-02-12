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
	endpoint := genApiPath([]string{"users", c.userId, "playlists", "list"})
	data := &schema.Response[[]*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить плейлист по kind.
//
// Доступно только для плейлистов в библиотеке пользователя.
//
// Доступно поле Tracks.
func (c Client) GetMyPlaylistByKind(ctx context.Context, playlistKind schema.KindID) (*schema.Playlist, error) {
	// GET /users/{userId}/playlists/{kind}
	endpoint := genApiPath([]string{"users", c.userId, "playlists", playlistKind.String()})
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Создать плейлист.
func (c Client) CreatePlaylist(ctx context.Context, name string, vis schema.Visibility) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/create
	body := schema.CreatePlaylistRequestBody{
		Title:      name,
		Visibility: vis,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath([]string{"users", c.userId, "playlists", "create"})
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Переименовать плейлист.
func (c Client) RenamePlaylist(ctx context.Context, pl *schema.Playlist, newName string) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/name
	if pl == nil {
		return nil, nil
	}
	body := schema.RenamePlaylistRequestBody{
		Value: newName,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath([]string{"users", c.userId, "playlists", pl.Kind.String(), "name"})
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Удалить плейлист.
func (c Client) DeletePlaylist(ctx context.Context, pl *schema.Playlist) error {
	// POST /users/{userId}/playlists/{kind}/delete
	if pl == nil {
		return nil
	}
	endpoint := genApiPath([]string{"users", c.userId, "playlists", pl.Kind.String(), "delete"})
	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetError(data).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
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
		return nil, nil
	}
	endpoint := genApiPath([]string{"users", c.userId, "playlists", pl.Kind.String(), "recommendations"})
	data := &schema.Response[*schema.PlaylistRecommendations]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Изменить видимость плейлиста.
func (c Client) ChangePlaylistVisibility(ctx context.Context, pl *schema.Playlist, vis schema.Visibility) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/visibility
	if pl == nil {
		return nil, nil
	}
	body := schema.ChangePlaylistVisibilityRequestBody{
		Value: vis,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath([]string{"users", c.userId, "playlists", pl.Kind.String(), "visibility"})
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
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
	if pl == nil || len(tracks) == 0 {
		return nil, nil
	}

	body := schema.AddDeleteTracksToPlaylistRequestBody{}
	if err := body.Add(pl, tracks); err != nil {
		return nil, err
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath([]string{"users", c.userId, "playlists", pl.Kind.String(), "change"})
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Удалить трек из плейлиста.
//
// Возвращает плейлист без поля Tracks.
//
// track - TrackItem из плейлиста (pl).
func (c Client) DeleteTrackFromPlaylist(ctx context.Context, pl *schema.Playlist, track *schema.TrackItem) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/change-relative
	//
	// POST /users/{userId}/playlists/{kind}/change
	if pl == nil {
		return nil, nil
	}
	if track == nil {
		return nil, nil
	}

	body := schema.AddDeleteTracksToPlaylistRequestBody{}
	if err := body.Delete(pl, track); err != nil {
		return nil, err
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath([]string{"users", c.userId, "playlists", pl.Kind.String(), "change"})
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Поставить лайк плейлисту.
func (c Client) LikePlaylist(ctx context.Context, pl *schema.Playlist) error {
	// POST /users/{userId}/likes/playlists/add
	if pl == nil {
		return nil
	}
	body := schema.LikePlaylistRequestBody{
		Kind:     pl.Kind,
		OwnerUid: pl.UID,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	endpoint := genApiPath([]string{"users", c.userId, "likes", "playlists", "add"})
	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return err
}

// Убрать лайк с плейлиста.
func (c Client) UnlikePlaylist(ctx context.Context, pl *schema.Playlist) error {
	// POST /users/{userId}/likes/playlists/{kind}/remove
	if pl == nil {
		return nil
	}

	uidAndKind := pl.UID.String() + "-" + pl.Kind.String()
	endpoint := genApiPath([]string{"users", c.userId, "likes", "playlists", uidAndKind, "remove"})
	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// Получить плейлисты.
//
// kindUid - map[kind плейлиста]uid_владельца
func (c Client) GetPlaylistsByKindUid(ctx context.Context, kindUid map[schema.KindID]schema.UniqueID) ([]*schema.Playlist, error) {
	if len(kindUid) == 0 {
		return nil, nil
	}

	// GET /playlists/list
	body := schema.PlaylistsIdsRequestBody{}
	body.AddMany(kindUid)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath([]string{"playlists", "list"})
	data := &schema.Response[[]*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
