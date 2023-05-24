package goym

import (
	"context"
	"errors"
	"strings"

	"github.com/oklookat/goym/schema"
)

// Получить плейлисты текущего пользователя.
//
// Без поля Tracks.
func (c Client) MyPlaylists(ctx context.Context) (schema.Response[[]schema.Playlist], error) {
	// GET /users/{userId}/playlists/list
	endpoint := genApiPath("users", string(c.UserId), "playlists", "list")
	data := &schema.Response[[]schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить лайкнутые плейлисты.
//
// Без поля Tracks.
func (c Client) LikedPlaylists(ctx context.Context) (schema.Response[[]schema.ResponseLikedPlaylist], error) {
	// GET /users/{userId}/likes/playlists
	return likesDislikes[[]schema.ResponseLikedPlaylist](ctx, &c, true, "playlists")
}

// Получить плейлист по kind.
//
// Доступно только для плейлистов в библиотеке пользователя.
//
// Доступно поле Tracks.
func (c Client) MyPlaylist(ctx context.Context, kind schema.ID) (schema.Response[*schema.Playlist], error) {
	// GET /users/{userId}/playlists/{kind}
	endpoint := genApiPath("users", string(c.UserId), "playlists", kind.String())
	data := &schema.Response[*schema.Playlist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить плейлисты.
//
// kindUid - map[kind плейлиста]uid_владельца
//
// Доступно поле Tracks.
func (c Client) PlaylistsByKindUid(ctx context.Context, kindUid map[schema.ID]schema.ID) (schema.Response[[]schema.Playlist], error) {
	data := &schema.Response[[]schema.Playlist]{}

	if len(kindUid) == 0 {
		return *data, nil
	}

	// GET /playlists/list
	body := schema.PlaylistsIdsRequestBody{}
	body.AddMany(kindUid)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("playlists", "list")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Создать плейлист.
func (c Client) CreatePlaylist(ctx context.Context, name, description string, vis schema.Visibility) (schema.Response[*schema.Playlist], error) {
	// POST /users/{userId}/playlists/create
	data := &schema.Response[*schema.Playlist]{}

	body := schema.CreatePlaylistRequestBody{
		Title:       name,
		Visibility:  vis,
		Description: description,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("users", string(c.UserId), "playlists", "create")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Переименовать плейлист.
func (c Client) RenamePlaylist(ctx context.Context, kind schema.ID, newName string) (schema.Response[*schema.Playlist], error) {
	// POST /users/{userId}/playlists/{kind}/name
	data := &schema.Response[*schema.Playlist]{}

	body := schema.ValuePlaylistRequestBody{
		Value: newName,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("users", string(c.UserId), "playlists", kind.String(), "name")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Изменить описание плейлиста.
func (c Client) SetPlaylistDescription(ctx context.Context, kind schema.ID, description string) (schema.Response[*schema.Playlist], error) {
	// POST /users/{userId}/playlists/{kind}/description
	data := &schema.Response[*schema.Playlist]{}

	body := schema.ValuePlaylistRequestBody{
		Value: description,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("users", string(c.UserId), "playlists", kind.String(), "description")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Удалить плейлист.
func (c Client) DeletePlaylist(ctx context.Context, kind schema.ID) (schema.Response[string], error) {
	// POST /users/{userId}/playlists/{kind}/delete
	endpoint := genApiPath("users", string(c.UserId), "playlists", kind.String(), "delete")
	data := &schema.Response[string]{}
	resp, err := c.Http.R().SetError(data).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить рекомендации на основе плейлиста.
//
// Только для плейлистов, созданных пользователем.
//
// Если в плейлисте нет треков, рекомендаций не будет.
func (c Client) PlaylistRecommendations(ctx context.Context, kind schema.ID) (schema.Response[*schema.PlaylistRecommendations], error) {
	// GET /users/{userId}/playlists/{kind}/recommendations
	endpoint := genApiPath("users", string(c.UserId), "playlists", kind.String(), "recommendations")
	data := &schema.Response[*schema.PlaylistRecommendations]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Изменить видимость плейлиста.
func (c Client) SetPlaylistVisibility(ctx context.Context, kind schema.ID, vis schema.Visibility) (schema.Response[*schema.Playlist], error) {
	// POST /users/{userId}/playlists/{kind}/visibility
	data := &schema.Response[*schema.Playlist]{}

	body := schema.ChangePlaylistVisibilityRequestBody{
		Value: vis,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("users", string(c.UserId), "playlists", kind.String(), "visibility")
	resp, err := c.Http.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Добавить треки в плейлист.
//
// Возвращает плейлист без поля Tracks.
func (c Client) AddToPlaylist(ctx context.Context, pl schema.Playlist, tracks []schema.Track) (schema.Response[*schema.Playlist], error) {
	// POST /users/{userId}/playlists/{kind}/change-relative
	// ||
	// POST /users/{userId}/playlists/{kind}/change
	data := &schema.Response[*schema.Playlist]{}

	if len(tracks) == 0 {
		return *data, nil
	}

	body := schema.NewAddPlaylistTracksRequestBody(pl)
	body.AddTracks(tracks)
	vals, err := body.ParamsToValues()
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("users", string(c.UserId), "playlists", pl.Kind.String(), "change")
	resp, err := c.Http.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Удалить трек из плейлиста.
//
// Возвращает плейлист без поля Tracks.
//
// track - TrackItem из плейлиста (pl).
func (c Client) DeleteTracksFromPlaylist(ctx context.Context, pl schema.Playlist, tracks []schema.ID) (schema.Response[*schema.Playlist], error) {
	// POST /users/{userId}/playlists/{kind}/change-relative
	//
	// POST /users/{userId}/playlists/{kind}/change
	data := &schema.Response[*schema.Playlist]{}

	if len(tracks) == 0 {
		return *data, nil
	}

	// Try to get playlist tracks if playlist tracks from args empty.
	if len(pl.Tracks) == 0 {
		myPl, err := c.MyPlaylist(ctx, pl.Kind)
		if err != nil {
			return *data, err
		}
		if myPl.Result == nil {
			return *data, errors.New("deleteFromPlaylist: nil result")
		}
		if len(myPl.Result.Tracks) == 0 {
			return *data, nil
		}
		pl = *myPl.Result
	}

	body := schema.NewDeletePlaylistTracksRequestBody(pl)
	for _, item := range pl.Tracks {
		for _, deleteID := range tracks {
			if strings.EqualFold(item.ID.String(), deleteID.String()) {
				body.AddTrack(item)
			}
		}
	}
	vals, err := body.ParamsToValues()
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("users", string(c.UserId), "playlists", pl.Kind.String(), "change")
	resp, err := c.Http.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Поставить лайк плейлисту.
//
// kind и uid можно получить из плейлиста.
func (c Client) LikePlaylist(ctx context.Context, kind schema.ID, ownerUid schema.ID) (schema.Response[string], error) {
	// POST /users/{userId}/likes/playlists/add
	body := schema.KindOwnerUidRequestBody{
		Kind:     kind,
		OwnerUid: ownerUid,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return schema.Response[string]{}, err
	}
	return addRemove(ctx, &c, vals, true, "playlists")
}

// Убрать лайк с плейлиста.
//
// kind и uid можно получить из плейлиста.
func (c Client) UnlikePlaylist(ctx context.Context, kind schema.ID, uid schema.ID) (schema.Response[string], error) {
	// POST /users/{userId}/likes/playlists/{kind}/remove
	body := schema.PlaylistsIdsRequestBody{}
	body.Add(kind, uid)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return schema.Response[string]{}, err
	}
	return addRemove(ctx, &c, vals, false, "playlists")
}

// Поставить лайк плейлистам.
//
// kind и uid можно получить из плейлиста.
func (c Client) LikePlaylists(ctx context.Context, kindUid map[schema.ID]schema.ID) (schema.Response[string], error) {
	return c.likeUnlikePlaylists(ctx, kindUid, true)
}

// Снять лайк с плейлистов.
//
// kind и uid можно получить из плейлиста.
func (c Client) UnlikePlaylists(ctx context.Context, kindUid map[schema.ID]schema.ID) (schema.Response[string], error) {
	return c.likeUnlikePlaylists(ctx, kindUid, false)
}

func (c Client) likeUnlikePlaylists(ctx context.Context, kindUid map[schema.ID]schema.ID, like bool) (schema.Response[string], error) {
	// POST /users/{userId}/likes/playlists/add-multiple
	// // POST /users/{userId}/likes/playlists/remove
	body := schema.PlaylistsIdsRequestBody{}
	body.AddMany(kindUid)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return schema.Response[string]{}, err
	}
	return addRemoveMultiple(ctx, &c, vals, like, "playlists")
}
