package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить плейлисты текущего пользователя.
//
// Без поля Tracks.
func (c Client) MyPlaylists(ctx context.Context) ([]*schema.Playlist, error) {
	// GET /users/{userId}/playlists/list
	endpoint := genApiPath("users", c.userId, "playlists", "list")
	data := &schema.Response[[]*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить лайкнутые плейлисты.
//
// Без поля Tracks.
func (c Client) LikedPlaylists(ctx context.Context) ([]*schema.ResponseLikedPlaylist, error) {
	// GET /users/{userId}/likes/playlists
	// todo тут еще pager есть
	return likesDislikes[[]*schema.ResponseLikedPlaylist](ctx, &c, true, "playlists")
}

// Получить плейлист по kind.
//
// Доступно только для плейлистов в библиотеке пользователя.
//
// Доступно поле Tracks.
func (c Client) MyPlaylist(ctx context.Context, kind schema.ID) (*schema.Playlist, error) {
	// GET /users/{userId}/playlists/{kind}
	endpoint := genApiPath("users", c.userId, "playlists", kind.String())
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить плейлисты.
//
// kindUid - map[kind плейлиста]uid_владельца
//
// Доступно поле Tracks.
func (c Client) PlaylistsByKindUid(ctx context.Context, kindUid map[schema.ID]schema.ID) ([]*schema.Playlist, error) {
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

	endpoint := genApiPath("playlists", "list")
	data := &schema.Response[[]*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
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

	endpoint := genApiPath("users", c.userId, "playlists", "create")
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Переименовать плейлист.
func (c Client) RenamePlaylist(ctx context.Context, kind schema.ID, newName string) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/name
	body := schema.RenamePlaylistRequestBody{
		Value: newName,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath("users", c.userId, "playlists", kind.String(), "name")
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Удалить плейлист.
func (c Client) DeletePlaylist(ctx context.Context, kind schema.ID) error {
	// POST /users/{userId}/playlists/{kind}/delete
	endpoint := genApiPath("users", c.userId, "playlists", kind.String(), "delete")
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
func (c Client) PlaylistRecommendations(ctx context.Context, kind schema.ID) (*schema.PlaylistRecommendations, error) {
	// GET /users/{userId}/playlists/{kind}/recommendations
	endpoint := genApiPath("users", c.userId, "playlists", kind.String(), "recommendations")
	data := &schema.Response[*schema.PlaylistRecommendations]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Изменить видимость плейлиста.
func (c Client) SetPlaylistVisibility(ctx context.Context, kind schema.ID, vis schema.Visibility) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/visibility
	body := schema.ChangePlaylistVisibilityRequestBody{
		Value: vis,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath("users", c.userId, "playlists", kind.String(), "visibility")
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
func (c Client) AddToPlaylist(ctx context.Context, pl *schema.Playlist, tracks []*schema.Track) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/change-relative
	// ||
	// POST /users/{userId}/playlists/{kind}/change
	body := schema.AddDeleteTracksToPlaylistRequestBody{}
	if err := body.Add(pl, tracks); err != nil {
		return nil, err
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath("users", c.userId, "playlists", pl.Kind.String(), "change")
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
func (c Client) DeleteFromPlaylist(ctx context.Context, pl *schema.Playlist, track *schema.TrackItem) (*schema.Playlist, error) {
	// POST /users/{userId}/playlists/{kind}/change-relative
	//
	// POST /users/{userId}/playlists/{kind}/change
	body := schema.AddDeleteTracksToPlaylistRequestBody{}
	if err := body.Delete(pl, track); err != nil {
		return nil, err
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath("users", c.userId, "playlists", pl.Kind.String(), "change")
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Поставить лайк плейлисту.
//
// kind и uid можно получить из плейлиста.
func (c Client) LikePlaylist(ctx context.Context, kind schema.ID, ownerUid schema.ID) error {
	// POST /users/{userId}/likes/playlists/add
	body := schema.KindOwnerUidRequestBody{
		Kind:     kind,
		OwnerUid: ownerUid,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}
	return addRemove(ctx, &c, vals, true, "playlists")
}

// Убрать лайк с плейлиста.
//
// kind и uid можно получить из плейлиста.
func (c Client) UnlikePlaylist(ctx context.Context, kind schema.ID, uid schema.ID) error {
	// POST /users/{userId}/likes/playlists/{kind}/remove
	body := schema.PlaylistsIdsRequestBody{}
	body.Add(kind, uid)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}
	return addRemove(ctx, &c, vals, false, "playlists")
}

// Поставить лайк плейлистам.
//
// kind и uid можно получить из плейлиста.
func (c Client) LikePlaylists(ctx context.Context, kindUid map[schema.ID]schema.ID) error {
	return c.likeUnlikePlaylists(ctx, kindUid, true)
}

// Снять лайк с плейлистов.
//
// kind и uid можно получить из плейлиста.
func (c Client) UnlikePlaylists(ctx context.Context, kindUid map[schema.ID]schema.ID) error {
	return c.likeUnlikePlaylists(ctx, kindUid, false)
}

func (c Client) likeUnlikePlaylists(ctx context.Context, kindUid map[schema.ID]schema.ID, like bool) error {
	// POST /users/{userId}/likes/playlists/add-multiple
	// // POST /users/{userId}/likes/playlists/remove
	body := schema.PlaylistsIdsRequestBody{}
	body.AddMany(kindUid)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}
	return addRemoveMultiple(ctx, &c, vals, like, "playlists")
}
